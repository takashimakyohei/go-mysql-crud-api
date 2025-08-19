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

	http.HandleFunc("/pointer", func(w http.ResponseWriter, r *http.Request) {
        valueStr := r.URL.Query().Get("value")
        i, err := strconv.Atoi(valueStr)
        if err != nil {
            fmt.Fprintf(w, "エラー: 整数値を入力してください\n")
            return
        }

        p := &i // iのアドレスをpに代入
        fmt.Fprintf(w, "初期値: i = %d, *p = %d\n", i, *p)

        *p = *p + 10 // ポインタ経由で値を変更
        fmt.Fprintf(w, "変更後: i = %d, *p = %d\n", i, *p)
    })

	http.HandleFunc("/slice", func(w http.ResponseWriter, r *http.Request) {
		// スライスの作成(var)
		var fruits []string

		// スライスの作成
		fruits2 := []string{"grape", "banana"}

		// 要素追加
		fruits = append(fruits, "orange")

		fmt.Fprintf(w, "fruits: %v\n", fruits)
		// スライスの特定要素にアクセス
        fmt.Fprintf(w, "index 0 value is: %v\n", fruits[0])
        fmt.Fprintf(w, "fruits: %v\n", fruits2)
		fmt.Fprintf(w, "len: %d\n", len(fruits))
	})
}

