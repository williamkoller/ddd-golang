package repositories

import "ddd-golang/domain/entities"

type OrderRepository interface {
	Save(order *entities.Order) error
	FindByID(id string) (*entities.Order, error)
}
