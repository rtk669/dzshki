package main

import "fmt"

func main() {
	var x int
	var y float64
	var z string
	var d bool

	var x1 *int
	y1 := &y
	z1 := &z
	d1 := &d

	fmt.Print(x, "\n", y, "\n", z, "\n", d, "\n", x1, "\n", y1, "\n", z1, "\n", d1, "\n")

}
