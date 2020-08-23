package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type square struct {
	length float64
}
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (sq square) area() float64 {
	return sq.length * sq.length
}

func info(sh shape) {
	fmt.Printf("%T\n", sh)
	fmt.Println(sh.area())
}

func main() {
	circ := circle{
		radius: 12.345,
	}

	squa := square{
		length: 15,
	}

	info(circ)
	info(squa)
}
