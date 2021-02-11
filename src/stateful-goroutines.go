package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// 読み込み用のリクエストをカプセル化したもの
type readOp struct {
	key  int
	resp chan int
}

// 書き込み用のリクエストをカプセル化したもの
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

// go run stateful-goroutines.go
func main() {
	// 操作を実行した回数を数える
	var readOps uint64
	var writeOps uint64

	// チャネルを使って他のゴルーチンは読み書きのリクエストする
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// このゴルーチンは state を所有する
	// この状態管理用のゴルーチンだけが読み書きをする
	// このゴルーチンは reads , writes を繰り返し select し届いたリクエストに返信する
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
