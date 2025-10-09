package main

import (
	"fmt"
	"math"
)

type Geometric interface {
	area() float64
}

type Square struct {
	side float64
} // lado ^ 2

func (s Square) area() float64 {
	return s.side * s.side
}

type Circle struct {
	radius float64
} // pi * (raio ^ 2)

func (c Circle) area() float64 {
	return math.Pi * (c.radius * c.radius)
}
func CreateInterface() {
	var s Geometric = Square{3}
	var c Geometric = Circle{5}

	fmt.Printf("A área do quadrado tem a área de %.2f\n", s.area())
	fmt.Printf("A área do círculo tem a área de %.2f\n", c.area())

}
