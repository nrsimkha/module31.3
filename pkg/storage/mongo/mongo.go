package mongo

import (
	"GoNews/pkg/storage"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Хранилище данных.
type Store struct {
	db *mongo.Client
}

var ctx context.Context = context.Background()

// Конструктор.
func New(constr string) (*Store, error) {
	mongoOpts := options.Client().ApplyURI(constr)
	client, err := mongo.Connect(ctx, mongoOpts)
	if err != nil {
		log.Fatal(err)
	}
	// не забываем закрывать ресурсы
	//defer client.Disconnect(ctx)
	// проверка связи с БД
	/* err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	} */
	s := Store{
		db: client,
	}
	return &s, nil
}

func (store Store) Posts() ([]storage.Post, error) {
	collection := store.db.Database("goNews").Collection("posts")
	filter := bson.D{}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	var data []storage.Post
	for cur.Next(ctx) {
		var l storage.Post
		err := cur.Decode(&l)
		if err != nil {
			return nil, err
		}
		data = append(data, l)
	}
	return data, cur.Err()

}

func (store Store) AddPost(post storage.Post) error {
	collection := store.db.Database("goNews").Collection("posts")
	_, err := collection.InsertOne(ctx, post)
	if err != nil {
		return err
	}
	return nil
}

func (store Store) UpdatePost(post storage.Post) error {

	coll := store.db.Database("goNews").Collection("posts")
	var id primitive.ObjectID

	// Find the document for which the _id field matches id and set the email to
	// "newemail@example.com".
	// Specify the Upsert option to insert a new document if a document matching
	// the filter isn't found.
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"id", id}}
	update := bson.D{{"$set", bson.D{{"author_id", post.Author_id}, {"title", post.Title}, {"content", post.Content}, {"created_ad", post.Created_at}}}}

	result, err := coll.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		return err
	}

	if result.MatchedCount != 0 {
		fmt.Println("matched and replaced an existing document")

	}
	if result.UpsertedCount != 0 {
		fmt.Printf("inserted a new document with ID %v\n", result.UpsertedID)
	}
	return nil
}

func (store Store) DeletePost(post storage.Post) error {
	coll := store.db.Database("goNews").Collection("posts")
	opts := options.Delete().SetCollation(&options.Collation{
		Locale:    "en_US",
		Strength:  1,
		CaseLevel: false,
	})
	res, err := coll.DeleteOne(context.TODO(), bson.D{{Key: "id", Value: post.ID}}, opts)
	if err != nil {
		return err
	}
	fmt.Printf("deleted %v documents\n", res.DeletedCount)
	return nil
}
