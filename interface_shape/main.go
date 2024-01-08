package main

import (
	"fmt"
)

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

func main() {
	t := triangle{
		base:   10,
		height: 2,
	}

	s := square{
		sideLength: 3,
	}

	printArea(t)
	printArea(s)
}

func (t triangle) getArea() float64 {
	area := 0.5 * t.base * t.height
	return area
}

func (s square) getArea() float64 {
	area := s.sideLength * s.sideLength
	return area
}

func printArea(sh shape) {
	a := sh.getArea()
	fmt.Println("Area of square is:", a)
}