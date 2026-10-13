package main

import (
	atm "dz_err/dz_2/ATM"
	card "dz_err/dz_2/Card"
	"fmt"
)

func main() {

	card := card.NewCard(1000)
	atm := atm.NewATM(card)

	if err := atm.GetCash(-200); err != nil {
		fmt.Println(err)
	}

	if err := atm.GetBalance(); err != nil {
		fmt.Println(err)
	}

	if err := atm.OnlinePay(7000); err != nil {
		fmt.Println(err)
	}
}
