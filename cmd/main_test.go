package main

import (
	"ddd-golang/domain/repositories"
	"ddd-golang/usecases"
	"testing"
)

func TestMainFunction(t *testing.T) {
	repo := repositories.NewInMemoryOrderRepository()
	createOrderUseCase := usecases.NewCreateOrderUseCase(repo)

	orderID := "1234"
	products := []string{"Product A", "Product B"}
	totalPrice := 100.0

	order, err := createOrderUseCase.Execute(orderID, products, totalPrice)
	if err != nil {
		t.Fatalf("Error creating order: %v", err)
	}

	if order.GetID() != orderID {
		t.Errorf("Expected order ID %s, got %s", orderID, order.GetID())
	}

	if len(order.GetProducts()) != len(products) {
		t.Errorf("Expected products %v, got %v", products, order.GetProducts())
	}

	if order.GetTotalPrice().GetAmount() != totalPrice {
		t.Errorf("Expected total price %.2f, got %.2f", totalPrice, order.GetTotalPrice().GetAmount())
	}
}
