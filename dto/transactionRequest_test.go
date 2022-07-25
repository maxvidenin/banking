package dto

import (
	"net/http"
	"testing"
)

func TestShouldReturnErrorForWrongTransactionTypes(t *testing.T) {
	// Arrange
	req := TransactionRequest{
		TransactionType: "wrong type",
		Amount:          100,
	}
	// Act
	appError := req.Validate()
	// Assert
	if appError.Message != "TransactionType must be either 'deposit' or 'withdrawal'" {
		t.Error("Invalid message while testing transaction type")
	}
	if appError.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code while testing transaction type")
	}

}

func TestShouldReturnErrorForNegativeAmount(t *testing.T) {
	// Arrange
	req := TransactionRequest{
		TransactionType: Deposit,
		Amount:          -100,
	}
	// Act
	appError := req.Validate()
	// Assert
	if appError.Message != "Amount must be greater than 0" {
		t.Error("Invalid message while testing transaction amount less than 0")
	}
	if appError.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code while testing transaction amount less than 0")
	}

}
