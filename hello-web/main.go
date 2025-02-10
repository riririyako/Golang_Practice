package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	// "/" パスに対するハンドラを登録
	http.HandleFunc("/", hello)

	// サーバーを起動し、第2引数に`nil`を渡すことでデフォルトのマルチプレクサ（`http.DefaultServeMux`）を使用
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("failed to start : ",err)
	}
}
