package domain

import (
	"database/sql"
	"time"

	"github.com/maxvidenin/banking/errs"
	"github.com/maxvidenin/banking/logger"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (cr CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var findAllSql string
	var rows *sql.Rows
	var err error

	findAllSql = "SELECT customer_id, name, city, zipcode,date_of_birth, status FROM customers"
	if status == "" {
		rows, err = cr.client.Query(findAllSql)
	} else {
		findAllSql += " WHERE status = ?"
		rows, err = cr.client.Query(findAllSql, status)
	}

	if err != nil {
		logger.Error(err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if err != nil {
			logger.Error(err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
		customers = append(customers, c)
	}

	return customers, nil
}

func (cr CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	findByIdSql := "SELECT customer_id, name, city, zipcode,date_of_birth, status FROM customers WHERE customer_id = ?"

	row := cr.client.QueryRow(findByIdSql, id)

	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)

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
	client, err := sql.Open("mysql", "root:123456@/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}
