package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

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

	return db, err
}
