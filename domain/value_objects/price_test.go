package value_objects

import (
	"testing"
)

func TestNewPrice_ValidAmount(t *testing.T) {
	price, err := NewPrice(100.0)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if price.GetAmount() != 100.0 {
		t.Errorf("expected amount to be 100.0, got %v", price.GetAmount())
	}
}

func TestNewPrice_NegativeAmount(t *testing.T) {
	_, err := NewPrice(-10.0)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	expectedError := "price cannot be negative"
	if err.Error() != expectedError {
		t.Errorf("expected error to be '%s', got '%s'", expectedError, err.Error())
	}
}

func TestPrice_Add(t *testing.T) {
	price1, _ := NewPrice(50.0)
	price2, _ := NewPrice(30.0)

	sum := price1.Add(price2)

	if sum.GetAmount() != 80.0 {
		t.Errorf("expected sum to be 80.0, got %v", sum.GetAmount())
	}
}

func TestPrice_GetAmount(t *testing.T) {
	price, _ := NewPrice(200.0)

	amount := price.GetAmount()
	if amount != 200.0 {
		t.Errorf("expected amount to be 200.0, got %v", amount)
	}
}
