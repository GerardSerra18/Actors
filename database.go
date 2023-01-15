package main

import (
	"database/sql"
)

// Create tables
func initDB() error {
	db, err := connectToDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE actors (id SERIAL PRIMARY KEY, first_name TEXT, last_name TEXT, gender CHAR(1), age INTEGER)`)
	if err != nil {
		return err
	}

	return nil
}

// Database connection
func connectToDB() (*sql.DB, error) {

	connStr := "user=gerard password=12345 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}