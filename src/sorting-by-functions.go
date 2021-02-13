package main

import (
	"fmt"
	"sort"
)

// Go で独自の関数を使ってソートするには、そのための型を定義しなければならない。
// これは []string のただのエイリアス
type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
