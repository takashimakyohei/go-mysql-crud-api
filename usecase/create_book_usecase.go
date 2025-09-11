package usecase

import (
	"go-mysql-crud/domain/iface"
	entity "go-mysql-crud/domain/model/book"
	dto "go-mysql-crud/dto"
)

type CreateBookUsecase struct {
	irepo iface.IBookRepository
}

func NewCreateBookUsecase(repo iface.IBookRepository) *CreateBookUsecase {
	return &CreateBookUsecase{irepo: repo}
}

func (uc *CreateBookUsecase) Execute(req dto.RequestParam) (*entity.Book, error) {
	entity := &entity.Book{
		Title:  req.Title,
		Author: req.Author,
	}
	return uc.irepo.CreateBook(entity)
}
