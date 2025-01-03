package usecases

import (
	"ddd-golang/domain/entities"
	"ddd-golang/domain/repositories"
	"errors"
	"testing"
)

func TestCreateOrderUseCase_Success(t *testing.T) {
	repo := repositories.NewInMemoryOrderRepository()
	useCase := NewCreateOrderUseCase(repo)

	orderID := "1234"
	products := []string{"product1", "product2"}
	totalPrice := 150.0

	order, err := useCase.Execute(orderID, products, totalPrice)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if order.GetID() != orderID {
		t.Errorf("expected order ID to be '%s', got '%s'", orderID, order.GetID())
	}

	if len(order.GetProducts()) != len(products) {
		t.Errorf("expected %d products, got %d", len(products), len(order.GetProducts()))
	}

	if order.GetTotalPrice().GetAmount() != totalPrice {
		t.Errorf("expected total price to be %.2f, got %.2f", totalPrice, order.GetTotalPrice().GetAmount())
	}
}

func TestCreateOrderUseCase_InvalidPrice(t *testing.T) {
	repo := repositories.NewInMemoryOrderRepository()
	useCase := NewCreateOrderUseCase(repo)

	orderID := "1234"
	products := []string{"product1", "product2"}
	totalPrice := -50.0 // Invalid price

	_, err := useCase.Execute(orderID, products, totalPrice)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	expectedError := "price cannot be negative"
	if err.Error() != expectedError {
		t.Errorf("expected error '%s', got '%s'", expectedError, err.Error())
	}
}

func TestCreateOrderUseCase_RepositoryError(t *testing.T) {
	repo := &repositories.MockOrderRepository{
		SaveFunc: func(order *entities.Order) error {
			return errors.New("repository error")
		},
	}

	useCase := NewCreateOrderUseCase(repo)

	orderID := "1234"
	products := []string{"product1", "product2"}
	totalPrice := 150.0

	_, err := useCase.Execute(orderID, products, totalPrice)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	expectedError := "repository error"
	if err.Error() != expectedError {
		t.Errorf("expected error '%s', got '%s'", expectedError, err.Error())
	}
}
