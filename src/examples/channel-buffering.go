package main

import "fmt"

func main() {
	// チャネルにはデフォルトではバッファはついていない。
	// 2 つまで値を溜められるバッファ付きのチャネルを作る
	messsages := make(chan string, 2)
	messsages <- "buffered"
	messsages <- "channel"
	fmt.Println(<-messsages)
	fmt.Println(<-messsages)
	// 例えば、make()の第二引数つかなかったら、以下のようなエラーが出る
	// fatal error: all goroutines are asleep - deadlock!

	// goroutine 1 [chan send]:
	// main.main()
	// 		/src/channel-buffering.go:9 +0x59
	// exit status 2
}
