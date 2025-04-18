package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewPostgresSQL(conn string) (*sql.DB, error) {
	// Initialize a new PostgresDB instance
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
