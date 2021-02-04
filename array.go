package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp: ", a)
	a[4] = 1000
	fmt.Println(a)
	fmt.Println(a[4])
	fmt.Println(len(a))

	// 一行で宣言と初期化はこのように書く
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	// 多次元配列
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("二次元配列: ", twoD)
}
