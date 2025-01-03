package repositories

import (
	"ddd-golang/domain/entities"
	"ddd-golang/domain/value_objects"
	"testing"
)

func TestInMemoryOrderRepository_SaveAndFindByID(t *testing.T) {
	repo := NewInMemoryOrderRepository()

	price, err := value_objects.NewPrice(150.0)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	order := entities.NewOrder("1234", []string{"Product1", "Product2"}, price)

	// Test Save
	err = repo.Save(order)
	if err != nil {
		t.Fatalf("expected no error while saving order, got %v", err)
	}

	// Test FindByID (Existing Order)
	foundOrder, err := repo.FindByID("1234")
	if err != nil {
		t.Fatalf("expected no error while finding order, got %v", err)
	}

	if foundOrder.GetID() != order.GetID() {
		t.Errorf("expected order ID to be '1234', got %v", foundOrder.GetID())
	}

	if len(foundOrder.GetProducts()) != 2 {
		t.Errorf("expected 2 products, got %v", len(foundOrder.GetProducts()))
	}

	if foundOrder.GetTotalPrice() != price {
		t.Errorf("expected total price to be %v, got %v", price, foundOrder.GetTotalPrice())
	}

	// Test FindByID (Non-Existing Order)
	_, err = repo.FindByID("5678")
	if err == nil || err.Error() != "order not found" {
		t.Errorf("expected error 'order not found', got %v", err)
	}
}
