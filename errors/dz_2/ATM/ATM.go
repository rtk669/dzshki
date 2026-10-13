package atm

import (
	card "dz_err/dz_2/Card"
	"errors"
	"fmt"
	"math/rand"
)

type ATM struct {
	Card *card.Card
}

func NewATM(card *card.Card) *ATM {
	return &ATM{
		Card: card,
	}
}

func err() error {
	x := rand.Intn(10)
	if x < 3 {
		return errors.New("ошибка")
	}
	return nil
}

func (a *ATM) GetCash(cash int) error {
	if err := err(); err != nil {
		return err
	}
	if a.Card.Balance-cash < 0 {
		return errors.New("недостаточно средств")
	}
	if cash < 0 {
		return errors.New("что за херню ты хочешь получить")
	}
	a.Card.Balance -= cash
	return nil
}

func (a *ATM) EnterCash(cash int) error {
	if err := err(); err != nil {
		return err
	}
	if cash < 0 {
		return errors.New("неправильно")
	}

	a.Card.Balance += cash

	return nil
}

func (a *ATM) GetBalance() error {
	if err := err(); err != nil {
		return err
	}

	fmt.Println(a.Card.Balance)

	return nil
}

func (a *ATM) OnlinePay(amount int) error {
	if err := err(); err != nil {
		return err
	}

	if a.Card.Balance-amount < 0 {
		return errors.New("недостаточно средств")
	}
	if amount < 0 {
		return errors.New("что за херню ты хочешь получить")
	}
	a.Card.Balance -= amount
	return nil

}
