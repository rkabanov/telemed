package pgstore

import (
	"database/sql"
	"log"
	"time"
)

type PGStore struct {
	db *sql.DB
}

func New(db *sql.DB) *PGStore {
	return &PGStore{
		db: db,
	}
}

func (store *PGStore) Now() (time.Time, error) {
	row := store.db.QueryRow("select now()")
	var now time.Time
	err := row.Scan(&now)
	if err != nil {
		log.Println("PGStore.Test", err)
	}
	return now, err
}
