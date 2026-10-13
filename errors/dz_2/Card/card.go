package card

type Card struct {
	Balance int
}

func NewCard(balance int) *Card {
	return &Card{
		Balance: balance,
	}
}
