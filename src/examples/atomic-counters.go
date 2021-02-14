package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Go で状態を管理する時、まずチャネル通信を使う方法を検討すべき。
// 別のやり方として、sync/atomic パッケージにある、複数のゴルーチンから
// アクセスされるアトミックなかうんたを使う方法。
func main() {
	// 符号なし整数型の変数で(負の値を取らない)カウンタを表現する
	var ops uint64
	// WaitGroupを使って全てのゴルーチンが仕事を終えるのを待つ。
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				// AddUint64関数にはopsのメモリアドレスを渡す。
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}
	// 待つ
	wg.Wait()
	fmt.Println("ops:", ops)

}
