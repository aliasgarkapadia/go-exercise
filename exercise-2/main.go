package main

import "fmt"

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
	return 3.14 * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(s shape) float64 {
	return s.area()
}

func main() {
	fmt.Println("Hello, 世界")
	c1 := circle{
		radius: 10,
	}
	s1 := square{
		length: 10,
		width:  10,
	}

	fmt.Println(info(c1))
	fmt.Println(info(s1))
}
