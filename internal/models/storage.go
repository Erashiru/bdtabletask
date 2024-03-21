package models

import "database/sql"

type Storage struct {
	id            int
	stelaj        string
	name          string
	request       int
	pieces        int
	additional_st string
}

type StorageDB struct {
	DB *sql.DB
}

func (m *StorageDB) Insert(stelaj, name string, request, pieces int, additional_st string) (int, error) {
	stmt := `
		INSERT INTO storage (stelaj, name, request, pieces, additional_st)
		VALUES(?, ?, ?, ?, ?)
	`

	result, err := m.DB.Exec(stmt, stelaj, name, request, pieces, additional_st)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), err
}

func (m *StorageDB) Get(id int) (*Storage, error) {
	stmt := `SELECT id, stelaj, name, request, pieces, additional_st
	FROM storage WHERE id = ?
	`

	row := m.DB.QueryRow(stmt, id)

	s := &Storage{}

	err := row.Scan(&s.id, &s.stelaj, &s.name, &s.request, &s.pieces, &s.additional_st)
	if err != nil {
		return nil, err
	}

	return s, nil
}
