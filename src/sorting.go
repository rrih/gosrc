package main

import (
	"fmt"
	"sort"
)

// そーと
func main() {
	strs := []string{"c", "a", "b"}
	// 文字列ソート。ソートは in-place
	// 引数に渡したスライスを変更し、新たなスライスを返すわけではない。
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:", ints)

	// ソート済みかどうかの判定
	s := sort.IntsAreSorted(ints)
	ss := sort.StringsAreSorted(strs)
	fmt.Println("Sorted:", s)
	fmt.Println("Sorted:", ss)
}
