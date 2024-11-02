package postgres

import (
	"database/sql"
	"log"
	"time"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (ps *Store) Now() (time.Time, error) {
	row := ps.db.QueryRow("select now()")
	var now time.Time
	err := row.Scan(&now)
	if err != nil {
		log.Println("Store.Test", err)
	}
	return now, err
}
