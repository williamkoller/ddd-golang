package main

import (
	"ddd-golang/domain/repositories"
	"ddd-golang/usecases"
	"fmt"
)

func main() {
	repo := repositories.NewInMemoryOrderRepository()
	createOrderUseCase := usecases.NewCreateOrderUseCase(repo)

	orderID := "1234"
	products := []string{"Product A", "Product B"}
	totalPrice := 100.0

	order, err := createOrderUseCase.Execute(orderID, products, totalPrice)
	if err != nil {
		fmt.Printf("Error creating order: %v\n", err)
		return
	}

	fmt.Printf("Order created with ID: %s\n", order.GetID())
	fmt.Printf("Products: %v\n", order.GetProducts())
	fmt.Printf("Total price: %v\n", order.GetTotalPrice())

}
