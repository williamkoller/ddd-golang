package repositories

import (
	"ddd-golang/domain/entities"
	"errors"
)

type MockOrderRepository struct {
	SaveFunc     func(order *entities.Order) error
	FindByIDFunc func(id string) (*entities.Order, error)
}

func (m *MockOrderRepository) Save(order *entities.Order) error {
	if m.SaveFunc != nil {
		return m.SaveFunc(order)
	}
	return nil
}

func (m *MockOrderRepository) FindByID(id string) (*entities.Order, error) {
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(id)
	}
	return nil, errors.New("not implemented")
}
