package main

import (
	"fmt"
	"time"
)

// https://www.spinute.org/go-by-example/timeouts.html
// タイムアウト機能実装編
// タイムアウト、実行時間に上限を設けるプログラムのこと
// チャネルと select を使うことでタイムアウト機能を実現する
func main() {
	c1 := make(chan string, 1)
	go func() {
		// 2秒後にチャネル c1 に送信する
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		// 2秒後に実行されるため time.After(1 * time.Second) によって2秒後に実行されるこちらが走る
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
