package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		// 変数 i を閉じ込めている
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()
	// 何度呼ばれても足される値(i)は 1 である
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}
