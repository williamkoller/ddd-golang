package entities

import "ddd-golang/domain/value_objects"

type Order struct {
	ID         string
	Products   []string
	TotalPrice value_objects.Price
}

func NewOrder(id string, products []string, totalPrice value_objects.Price) *Order {
	return &Order{
		ID:         id,
		Products:   products,
		TotalPrice: totalPrice,
	}
}

func (o *Order) GetID() string {
	return o.ID
}

func (o *Order) GetProducts() []string {
	return o.Products
}

func (o *Order) GetTotalPrice() value_objects.Price {
	return o.TotalPrice
}
