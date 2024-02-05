package db

import (
	"database/sql"
	"tg-backend/internal/db/sqlc"

	_ "github.com/lib/pq"
)

type Store struct {
	*sqlc.Queries
	db *sql.DB
}

func New(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: sqlc.New(db),
	}
}
