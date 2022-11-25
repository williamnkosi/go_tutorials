package main

import "fmt"

type shape interface {
	getArea() float64
}
type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func main() {
	tri := triangle{height: 5, base: 4}
	squ := square{
		sideLength: 10,
	}
	printArea(tri)
	printArea(squ)

}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	area := t.height * t.base * 0.5
	return area
}

func (s square) getArea() float64 {
	area := s.sideLength * s.sideLength
	return area
}
