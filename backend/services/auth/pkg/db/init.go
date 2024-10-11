package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var Instance *sql.DB

func InitDB() {
	var err error
	connStr := "postgres://user:password@localhost:5432/auth_db?sslmode=disable"
	Instance, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = Instance.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
