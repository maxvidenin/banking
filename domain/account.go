package domain

import (
	"github.com/maxvidenin/banking/dto"
	"github.com/maxvidenin/banking/errs"
)

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		AccountId:   a.AccountId,
		CustomerId:  a.CustomerId,
		AccountType: a.AccountType,
		Amount:      a.Amount,
	}
}
