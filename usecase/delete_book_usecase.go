package usecase

import (
	"go-mysql-crud/infra/repository"
)

type DeleteBookUsecase struct {
	irepo repository.IBookRepository
}

func NewDeleteBookUsecase(repo repository.IBookRepository) *DeleteBookUsecase {
	return &DeleteBookUsecase{irepo: repo}
}

func (uc *DeleteBookUsecase) Execute(id int) error {
	dbBook, err := uc.irepo.GetBook(id)
	if err != nil || dbBook == nil {
		return err
	}

	return uc.irepo.DeleteBook(id)
}
