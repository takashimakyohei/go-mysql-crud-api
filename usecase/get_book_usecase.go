package usecase

import (
	entity "go-mysql-crud/domain/model/book"
	"go-mysql-crud/infra/repository"
)

type GetBookUsecase struct {
	irepo repository.IBookRepository
}

func NewGetBookUsecase(repo repository.IBookRepository) *GetBookUsecase {
	return &GetBookUsecase{irepo: repo}
}

func (uc *GetBookUsecase) Execute(id int) (*entity.Book, error) {
	return uc.irepo.GetBook(id)
}
