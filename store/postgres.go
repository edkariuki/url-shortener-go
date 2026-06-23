package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	db *sql.DB
}

func New(connStr string) (*Store, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &Store{db: db}, nil
}

func (s *Store) Save(longURL string) (uint64, error) {
	var id uint64
	err := s.db.QueryRow(
		`INSERT INTO urls (original) VALUES ($1) RETURNING id`,
		longURL,
	).Scan(&id)
	return id, err
}

func (s *Store) FindByURL(longURL string) (uint64, error) {
	var id uint64
	err := s.db.QueryRow(
		`SELECT id FROM urls WHERE original = $1`,
		longURL,
	).Scan(&id)
	if err == sql.ErrNoRows {
		return 0, nil
	}
	return id, err
}

func (s *Store) FindByID(id uint64) (string, error) {
	var longURL string
	err := s.db.QueryRow(
		`SELECT original FROM urls WHERE id = $1`,
		id,
	).Scan(&longURL)
	if err == sql.ErrNoRows {
		return "", nil
	}
	return longURL, err
}
