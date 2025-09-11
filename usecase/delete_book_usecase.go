package usecase

import (
	"go-mysql-crud/domain/iface"
)

type DeleteBookUsecase struct {
	irepo iface.IBookRepository
}

func NewDeleteBookUsecase(repo iface.IBookRepository) *DeleteBookUsecase {
	return &DeleteBookUsecase{irepo: repo}
}

func (uc *DeleteBookUsecase) Execute(id int) error {
	dbBook, err := uc.irepo.GetBook(id)
	if err != nil || dbBook == nil {
		return err
	}

	return uc.irepo.DeleteBook(id)
}
