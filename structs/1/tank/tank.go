package tank

import (
	"errors"
	"fmt"
)

type ton float64
type mm int

type Tank struct {
	name   string
	weight ton
	gun    mm
	amount int
}

func NewTank(
	name string,
	weight ton,
	gun mm,
	amount int,
) (*Tank, error) {
	if weight < 5 {
		return nil, errors.New("танк пробьет даже плевок")
	}

	if weight > 60 {
		return nil, errors.New("танк утонет в грязи")
	}

	return &Tank{
		name:   name,
		weight: weight,
		gun:    gun,
		amount: amount,
	}, nil
}

func (t *Tank) Produce(n int) {
	t.amount += n
}

func (t *Tank) Lose(n int) {
	t.amount -= n
}

func (t Tank) TankStats() {
	fmt.Println()
	fmt.Println("tank stats")
	fmt.Printf("%+v\n%+v\n%+v\n", t.name, t.weight, t.gun)
	fmt.Println()
}
