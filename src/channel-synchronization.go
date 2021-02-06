package main

import (
	"fmt"
	"time"
)

func main() {
	// done チャネルを作る
	done := make(chan bool, 1)
	// worker ゴルーチンを実行する
	go worker(done)
	// ブロック
	<-done
}

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}
