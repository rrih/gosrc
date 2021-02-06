package main

import (
	"fmt"
	"time"
)

// レート制限
// ゴルーチン、チャネル、tickersを使ってリソース使用量を管理する
func main() {
	// リクエストを受け取る量を指定
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		// 流し込む
		requests <- i
	}
	// 閉じる
	close(requests)

	// limiter チャネルからの受信がブロックするのを利用して200msごとにリクエストを受信
	limiter := time.Tick(200 * time.Millisecond)
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// リクエストの短期的なバーストを許容しながらも長期的には
	// レート制限を守らせることもできる
	// limiterにバッファをつける。
	// チャネルburstyLimiterは3つまでのリクエストのバーストを許容する
	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		// バーストを表す3つの要素をチャネルに送る
		burstyLimiter <- time.Now()
	}

	go func() {
		// 200msごとにburstyLimiterに値を送信する。
		// ただしチャネルの要素数は最大3
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	// ここで5つリクエストが届いたとする。
	// そのうちはじめの3つはバースト耐性によって処理される
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
