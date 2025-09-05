package usecase

import (
	entity "go-mysql-crud/domain/model/book"
	dto "go-mysql-crud/dto"
	"go-mysql-crud/infra/repository"
)

type UpdateBookUsecase struct {
	irepo repository.IBookRepository
}

func NewUpdateBookUsecase(repo repository.IBookRepository) *UpdateBookUsecase {
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
