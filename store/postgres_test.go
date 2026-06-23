package store

import (
	"testing"
)

func TestDBConnection(t *testing.T) {
	connStr := "postgres://postgres:postgres@localhost/shortener?sslmode=disable"

	s, err := New(connStr)
	if err != nil {
		t.Fatalf("could not connect: %v", err)
	}

	if err := s.db.Ping(); err != nil {
		t.Fatalf("ping failed: %v", err)
	}

	t.Log("connected to database")
}
