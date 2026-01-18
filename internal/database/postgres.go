package database

import (
	"database/sql"
	"fmt"
)

// handle database connection

func Connect() (*sql.DB, error) {
	connStr := "user=postgres password=Golang2026! dbname=todo-database port=5431"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return db, err
	}

	err = db.Ping()
	if err != nil {
		return db, err
	}

	fmt.Println("Successfully connected to todo-database container")

	return db, err
}
