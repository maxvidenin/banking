package domain

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/maxvidenin/banking/errs"
	"github.com/maxvidenin/banking/logger"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (ar AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	// sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES (:customer_id, :opening_date, :account_type, :amount, :status)"
	// res, err := ar.client.NamedExec(sqlInsert, a)

	sqlInsert2 := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES (?, ?, ?, ?, ?)"
	res, err := ar.client.Exec(sqlInsert2, a.AccountId, a.OpeningDate, a.AccountType, a.Amount, a.Status)

	if err != nil {
		logger.Error(err.Error())
		return nil, errs.NewUnexpectedError(err.Error())
	}
	id, err := res.LastInsertId()
	if err != nil {
		logger.Error(err.Error())
		return nil, errs.NewUnexpectedError(err.Error())
	}
	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func NewAccountRepositoryDb(client *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{client: client}
}
