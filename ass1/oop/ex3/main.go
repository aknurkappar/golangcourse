package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	a float64
	b float64
}

func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func (rectangle Rectangle) area() float64 {
	return rectangle.a * rectangle.b
}

func printArea(shape Shape) {
	fmt.Println("Area:", shape.area())
}

func main() {
	c := Circle{radius: 4}
	r := Rectangle{a: 2, b: 10}

	printArea(c)
	printArea(r)
}
