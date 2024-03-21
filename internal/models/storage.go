package models

import "database/sql"

type Storage struct {
	id            int
	name          string
	request       int
	pieces        int
	additional_st string
}

type StorageDB struct {
	DB *sql.DB
}

func (m *StorageDB) Insert(name string, request, pieces int, additional_st string) (int, error) {
	stmt := `
		INSERT INTO storage (name, request, pieces, additional_st)
		VALUES(?, ?, ?, ?)
	`

	result, err := m.DB.Exec(stmt, name, request, pieces, additional_st)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), err
}

func (m *StorageDB) Get(request int) (*Storage, error) {
	stmt := `SELECT id, name, request, pieces, additional_st
	FROM storage_A WHERE request = ?
	`

	row := m.DB.QueryRow(stmt, request)

	s := &Storage{}

	err := row.Scan(&s.id, &s.name, &s.request, &s.pieces, &s.additional_st)
	if err != nil {
		return nil, err
	}

	return s, nil
}
