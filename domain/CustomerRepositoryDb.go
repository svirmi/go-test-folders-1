package domain

import (
	"database/sql"
	"folders-one/errs"
	"folders-one/logger"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {

	var err error

	customers := make([]Customer, 0)

	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findAllSql)

	} else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = d.client.Select(&customers, findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while querying customer table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected DB error")
	}

	return customers, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sqlx.Open("mysql", "root:secret@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client}
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	var c Customer
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	err := d.client.Get(&c, customerSql, id)

	if err != nil {

		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning customer: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected DB error")
		}
	}

	return &c, nil
}
