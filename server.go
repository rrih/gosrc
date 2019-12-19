package main

import (
	"fmt"
	"net/http"
)

func main() {
	// ルーティング
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World")
	})
	// ポート8080番でサーバー起動
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}