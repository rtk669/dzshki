package main

import (
	"fmt"
	"str/tank"
)

func main() {
	tank1, err := tank.NewTank(
		"T-34",
		23.7,
		76,
		3000,
	)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(*tank1)
		tank1.Produce(300)
		fmt.Println(*tank1)
		tank1.Lose(1700)
		fmt.Println(*tank1)
		tank1.TankStats()
	}

	tank2, err := tank.NewTank(
		"Big daddy",
		120,
		250,
		3,
	)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(*tank2)
	}

}
