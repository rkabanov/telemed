package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/lib/pq"
	"github.com/rkabanov/service/store"

	"github.com/stretchr/testify/require"
)

var (
	driver     = "postgres"
	source     = "postgresql://root:secret@localhost:5433/servicedb?sslmode=disable"
	testStore  *Store
	testDoctor store.DoctorRecord
)

func init() {
	testDB, err := sql.Open(driver, source)
	if err != nil {
		log.Fatal("failed top open DB connection")
	}

	testStore = NewStore(testDB)

	testDoctor = store.DoctorRecord{
		ID:         "",
		Name:       "Markus Davidson",
		Email:      "dm@yopmail.com",
		Role:       "radiologist",
		Speciality: "cardiology",
	}
}

func TestNow(t *testing.T) {
	fmt.Println("TestNow")
	now, err := testStore.Now()
	require.NoError(t, err)
	log.Println("log NOW:", now)
}

func TestCreateDoctor(t *testing.T) {
	fmt.Println("TestCreateDoctor")
	var err error
	testDoctor.ID, err = testStore.CreateDoctor(testDoctor)
	require.NoError(t, err)
	require.NotEmpty(t, testDoctor.ID)
}

func TestGetDoctor(t *testing.T) {
	fmt.Println("TestGetDoctor")
	d, err := testStore.GetDoctor(testDoctor.ID)
	require.NoError(t, err)
	require.NotEmpty(t, d)
	require.Equal(t, testDoctor.ID, d.ID)
	require.Equal(t, testDoctor.Name, d.Name)
	require.Equal(t, testDoctor.Email, d.Email)
	require.Equal(t, testDoctor.Role, d.Role)
	require.Equal(t, testDoctor.Speciality, d.Speciality)
}
