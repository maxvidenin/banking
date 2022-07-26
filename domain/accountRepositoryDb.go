package domain

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/maxvidenin/banking-lib/errs"
	"github.com/maxvidenin/banking-lib/logger"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (ar AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	// sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES (:customer_id, :opening_date, :account_type, :amount, :status)"
	// res, err := ar.client.NamedExec(sqlInsert, a)

	sqlInsert2 := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES (?, ?, ?, ?, ?)"
	res, err := ar.client.Exec(sqlInsert2, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)

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

func (ar AccountRepositoryDb) ById(accountId string) (*Account, *errs.AppError) {
	sqlSelect := "SELECT * FROM accounts WHERE account_id = ?"
	a := Account{}
	err := ar.client.Get(&a, sqlSelect, accountId)
	if err != nil {
		logger.Error(err.Error())
		return nil, errs.NewUnexpectedError(err.Error())
	}
	return &a, nil
}

func (ar AccountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	tx, err := ar.client.Beginx()
	if err != nil {
		logger.Error(err.Error())
		return nil, errs.NewUnexpectedError(err.Error())
	}
	sqlInsert := "INSERT INTO transactions (account_id, transaction_type, amount) VALUES	(:account_id, :transaction_type, :amount)"
	res, _ := tx.NamedExec(sqlInsert, t)

	var sqlUpdate string
	if t.TransactionType == "withdrawal" {
		sqlUpdate = "UPDATE accounts SET amount = amount - :amount WHERE account_id = :account_id"
	} else {
		sqlUpdate = "UPDATE accounts SET amount = amount + :amount WHERE account_id = :account_id"
	}
	_, err = tx.NamedExec(sqlUpdate, t)
	if err != nil {
		tx.Rollback()
		logger.Error(err.Error())
		return nil, errs.NewUnexpectedError(err.Error())
	}
	err = tx.Commit()
	if err != nil {
		logger.Error(err.Error())
		return nil, errs.NewUnexpectedError(err.Error())
	}

	transactionId, err := res.LastInsertId()
	if err != nil {
		logger.Error(err.Error())
		return nil, errs.NewUnexpectedError(err.Error())
	}
	t.TransactionId = strconv.FormatInt(transactionId, 10)

	return &t, nil
}

func NewAccountRepositoryDb(client *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{client: client}
}
