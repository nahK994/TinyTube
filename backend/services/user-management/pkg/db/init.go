package db

import (
	"database/sql"
	"fmt"
	"log"
	"user-management/pkg/app"

	_ "github.com/lib/pq"
)

type DB struct {
	db *sql.DB
}

func InitDB(dbConfig app.DBConfig) (*DB, error) {
	var err error
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &DB{
		db: db,
	}, nil
}
