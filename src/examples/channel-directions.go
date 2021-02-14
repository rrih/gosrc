package main

import (
	"fmt"
	"time"
)

// 送信専用
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// 受信専用
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
	time.Sleep(time.Second)
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	// 送信用、受信用とそれぞれ明示的にすることで型安全性が増す
	ping(pings, "passed message") // ping()の引数であるpingsは送信専用。
	pong(pings, pongs)            // pingsから受信しようとするとコンパイルエラーになる
	fmt.Println(<-pongs)
}
