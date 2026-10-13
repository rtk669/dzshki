package main

import (
	"fmt"
	"math/rand/v2"
)

func Zaselenie(n int, nomera []bool) {
	if nomera[n] == false {
		fmt.Println("вот ключ от вашего номера", n)
		nomera[n] = true
	} else {
		fmt.Println("номер занят")
	}

}

func main() {
	var n int

	fmt.Scan(&n)

	nomera := make([]bool, 5)

	fmt.Println(nomera)

	for range n {
		fmt.Println("здравствйте")

		x := rand.IntN(5)
		Zaselenie(x, nomera)
	}

	fmt.Println(nomera)
}
