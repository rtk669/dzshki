package main

import (
	"fmt"
	"math"
)

type figure interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	a float64
	b float64
}

func (s Rectangle) Area() float64 {
	return s.a * s.b
}

func (s Rectangle) Perimeter() float64 {

	return s.a*2 + s.b*2
}

type Triangle struct {
	a float64
	b float64
	c float64
}

func (t Triangle) SomeSh() {

}

func (s Triangle) Area() float64 {
	p := (s.a + s.b + s.c) / 2

	// Формула Герона: S = sqrt(p * (p - a) * (p - b) * (p - c))
	return math.Sqrt(p * (p - s.a) * (p - s.b) * (p - s.c))
}

func (s Triangle) Perimeter() float64 {

	return s.a + s.b + s.c
}

type Calculator struct {
	calc figure
}

func main() {
	calc1 := Calculator{
		Triangle{
			a: 7,
			b: 3,
			c: 4,
		},
	}
	fmt.Println(calc1.calc.Area())

}
