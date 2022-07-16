package dto

import (
	"strings"

	"github.com/maxvidenin/banking/errs"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (req NewAccountRequest) Validate() *errs.AppError {
	if req.CustomerId == "" {
		return errs.NewValidationError("CustomerId is required")
	}

	if req.AccountType == "" {
		return errs.NewValidationError("AccountType is required")
	}

	if strings.ToLower(req.AccountType) != "checking" && strings.ToLower(req.AccountType) != "savings" {
		return errs.NewValidationError("AccountType must be either 'checking' or 'savings'")
	}

	if req.Amount <= 5000 {
		return errs.NewValidationError("Amount must be greater than 5000.00")
	}

	return nil
}
