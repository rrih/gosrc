package main

import "fmt"

func zeroval(ival int) {
	// 意味ない
	ival = 0
}

func zeroptr(iptr *int) {
	fmt.Println("iptr pointer:", *iptr)
	// ポインタの参照を剥がしてる
	*iptr = 2
}

func main() {
	i := 1
	fmt.Println("pointer:", &i)
	fmt.Println("initial:", i)
	zeroval(i)
	fmt.Println("zeroval:", i)
	// i をポインタとして渡す
	zeroptr(&i)
	fmt.Println("zeroptr:", i)
	fmt.Println("pointer:", &i)
}
