package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func InitDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:superpostgres@localhost:5432/userdb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db
}