package postgres

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

var testDB *sql.DB

func init() {
	var driver = "postgres"
	var source = "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable"

	var err error
	testDB, err = sql.Open(driver, source)
	if err != nil {
		log.Fatal("failed top open DB connection")
	}
}

func TestNow(t *testing.T) {
	store := NewStore(testDB)
	now, err := store.Now()
	require.NoError(t, err)

	log.Println("log NOW:", now)
}
