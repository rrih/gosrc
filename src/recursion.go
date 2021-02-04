package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println("1:", fact(1))
	fmt.Println("2:", fact(2))
	fmt.Println("3:", fact(3))
	fmt.Println("7:", fact(7))
}
