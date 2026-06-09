package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func connectDB() *sql.DB {
	connStr := "postgres://postgres:postgres@localhost:5432/shortener?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("failed to open db:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("failed to connect to db:", err)
	}

	return db
}
