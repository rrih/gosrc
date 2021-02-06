package main

import "fmt"

// https://www.spinute.org/go-by-example/channels.html
// チャネル( channel )は並行に動くゴルーチンを繋ぐパイプ。
// あるゴルーチンからチャネルに値を送り別のゴルーチンで受け取れる。
func main() {
	// make(chan value-type) で新しいチャネルを作る。
	// チャネルの型にはそれを通る値の型を入れる
	messages := make(chan string)

	go func() {
		// channel <- と描けばチャネルに値を送信できる。
		// ここでは新たなゴルーチンから"ping"を先に作ったチャネル messages に送る
		messages <- "ping"
	}()

	// <-channel と描けばそのチャネルから値を受信。
	// ここでは"ping"を受信、表示してる。
	msg := <-messages
	fmt.Println(msg)
}
