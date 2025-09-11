package usecase

import (
	"go-mysql-crud/domain/iface"
	entity "go-mysql-crud/domain/model/book"
	dto "go-mysql-crud/dto"
)

type UpdateBookUsecase struct {
	irepo iface.IBookRepository
}

func NewUpdateBookUsecase(repo iface.IBookRepository) *UpdateBookUsecase {
	return &UpdateBookUsecase{irepo: repo}
}

func (uc *UpdateBookUsecase) Execute(id int, req dto.RequestParam) (*entity.Book, error) {
	dbBook, err := uc.irepo.GetBook(id)
	if err != nil || dbBook == nil {
		return nil, err
	}

	dbBook.Title = req.Title
	dbBook.Author = req.Author

	return uc.irepo.UpdateBook(dbBook)
}
