package entities

import (
	"ddd-golang/domain/value_objects"
	"testing"
)

func TestNewOrder(t *testing.T) {
	price, err := value_objects.NewPrice(100.0)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	products := []string{"Product1", "Product2"}
	order := NewOrder("1234", products, price)

	if order.GetID() != "1234" {
		t.Errorf("expected ID to be '1234', got %v", order.GetID())
	}

	if len(order.GetProducts()) != 2 {
		t.Errorf("expected 2 products, got %v", len(order.GetProducts()))
	}

	if order.GetProducts()[0] != "Product1" || order.GetProducts()[1] != "Product2" {
		t.Errorf("expected products to be ['Product1', 'Product2'], got %v", order.GetProducts())
	}

	if order.GetTotalPrice() != price {
		t.Errorf("expected total price to be %v, got %v", price, order.GetTotalPrice())
	}
}

func TestOrderWithInvalidPrice(t *testing.T) {
	_, err := value_objects.NewPrice(-10.0)
	if err == nil {
		t.Fatal("expected error for negative price, got none")
	}
}
