package service

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	realdomain "github.com/maxvidenin/banking/domain"
	"github.com/maxvidenin/banking/dto"
	"github.com/maxvidenin/banking/errs"
	"github.com/maxvidenin/banking/mocks/domain"
)

var mockRepo *domain.MockAccountRepository
var service AccountService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockRepo = domain.NewMockAccountRepository(ctrl)
	service = NewAccountService(mockRepo)
	return func() {
		service = nil
		defer ctrl.Finish()
	}
}

func TestShouldReturnAmountValidationError(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()
	req := dto.NewAccountRequest{
		CustomerId:  "1001",
		AccountType: "savings",
		Amount:      0,
	}
	// Act
	_, appErr := service.NewAccount(req)
	// Assert
	if appErr.Message != "Amount must be greater than 5000.00" {
		t.Error("Expected amount validation error")
	}

}

func TestShouldReturnDbSaveError(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()
	req := dto.NewAccountRequest{
		CustomerId:  "1001",
		AccountType: "savings",
		Amount:      10000.00,
	}
	account := realdomain.Account{
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
	mockRepo.EXPECT().Save(account).Return(nil, errs.NewUnexpectedError("Unxpected database error"))
	// Act
	_, appErr := service.NewAccount(req)
	// // Assert
	if appErr == nil {
		t.Error("Expected db save error")
	}

}

func TestShouldReturnNewAccountRespinse(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()
	req := dto.NewAccountRequest{
		CustomerId:  "1001",
		AccountType: "savings",
		Amount:      10000.00,
	}
	account := realdomain.Account{
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
	accountWithId := account
	accountWithId.AccountId = "2001"
	mockRepo.EXPECT().Save(account).Return(&accountWithId, nil)
	// Act
	newAccount, appErr := service.NewAccount(req)
	// // Assert
	if appErr != nil {
		t.Error("Expected to create new account")
	}

	if newAccount.AccountId != accountWithId.AccountId {
		t.Error("Expected account id to match")
	}
}
