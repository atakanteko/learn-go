package domain

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (receiver CustomerRepositoryDb) FindAll() ([]Customer, error) {

	//define sql
	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	//define query
	rows, err := receiver.client.Query(findAllSql)
	if err != nil {
		log.Println("Error while querying customer table" + err.Error())
		return nil, err
	}

	//iterate all the row that received from the db
	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if err != nil {
			log.Println("Error while querying customer table" + err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:12345678@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client: client}
}
