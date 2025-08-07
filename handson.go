package main

import (
	"fmt"
	"net/http"
)

func HandsonRoutes() {
	http.HandleFunc("/for", func(w http.ResponseWriter, r *http.Request) {
		for i := 1; i <= 5; i++ {
			fmt.Fprintf(w, "Line %d\n", i)
		}
	})
}


