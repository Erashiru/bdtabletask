package models

import (
	"database/sql"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

const (
	migrate = "./migrate/create.sql"
)

func New(storagePath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, err
	}

	if err = Create(db); err != nil {
		return nil, err
	}

	return db, nil
}

func Create(db *sql.DB) error {
	file, err := os.ReadFile(migrate)
	if err != nil {
		return err
	}
	requests := strings.Split(string(file), ";")
	for _, request := range requests {
		_, err := db.Exec(request)
		if err != nil {
			return err
		}
	}
	return nil
}
