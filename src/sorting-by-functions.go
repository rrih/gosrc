package main

import (
	"fmt"
	"sort"
)

// Go で独自の関数を使ってソートするには、そのための型を定義しなければならない。
// これは []string のただのエイリアス
type byLength []string

// Less 用の Length を取得する自作メソッド
func (s byLength) Len() int {
	return len(s)
}

// Swap の自作メソッド
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less の自作メソッド
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwiiiiiiii"}
	// 肩をキャストしてソート
	sort.Sort(byLength(fruits))
	nums := []int{1, 4, 66, 44, 12, -1, 10}
	sort.Sort(byLength(nums))
	fmt.Println(fruits)
	fmt.Println(nums)
}
