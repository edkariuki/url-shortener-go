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
		`INSERT INTO urls (original_url) VALUES ($1) RETURNING id`,
		longURL,
	).Scan(&id)
	return id, err
}

func (s *Store) FindByCode(code string) (string, error) {
	var longURL string
	err := s.db.QueryRow(
		`SELECT original_url FROM urls WHERE short_code = $1`,
		code,
	).Scan(&longURL)
	if err == sql.ErrNoRows {
		return "", nil
	}
	return longURL, err
}

func (s *Store) UpdateCode(id uint64, code string) error {
	_, err := s.db.Exec(
		`UPDATE urls SET short_code = $1 WHERE id = $2`,
		code, id,
	)
	return err
}
