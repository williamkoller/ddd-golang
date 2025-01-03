package repositories

import (
	"ddd-golang/domain/entities"
	"errors"
)

type InMemoryOrderRepository struct {
	data map[string]*entities.Order
}

func NewInMemoryOrderRepository() *InMemoryOrderRepository {
	return &InMemoryOrderRepository{
		data: make(map[string]*entities.Order),
	}
}

func (r *InMemoryOrderRepository) Save(order *entities.Order) error {
	r.data[order.ID] = order
	return nil
}

func (r *InMemoryOrderRepository) FindByID(id string) (*entities.Order, error) {
	order, exists := r.data[id]
	if !exists {
		return nil, errors.New("order not found")
	}
	return order, nil
}
