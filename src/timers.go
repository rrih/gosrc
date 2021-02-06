package main

import (
	"fmt"
	"time"
)

func main() {
	// 2秒待つタイマー
	timer1 := time.NewTimer(2 * time.Second)

	// タイマーのチャネルCにタイマーが切れたことを示す値が届くまでブロック
	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	// キャンセルしてる
	stop2 := timer2.Stop()
	// キャンセルの確認
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
	time.Sleep(2 * time.Second)
}
