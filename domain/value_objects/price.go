package value_objects

import "fmt"

type Price struct {
	Amount float64
}

func NewPrice(amount float64) (Price, error) {
	if amount < 0 {
		return Price{}, fmt.Errorf("price cannot be negative")
	}
	return Price{Amount: amount}, nil
}

func (p Price) Add(other Price) Price {
	return Price{Amount: p.Amount + other.Amount}
}

func (p Price) GetAmount() float64 {
	return p.Amount
}
