package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	if a <= 10 {
		fmt.Println(a, "は10以下")
		return
	}
	fmt.Println(a, "は10より高い")
	return
}