package main

import (
	"fmt"
	"net/http"
)

func main() {
	// ルーティング
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmtパッケージによって提供されているprintf関数
		// 第一引数で書き込み先を明示的に指定 ここではwを指定している
		fmt.Fprintf(w, "Hello, World\n")
		fmt.Fprintf(w, "This is a Pen\n")
	})
	// ポート8080番でサーバー起動
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}