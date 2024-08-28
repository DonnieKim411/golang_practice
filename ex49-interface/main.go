package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.width
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type shape interface {
	area() float64
}

func info(s shape) float64 {
	return s.area()
}

func main() {
	s1 := square{
		length: 5.0,
		width:  4.0,
	}
	c1 := circle{
		radius: 1.0,
	}

	fmt.Println(info(s1))
	fmt.Println(info(c1))

}
