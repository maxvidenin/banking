package domain

import (
	"github.com/maxvidenin/banking/dto"
	"github.com/maxvidenin/banking/errs"
)

type Account struct {
	AccountId   string  `db:"account_id"`
	CustomerId  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	SaveTransaction(Transaction) (*Transaction, *errs.AppError)
	ById(string) (*Account, *errs.AppError)
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		AccountId:   a.AccountId,
		CustomerId:  a.CustomerId,
		AccountType: a.AccountType,
		Amount:      a.Amount,
	}
}

func (a Account) CanWithdraw(amount float64) bool {
	return a.Amount >= amount
}
