package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp: ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s)

	s = append(s, "d")
	s = append(s, "s", "d", "f", "eee")
	fmt.Println(s)

	// copy
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)

	// 2 番目から 5 番目までを抽出
	l := s[2:5]
	fmt.Println("slice: ", l)

	// 5 番目までを抽出
	t := s[:5]
	fmt.Println(t)

	// 2 番目以降を抽出
	tt := s[2:]
	fmt.Println(tt)

	str := []string{"gg", "hh", "ii"}
	fmt.Println("str: ", str)

	twoD := make([][]int, 3)
	// 空の二次元配列を 3 列生成
	fmt.Println(twoD)

	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
