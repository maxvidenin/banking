package domain

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/maxvidenin/banking/errs"
	"github.com/maxvidenin/banking/logger"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (cr CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var findAllSql string
	var err error
	customers := make([]Customer, 0)

	findAllSql = "SELECT customer_id, name, city, zipcode,date_of_birth, status FROM customers"
	if status == "" {
		err = cr.client.Select(&customers, findAllSql)
	} else {
		findAllSql += " WHERE status = ?"
		err = cr.client.Select(&customers, findAllSql, status)
	}

	if err != nil {
		logger.Error(err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	return customers, nil
}

func (cr CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	findByIdSql := "SELECT customer_id, name, city, zipcode,date_of_birth, status FROM customers WHERE customer_id = ?"

	var c Customer
	err := cr.client.Get(&c, findByIdSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			logger.Error(err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}

	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sqlx.Open("mysql", "root:123456@/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}
