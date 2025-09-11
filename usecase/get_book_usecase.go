package usecase

import (
	"go-mysql-crud/domain/iface"
	entity "go-mysql-crud/domain/model/book"
)

type GetBookUsecase struct {
	irepo iface.IBookRepository
}

func NewGetBookUsecase(repo iface.IBookRepository) *GetBookUsecase {
	return &GetBookUsecase{irepo: repo}
}

func (uc *GetBookUsecase) Execute(id int) (*entity.Book, error) {
	return uc.irepo.GetBook(id)
}
