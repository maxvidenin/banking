package service

import (
	"github.com/maxvidenin/banking/domain"
	"github.com/maxvidenin/banking/dto"
	"github.com/maxvidenin/banking/errs"
)

type CustomerService interface {
	GetAllCustomers(string) ([]domain.Customer, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerReporsitory
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]domain.Customer, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	return s.repo.FindAll(status)
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}

	response := c.ToDto()

	return &response, nil
}

func NewCustomerService(repo domain.CustomerReporsitory) DefaultCustomerService {
	return DefaultCustomerService{repo}
}
