package main

import (
	"database/sql"
	"fmt"
	bookHandler "go-mysql-crud/handler"
	"log"
	"net/http"
	"os"
	"strings"
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

// データベースの初期化（テーブル作成とサンプルデータ挿入）
func initDatabase(db *sql.DB) error {
	// SQLファイルを読み込み
	sqlPath := "sql/init.sql"
	log.Printf("SQLファイルパス: %s", sqlPath)

	// ファイルの存在確認
	if _, err := os.Stat(sqlPath); os.IsNotExist(err) {
		return fmt.Errorf("SQLファイルが存在しません: %s", sqlPath)
	}

	sqlBytes, err := os.ReadFile(sqlPath)
	if err != nil {
		return fmt.Errorf("SQLファイルの読み込みに失敗: %v", err)
	}

	log.Printf("SQLファイルサイズ: %d bytes", len(sqlBytes))

	// SQLを実行（複数文を分けて実行）
	sqlContent := string(sqlBytes)
	log.Printf("読み込んだSQL内容: %s", sqlContent)

	// セミコロンで分割し、空行とコメントを除去
	statements := strings.Split(sqlContent, ";")
	log.Printf("分割されたSQL文の数: %d", len(statements))

	for i, stmt := range statements {
		stmt = strings.TrimSpace(stmt)
		log.Printf("SQL文 [%d] (trim前): '%s'", i+1, statements[i])
		log.Printf("SQL文 [%d] (trim後): '%s'", i+1, stmt)

		if stmt == "" || strings.HasPrefix(stmt, "--") {
			log.Printf("SQL文 [%d] をスキップ (空またはコメント)", i+1)
			continue
		}

		log.Printf("SQL実行 [%d]: %s", i+1, stmt)
		_, err := db.Exec(stmt)
		if err != nil {
			return fmt.Errorf("SQL実行エラー [%d]: %v, SQL: %s", i+1, err, stmt)
		}
		log.Printf("SQL実行成功 [%d]: %s", i+1, stmt)
	}

	log.Println("データベースの初期化が完了しました")
	return nil
}

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=true&loc=Local",
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

	// データベースの初期化
	if err := initDatabase(db); err != nil {
		log.Fatal("データベース初期化に失敗しました: ", err)
	}

	// 標準のhttpパッケージのみを使用したルーティング
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	// ハンズオン用ルーティングを別ファイルから登録
	HandsonRoutes()

	// Bookハンドラーのルーティングを登録
	bh := &bookHandler.Handler{DB: db}
	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			bh.Index(w, r)
		case http.MethodPost:
			bh.Create(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	// 詳細・更新・削除
	http.HandleFunc("/books/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/books/")
		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		switch r.Method {
		case http.MethodGet:
			bh.Show(w, r, id)
		case http.MethodPut:
			bh.Update(w, r, id)
		case http.MethodDelete:
			bh.Delete(w, r, id)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	log.Println("APIサーバー起動 :8080")
	http.ListenAndServe(":8080", nil)
}
