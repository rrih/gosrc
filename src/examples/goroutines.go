package main

import (
	"fmt"
	"time"
)

func main() {
	f("direct")

	// 関数をゴルーチンの中で実行する
	// 新たなゴルーチンが呼び出し側と並行に実行される
	go f("goroutine")

	// 無名関数のゴルーチン
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// 処理が終わるのを待つ
	time.Sleep(time.Second)
	fmt.Println("done")
}

func f(from string) {
	for i := 0; i < 10; i++ {
		fmt.Println(from, ":", i)
	}
}
