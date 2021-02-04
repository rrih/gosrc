package main

import "fmt"

// 四角の構造体
type rect struct {
	width  int
	height int
}

// area のレシーバの型は *rect
func (r *rect) area() int {
	return r.width * r.height
}

// 周囲 perimeter
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 10}
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	// レシーバのポインタを渡す
	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
}
