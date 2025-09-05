package usecase

import (
	entity "go-mysql-crud/domain/model/book"
	"go-mysql-crud/infra/repository"
)

type ListBookUsecase struct {
	irepo repository.IBookRepository
}

func NewListBookUsecase(repo repository.IBookRepository) *ListBookUsecase {
	return &ListBookUsecase{irepo: repo}
}

func (uc *ListBookUsecase) Execute() ([]*entity.Book, error) {
	return uc.irepo.ListBook()
}
