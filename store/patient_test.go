package store

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

var testPatientStore *PatientStore

func init() {
	testPatientStore = NewPatientStore([]PatientRecord{
		{ID: "1", Name: "John", Age: 30, External: false},
		{ID: "2", Name: "Mary", Age: 25, External: true},
	})
}

func TestNextPatientID(t *testing.T) {
	id, err := testPatientStore.NextPatientID()
	require.NoError(t, err)
	require.NotEmpty(t, id)
	require.Equal(t, id, "3")
}

func TestGetPatient(t *testing.T) {
	p, err := testPatientStore.GetPatient("1")
	require.NoError(t, err)
	require.NotEmpty(t, p)
	require.Equal(t, "1", p.ID)
	require.Equal(t, "John", p.Name)
	require.Equal(t, 30, p.Age)
	require.Equal(t, false, p.External)
}

func TestGetPatients(t *testing.T) {
	list, err := testPatientStore.GetPatients()
	require.NoError(t, err)
	require.NotEmpty(t, list)
	require.Equal(t, 2, len(list))
	require.Equal(t, "1", list[0].ID)
	require.Equal(t, "2", list[1].ID)
}

func TestCreatePatient(t *testing.T) {
	p := PatientRecord{
		ID:       "",
		Name:     "Charles",
		Age:      21,
		External: true,
	}

	newID, err := testPatientStore.CreatePatient(p)
	require.NoError(t, err)
	require.NotEmpty(t, newID)

	var newPat PatientRecord
	newPat, err = testPatientStore.GetPatient(newID)
	require.NoError(t, err)
	require.NotEmpty(t, newPat)
	require.Equal(t, newPat.ID, newID)
	require.Equal(t, newPat.Name, p.Name)
	require.Equal(t, newPat.Age, p.Age)
	require.Equal(t, newPat.External, p.External)

	log.Printf("TestCreatePatient: %v", newPat)
}
