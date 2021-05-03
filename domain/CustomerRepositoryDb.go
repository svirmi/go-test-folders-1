package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type NewCustomerRepositoryDb struct {
}

func (d NewCustomerRepositoryDb) FindAll() ([]Customer, error) {
	client, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	rows, err := client.Query(findAllSql)

	if err != nil {
		log.Println("Error while querying customer table " + err.Error())
		return nil, err
	}

	customers := make([]Customer, 0)

	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)

		if err != nil {
			log.Println("Error while scanning customers " + err.Error())
			return nil, err
		}

		customers = append(customers, c)
	}

	return customers, nil

}
