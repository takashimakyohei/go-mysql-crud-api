package repository

import (
	"database/sql"
	entity "go-mysql-crud/domain/model/book"
	"log"
)

type (
	IBookRepository interface {
		ListBook() ([]*entity.Book, error)
		GetBook(id int) (*entity.Book, error)
		CreateBook(*entity.Book) (*entity.Book, error)
	}
	BookRepository struct {
		DB *sql.DB
	}
)

func NewBookRepository(db *sql.DB) IBookRepository {
	return &BookRepository{DB: db}
}

func (r *BookRepository) ListBook() ([]*entity.Book, error) {
	books := []*entity.Book{}
	rows, err := r.DB.Query("SELECT id, title, author, created_at, updated_at FROM books ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		b := &entity.Book{}
		err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.CreatedAt, &b.UpdatedAt)
		if err != nil {
			return nil, err
		}
		log.Printf("書籍データ読み込み: ID=%d, Title=%s, Author=%s", b.ID, b.Title, b.Author)
		books = append(books, b)
	}

	return books, nil
}

func (r *BookRepository) GetBook(id int) (*entity.Book, error) {
	b := entity.Book{}
	row := r.DB.QueryRow(`
        SELECT id, title, author, created_at, updated_at
        FROM books
        WHERE id = ?
    `, id)

	err := row.Scan(&b.ID, &b.Title, &b.Author, &b.CreatedAt, &b.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &b, nil
}

func (r *BookRepository) CreateBook(b *entity.Book) (*entity.Book, error) {
	// Insert
	result, err := r.DB.Exec(
		"INSERT INTO books (title, author) VALUES (?, ?)",
		b.Title, b.Author,
	)
	if err != nil {
		return nil, err
	}

	// Insert後の自動採番IDを取得
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	b.ID = int(id)

	// 完成したEntityを返す
	return b, nil
}
