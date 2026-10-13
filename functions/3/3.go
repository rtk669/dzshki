package main

import (
	"errors"
	"fmt"
	"math"
)

func Square(a float64, b float64, c float64) (float64, float64, error) {
	var x1 float64
	var x2 float64

	d := math.Pow(b, 2) - 4*a*c

	if d < 0 {
		return 0, 0, errors.New("нет корней")
	}

	x1 = (-b + math.Sqrt(d)) / (2 * a)
	x2 = (-b - math.Sqrt(d)) / (2 * a)

	return x1, x2, nil
}

func main() {
	var a float64 = 1
	var b float64 = 5
	var c float64 = 1

	x1, x2, err := Square(a, b, c)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("корни уравнения", x1, x2)
}
