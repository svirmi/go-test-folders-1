package domain

import (
	"folders-one/errs"
	"folders-one/logger"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?)"

	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)

	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from DB")
	}

	id, err := result.LastInsertId()

	if err != nil {
		logger.Error("Error while getting last insert id for a new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from DB")
	}

	a.AccountId = strconv.FormatInt(id, 10)

	return &a, nil
}
