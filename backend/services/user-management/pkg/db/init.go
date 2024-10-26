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

func createTables(db *sql.DB) error {
	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		profile_pic TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	if _, err := db.Exec(createUserTable); err != nil {
		return fmt.Errorf("could not create users table: %w", err)
	}

	return nil
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
	if err := createTables(db); err != nil {
		return nil, err
	}

	return &DB{
		db: db,
	}, nil
}
