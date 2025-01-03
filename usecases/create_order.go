package usecases

import (
	"ddd-golang/domain/entities"
	"ddd-golang/domain/repositories"
	"ddd-golang/domain/value_objects"
)

type CreateOrderUseCase struct {
	OrderRepository repositories.OrderRepository
}

func NewCreateOrderUseCase(repo repositories.OrderRepository) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		OrderRepository: repo,
	}
}

func (uc *CreateOrderUseCase) Execute(orderID string, products []string, totalPrice float64) (*entities.Order, error) {
	price, err := value_objects.NewPrice(totalPrice)
	if err != nil {
		return nil, err
	}

	order := entities.NewOrder(orderID, products, price)
	if err := uc.OrderRepository.Save(order); err != nil {
		return nil, err
	}

	return order, nil
}
