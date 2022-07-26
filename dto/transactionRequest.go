package dto

import "github.com/maxvidenin/banking-lib/errs"

const Withdrawal = "withdrawal"
const Deposit = "deposit"

type TransactionRequest struct {
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
	AccountId       string  `json:"account_id"`
	CustomerId      string  `json:"customer_id"`
}

func (req TransactionRequest) IsWithdrawal() bool {
	return req.TransactionType == Withdrawal
}

func (req TransactionRequest) IsDeposit() bool {
	return req.TransactionType == Deposit
}

func (req TransactionRequest) Validate() *errs.AppError {
	if !req.IsWithdrawal() && !req.IsDeposit() {
		return errs.NewValidationError("TransactionType must be either 'deposit' or 'withdrawal'")
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
