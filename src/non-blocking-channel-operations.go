package main

import "fmt"

//  https://www.spinute.org/go-by-example/non-blocking-channel-operations.html
// select と default を使うことでノンブロッキングな送受信や選択的な受信を実装できる。
func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// ノンブロッキングな受信
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message recceived")
	}

	// ノンブロッキングな送信
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
