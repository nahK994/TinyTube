package db

import (
	"auth-service/pkg/app"
	"database/sql"
	"fmt"
	"log"

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
		password VARCHAR(255) NOT NULL,
		profile_pic TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	createTokensTable := `
	CREATE TABLE IF NOT EXISTS tokens (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		refresh_token TEXT NOT NULL,
		expires_at TIMESTAMP NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
	);
	`

	if _, err := db.Exec(createUserTable); err != nil {
		return fmt.Errorf("could not create users table: %w", err)
	}

	if _, err := db.Exec(createTokensTable); err != nil {
		return fmt.Errorf("could not create tokens table: %w", err)
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
