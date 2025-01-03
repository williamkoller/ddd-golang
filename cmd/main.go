package main

import (
	"ddd-golang/domain/repositories"
	"ddd-golang/usecases"
	"fmt"
)

func Run() error {
	repo := repositories.NewInMemoryOrderRepository()
	createOrderUseCase := usecases.NewCreateOrderUseCase(repo)

	orderID := "1234"
	products := []string{"Product A", "Product B"}
	totalPrice := 100.0

	order, err := createOrderUseCase.Execute(orderID, products, totalPrice)
	if err != nil {
		return fmt.Errorf("error creating order: %w", err)
	}

	fmt.Printf("Order created with ID: %s\n", order.GetID())
	fmt.Printf("Products: %v\n", order.GetProducts())
	fmt.Printf("Total price: %.2f\n", order.GetTotalPrice().GetAmount())
	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
