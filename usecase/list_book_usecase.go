package usecase

import (
	"go-mysql-crud/domain/iface"
	entity "go-mysql-crud/domain/model/book"
)

type ListBookUsecase struct {
	irepo iface.IBookRepository
}

func NewListBookUsecase(repo iface.IBookRepository) *ListBookUsecase {
	return &ListBookUsecase{irepo: repo}
}

func (uc *ListBookUsecase) Execute() ([]*entity.Book, error) {
	return uc.irepo.ListBook()
}
