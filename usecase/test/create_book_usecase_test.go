package usecase_test

import (
	"go-mysql-crud/domain/model/book"
	"go-mysql-crud/dto"
	"go-mysql-crud/mock"
	"go-mysql-crud/usecase"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateBookUsecase_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockIBookRepository(ctrl)
	uc := usecase.NewCreateBookUsecase(mockRepo)

	req := dto.RequestParam{
		Title:  "Go入門",
		Author: "山田太郎",
	}

	expected := &book.Book{
		Title:  req.Title,
		Author: req.Author,
	}

	// mock の振る舞いを定義
	mockRepo.
		EXPECT().
		CreateBook(gomock.Any()).
		Return(expected, nil).
		Times(1)

	// 実行
	res, err := uc.Execute(req)

	// 検証
	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}
