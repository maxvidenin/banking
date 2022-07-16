package service

import (
	"time"

	"github.com/maxvidenin/banking/domain"
	"github.com/maxvidenin/banking/dto"
	"github.com/maxvidenin/banking/errs"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	account, err := s.repo.Save(domain.Account{
		CustomerId:  req.CustomerId,
		AccountType: req.AccountType,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		Amount:      req.Amount,
		Status:      "1",
	})
	if err != nil {
		return nil, err
	}

	response := account.ToNewAccountResponseDto()

	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repo}
}
