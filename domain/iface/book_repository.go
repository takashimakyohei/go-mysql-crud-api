package iface

import (
	entity "go-mysql-crud/domain/model/book"
)

type (
	IBookRepository interface {
		ListBook() ([]*entity.Book, error)
		GetBook(id int) (*entity.Book, error)
		CreateBook(*entity.Book) (*entity.Book, error)
		UpdateBook(*entity.Book) (*entity.Book, error)
		DeleteBook(id int) error
	}
)
