package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// ワーカー実行
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// 3つの仕事をワーカーに送る
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	// チャネルを閉じる
	close(jobs)
	fmt.Println("sent all jobs")

	// ワーカーを待つ
	<-done
}
