package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var db *sql.DB

func initializeDb() error {
	var err error

	db, err = sql.Open("postgres", "postgres://user:pass@127.0.0.1:5432/house_db?sslmode=disable")
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	return nil
}

func storeInDb(data temperatureReading) error {
	_, err := db.Exec("insert into data (id, name) values (1, 'a')")
	if err != nil {
		return err
	}
	return nil
}
