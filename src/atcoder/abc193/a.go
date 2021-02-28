package main

import (
	"fmt"
)

func main() {
	var a, b float64
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Println(float64((a-b)/a) * 100)
}
