package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func HandsonRoutes() {
	http.HandleFunc("/for", func(w http.ResponseWriter, r *http.Request) {
		for i := 1; i <= 5; i++ {
			fmt.Fprintf(w, "Line %d\n", i)
		}
	})

	http.HandleFunc("/if-bool", func(w http.ResponseWriter, r *http.Request) {
		// クエリパラメータから値を取��
		valueStr := r.URL.Query().Get("value")

		// 文字列を整数に変換
		value, err := strconv.Atoi(valueStr)

		if err != nil {
			fmt.Fprintf(w, "エラー: 整数値を入力してください\n")
			return
		}

		if value > 10 {
			fmt.Fprintf(w, "入力された値は10より大きいです: %d\n", value)
		} else {
			fmt.Fprintf(w, "入力された値は10以下です: %d\n", value)
		}
	})
}
