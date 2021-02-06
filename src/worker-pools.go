package main

import (
	"fmt"
	"time"
)

// このワーカーを複数並行に動かす
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		// 仕事をするたびに1秒スリープする
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

// ゴルーチンとチャネルを使ってワーカープールを実装
// ワーカープールを使うために、仕事を送り、結果を収集する。
// そのために使う 2 つのチャネルを作る。
func main() {
	const numJobs = 5
	// 仕事
	jobs := make(chan int, numJobs)
	// 結果
	results := make(chan int, numJobs)

	// 3つワーカーを起動。まだジョブがないのでいずれのワーカーもブロック。
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// jobs に5つ値を送る。
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	// チャネルを閉じて仕事を送り終えたことを伝える。
	close(jobs)

	// 仕事の結果を集める。ここでゴルーチンのワーカーが終了したことも保証される。
	// WaitGroup を使うやり方もある。わからない。
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
