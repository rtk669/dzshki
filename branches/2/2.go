package main

import "fmt"

func main() {
	var weather int

	fmt.Scan(&weather)
	fmt.Println(weather)

	switch {
	case weather != 3:
		fmt.Println("pivo cool")
	default:
		fmt.Println("idi nahui")
	}
}
