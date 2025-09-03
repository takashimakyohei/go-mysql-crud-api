package book

import (
	"database/sql"
	"encoding/json"
	"fmt"
	repo "go-mysql-crud/infra/repository"
	usecase "go-mysql-crud/usecase"
	"net/http"
)

type Handler struct {
	DB *sql.DB
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {

	bookRepo := repo.NewBookRepository(h.DB)
	uc := usecase.NewListBookUsecase(bookRepo)

	// データ取得
	books, err := uc.Execute()
	if err != nil {
		http.Error(w, "データ取得に失敗しました", http.StatusInternalServerError)
		return
	}

	// JSONに変換
	jsonData, err := json.Marshal(books)
	if err != nil {
		http.Error(w, "JSONエンコードに失敗しました", http.StatusInternalServerError)
		return
	}

	// レスポンス返却
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (h *Handler) Show(w http.ResponseWriter, r *http.Request, id string) {
	fmt.Println("Request ID:", id)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Book detail"))
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Book create"))
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request, id string) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Book update"))
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request, id string) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Book delete"))
}
