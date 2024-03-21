package models

import "database/sql"

func New(storagePath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS storage (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			stelaj VARCHAR(1) NOT NULL,
			name TEXT NOT NULL,
			request INTEGER NOT NULL,
			pieces INTEGER NOT NULL,
			additional_st VARCHAR(10)
		)`)
	if err != nil {
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
