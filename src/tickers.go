package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			// done
			case <-done:
				return
			// 500 msごとに届く値を繰り返し受け取る
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// 1600 ms待つ
	time.Sleep(1600 * time.Millisecond)
	// ticker を止める。それ以降そのチャネルは値を受信しなくなる。
	ticker.Stop()
	// 無名関数を抜けさせる
	done <- true
	fmt.Println("Ticker stopped")
}
