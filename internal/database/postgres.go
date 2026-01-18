package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// handle database connection

func Connect() (*sql.DB, error) {
	connStr := "user=postgres password=Golang2026! dbname=postgres port=5431 sslmode=disable"
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
