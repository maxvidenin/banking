package dto

import "github.com/maxvidenin/banking/errs"

type TransactionRequest struct {
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
	AccountId       string  `json:"account_id"`
	CustomerId      string  `json:"customer_id"`
}

func (req TransactionRequest) Validate() *errs.AppError {
	if req.TransactionType != "deposit" && req.TransactionType != "withdrawal" {
		return errs.NewValidationError("AccountType must be either 'deposit' or 'withdrawal'")
	}
	if req.Amount <= 0 {
		return errs.NewValidationError("Amount must be greater than 0")
	}
	if req.AccountId == "" {
		return errs.NewValidationError("AccountId is required")
	}
	if req.CustomerId == "" {
		return errs.NewValidationError("CustomerId is required")
	}
	return nil
}
