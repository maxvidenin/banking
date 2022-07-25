package service

import (
	"time"

	"github.com/maxvidenin/banking/domain"
	"github.com/maxvidenin/banking/dto"
	"github.com/maxvidenin/banking/errs"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	MakeTransaction(dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	if account, err := s.repo.Save(domain.NewAccount(
		req.CustomerId,
		req.AccountType,
		req.Amount,
	)); err != nil {
		return nil, err
	} else {
		return account.ToNewAccountResponseDto(), nil
	}
}

func (s DefaultAccountService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	account, err := s.repo.ById(req.AccountId)
	if err != nil {
		return nil, err
	}
	if account.Status != "1" {
		return nil, errs.NewValidationError("Account is closed")
	}
	if req.TransactionType == "withdrawal" && !account.CanWithdraw(req.Amount) {
		return nil, errs.NewValidationError("Insufficient funds")
	}
	nt := domain.Transaction{
		AccountId:       req.AccountId,
		TransactionType: req.TransactionType,
		Amount:          req.Amount,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}
	t, appErr := s.repo.SaveTransaction(nt)
	if appErr != nil {
		return nil, appErr
	}
	response := t.ToDto()
	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repo}
}
