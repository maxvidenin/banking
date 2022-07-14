package domain

import "github.com/maxvidenin/banking/errs"

type Customer struct {
	Id          int64  `json:"id db:"customer_id"`
	Name        string `json:"name"` 
	City        string `json:"city"` 
	Zipcode     string `json:"zipcode"`
	DateOfBirth string `json:"date_of_birth"`
	Status      string `json:"status"`
}

type CustomerReporsitory interface {
	FindAll(string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
