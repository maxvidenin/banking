package domain

import "github.com/maxvidenin/banking/dto"

type Transaction struct {
	TransactionId   string  `db:"transaction_id"`
	TransactionType string  `db:"transaction_type"`
	Amount          float64 `db:"amount"`
	AccountId       string  `db:"account_id"`
	TransactionDate string  `db:"transaction_date"`
}

func (t Transaction) ToDto() dto.TransactionResponse {
	return dto.TransactionResponse{
		TransactionId:   t.TransactionId,
		AccountId:       t.AccountId,
		Amount:          t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}
}
