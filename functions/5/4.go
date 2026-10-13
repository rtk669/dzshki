package main

import "fmt"

var x int

func PlusX() {
	x++
}

func PlusY(y int) int {
	y++
	return y
}

func main() {
	var y int = 0
	fmt.Println(x, y)

	PlusX()
	y = PlusY(y)
	fmt.Println(x, y)
}
