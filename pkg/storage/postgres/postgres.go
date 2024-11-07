package postgres

import (
	"GoNews/pkg/storage"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Хранилище данных.
type Store struct {
	db *pgxpool.Pool
}

var ctx context.Context = context.Background()

// Конструктор.
func New(constr string) (*Store, error) {
	db, err := pgxpool.New(context.Background(), constr)
	if err != nil {
		return nil, err
	}
	s := Store{
		db: db,
	}
	return &s, nil
}

func (store Store) Posts() ([]storage.Post, error) {
	rows, err := store.db.Query(ctx, `
	SELECT * FROM posts ORDER BY id;
`)
	if err != nil {
		return nil, err
	}
	var posts []storage.Post

	for rows.Next() {
		var p storage.Post
		fmt.Println("start")
		err = rows.Scan(
			&p.ID,
			&p.Author_id,
			&p.Title,
			&p.Content,
			&p.Created_at,
		)
		if err != nil {
			return nil, err
		}
		fmt.Println(p)
		posts = append(posts, p)

	}
	return posts, rows.Err()
}

func (store Store) AddPost(post storage.Post) error {
	_, err := store.db.Exec(ctx, `
			CREATE TABLE IF NOT EXISTS posts (
				id SERIAL PRIMARY KEY, 
				author_id INTEGER REFERENCES authors(id) NOT NULL,
				title TEXT NOT NULL,
				content TEXT NOT NULL,
				created_at BIGINT NOT NULL DEFAULT extract(epoch from now())
			);
	`)
	if err != nil {
		return err
	}
	_, err = store.db.Exec(ctx, `INSERT INTO posts (author_id, title, content) VALUES ($1, $2, $3);`,
		post.Author_id, post.Title, post.Content)
	if err != nil {
		return err
	}

	return nil
}

func (store Store) UpdatePost(post storage.Post) error {
	_, err := store.db.Exec(ctx, `UPDATE posts SET author_id=$1, title=$2, content=$3 WHERE id = $4;`,
		post.Author_id, post.Title, post.Content, post.ID)
	if err != nil {
		return err
	}
	return nil
}

func (store Store) DeletePost(post storage.Post) error {
	_, err := store.db.Exec(ctx, `DELETE FROM posts WHERE id = $1;`,
		post.ID)
	if err != nil {
		return err
	}
	return nil
}
