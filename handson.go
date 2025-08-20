package main

import (
	"fmt"
	"net/http"
	"strconv"
    "go-mysql-crud/user"
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

        // あらかじめ容量を指定してスライスを作成
        // tip: スライスの容量を指定することで、メモリの再割り当てを減らすことができる
		b := make([]int, 0, 5)
		b = append(b, 1, 2, 3, 4, 5)
        b = append(b, 1, 2)
		fmt.Fprintf(w, "b: len=%d cap=%d %v\n", len(b), cap(b), b)
	})

	http.HandleFunc("/struct", func(w http.ResponseWriter, r *http.Request) {
        type Student struct {
            Number int
            Name string
        }

        students := []Student{
            {Number: 1, Name: "Alice"},
            {Number: 2, Name: "Bob"},
        }

        // 構造体スライスの出力
        fmt.Fprintf(w, "%v\n", students)
        for _, student := range students {
            fmt.Fprintf(w, "Name: %s\n", student.Name)
        }

        // 構造体の追加
        students = append(students, Student{Number: 3, Name: "Carol"})
        fmt.Fprintf(w, "追加後: %v\n", students)

        // 構造体のフィールド更新
        students[1].Name = "Bob Jr."
        fmt.Fprintf(w, "2番目の生徒の名前を更新: %v\n", students[1])

        // 構造体の比較
        equal := students[0] == students[1]
        fmt.Fprintf(w, "1番目と2番目は同じ？: %v\n", equal)


    })

    http.HandleFunc("/map", func(w http.ResponseWriter, r *http.Request) {
        // 要素が入ったmapの作成
        prices := map[string]int{
            "apple": 120,
            "banana": 80,
        }
        fmt.Fprintf(w, "初期map: %+v\n", prices)

        // 空のmapの作成
        emptymap := map[string]int{}
        fmt.Fprintf(w, "初期map: %+v\n", emptymap)

        prices["orange"] = 150 // 追加
        prices["banana"] = 90  // 更新
        fmt.Fprintf(w, "追加・更新後: %+v\n", prices)

        fruit := "apple"
        price, exists := prices[fruit]
        if exists {
            fmt.Fprintf(w, "%sの値段: %d\n", fruit, price)
        } else {
            fmt.Fprintf(w, "%sは登録されていません\n", fruit)
        }

        delete(prices, "banana")
        fmt.Fprintf(w, "削除後: %+v\n", prices)
    })

	http.HandleFunc("/receiver", func(w http.ResponseWriter, r *http.Request) {
		u := user.User{"Taro", 25}
		// レシーバ付きメソッドの呼び出し
        fmt.Fprintf(w, "名前: %s\n", u.GetName())
        fmt.Fprintf(w, "年齢: %s\n", u.GetAge())

        u.HaveBirthday()
        fmt.Fprintf(w, "誕生日を迎えた後の年齢（値レシーバ）: %s\n", u.GetAge())

        u.HaveBirthdayPointer()
        fmt.Fprintf(w, "誕生日を迎えた後の年齢（ポインタレシーバ）: %s\n", u.GetAge())

	})
}
