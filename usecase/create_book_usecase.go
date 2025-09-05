package usecase

import (
	entity "go-mysql-crud/domain/model/book"
	dto "go-mysql-crud/dto"
	"go-mysql-crud/infra/repository"
)

type CreateBookUsecase struct {
	irepo repository.IBookRepository
}

func NewCreateBookUsecase(repo repository.IBookRepository) *CreateBookUsecase {
	return &CreateBookUsecase{irepo: repo}
}

func (uc *CreateBookUsecase) Execute(req dto.RequestParam) (*entity.Book, error) {
	entity := &entity.Book{
		Title:  req.Title,
		Author: req.Author,
	}
	return uc.irepo.CreateBook(entity)
}
