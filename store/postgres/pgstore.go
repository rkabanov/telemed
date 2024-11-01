package pgstore

import (
	"database/sql"
	"log"
	"time"
)

type PostgresStore struct {
	db *sql.DB
}

func New(db *sql.DB) *PostgresStore {
	return &PostgresStore{
		db: db,
	}
}

func (ps *PostgresStore) Now() (time.Time, error) {
	row := ps.db.QueryRow("select now()")
	var now time.Time
	err := row.Scan(&now)
	if err != nil {
		log.Println("PostgresStore.Test", err)
	}
	return now, err
}
