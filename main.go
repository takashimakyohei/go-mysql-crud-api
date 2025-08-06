package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// DB接続をリトライ
func connectWithRetry(driver, dsn string, maxRetries int, waitSec int) (*sql.DB, error) {
	var db *sql.DB
	var err error
	for i := 0; i < maxRetries; i++ {
		db, err = sql.Open(driver, dsn)
		if err == nil {
			err = db.Ping()
			if err == nil {
				return db, nil
			}
			db.Close()
		}
		log.Printf("DB接続失敗: %v (リトライ %d/%d)", err, i+1, maxRetries)
		time.Sleep(time.Duration(waitSec) * time.Second)
	}
	return nil, err
}

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)
	var err error
	db, err := connectWithRetry("mysql", dsn, 10, 3)
	if err != nil {
		log.Fatal("DB接続に失敗しました: ", err)
	}
	defer db.Close()
	fmt.Println("DB接続成功")

	// 標準のhttpパッケージのみを使用したルーティング
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	log.Println("APIサーバー起動 :8080")
	http.ListenAndServe(":8080", nil)
}
