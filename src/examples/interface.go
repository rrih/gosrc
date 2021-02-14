package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

// react のメソッド
func (r rect) area() float64 {
	return r.width * r.height
}

// react のメソッド
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// circle のメソッド
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// circle のメソッド
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	r := rect{width: 10, height: 4}
	c := circle{radius: 5}
	measure(r)
	measure(c)
}

// 汎用メソッド
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
