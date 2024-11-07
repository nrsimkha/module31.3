package storage

// Post - публикация.
type Post struct {
	ID         int
	Author_id  int
	Title      string
	Content    string
	Created_at int
}

// Interface задаёт контракт на работу с БД.
type Interface interface {
	Posts() ([]Post, error) // получение всех публикаций
	AddPost(Post) error     // создание новой публикации
	UpdatePost(Post) error  // обновление публикации
	DeletePost(Post) error  // удаление публикации по ID
}
