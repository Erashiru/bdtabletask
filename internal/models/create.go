package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func New(storagePath string) (*sql.DB, error) {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS storage_A (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			request INTEGER NOT NULL,
			pieces INTEGER NOT NULL,
			additional_st VARCHAR(10)
		)`)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS storage_B (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			request INTEGER NOT NULL,
			pieces INTEGER NOT NULL,
			additional_st VARCHAR(10)
		)`)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS storage_J (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			request INTEGER NOT NULL,
			pieces INTEGER NOT NULL,
			additional_st VARCHAR(10)
		)`)
	if err != nil {
		return nil, err
	}

	return db, nil
}
