package routes

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type ClientManager struct {
	db      *sql.DB
	secrets map[string]string
}

func NewClientManager(dbPath string) (*ClientManager, error) {
	if dbPath == "" {
		dbPath = "user.db"
	}

	database, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Initialize database if doesn't exist
	_, err = database.Exec(`CREATE TABLE IF NOT EXISTS Users (
		Id TEXT NOT NULL PRIMARY KEY,
		FirstName TEXT NOT NULL,
		LastName TEXT NOT NULL,
		Email TEXT NOT NULL,
		Username TEXT NOT NULL,
		Password TEXT NOT NULL
	)`)
	if err != nil {
		return nil, err
	}

	return &ClientManager{
		db:      database,
		secrets: make(map[string]string),
	}, nil
}
