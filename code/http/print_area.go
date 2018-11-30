/*
Write a program that creates two custom struct types called triangle and square.

The square type should a struct with a field called sideLength of type float64.

The triangle type should be struct with a field called base of type float64 and a field called height of type float64.

Both types should have function called getArea that returns the calculated area of the square or triangle.

Add a shape interface that defines a function called getArea. This function should calculate the area of the given shape
and return it. So that getArea function can be called with either a triangle or a square.
 */

package main

import (
	"fmt"
	"math"
)

type triangle struct {
	height float64
	base float64
}

type square struct {
	sideLength float64
}

type circle struct {
	radius float64
}

type shape interface {
	getArea() float64
}

func (t triangle) getArea () float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea () float64 {
	return s.sideLength * s.sideLength
}

func (c circle) getArea () float64 {
	return math.Pi * c.radius * c.radius
}
func printArea(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	t := triangle{height:4, base:5}
	s := square{sideLength:3}
	c := circle{radius:2}
	printArea(t)
	printArea(s)
	printArea(c)
}
