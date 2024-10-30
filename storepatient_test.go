package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

var testPatientStore *PatientStore

// type Patient struct {
// 	ID       PatientID `json:"id"`
// 	Name     string    `json:"name"`
// 	Age      int       `json:"age"`
// 	External bool      `json:"external"`
// }

func init() {
	testPatientStore = NewPatientStore([]Patient{
		{ID: "1", Name: "John", Age: 30, External: false},
		{ID: "2", Name: "Mary", Age: 25, External: true},
	})
}

func TestNextPatientID(t *testing.T) {
	id, err := testPatientStore.NextPatientID()
	require.NoError(t, err)
	require.NotEmpty(t, id)
	require.Equal(t, id, PatientID("3"))
}

func TestGetPatient(t *testing.T) {
	p, err := testPatientStore.GetPatient("1")
	require.NoError(t, err)
	require.NotEmpty(t, p)
	require.Equal(t, PatientID("1"), p.ID)
	require.Equal(t, "John", p.Name)
	require.Equal(t, 30, p.Age)
	require.Equal(t, false, p.External)
}

func TestGetPatients(t *testing.T) {
	list, err := testPatientStore.GetPatients()
	require.NoError(t, err)
	require.NotEmpty(t, list)
	require.Equal(t, 2, len(list))
	require.Equal(t, PatientID("1"), list[0].ID)
	require.Equal(t, PatientID("2"), list[1].ID)
}

func TestCreatePatient(t *testing.T) {
	p := Patient{
		ID:       "",
		Name:     "Charles",
		Age:      21,
		External: true,
	}

	newID, err := testPatientStore.CreatePatient(p)
	require.NoError(t, err)
	require.NotEmpty(t, newID)

	var newP Patient
	newP, err = testPatientStore.GetPatient(newID)
	require.NoError(t, err)
	require.NotEmpty(t, newP)
	require.Equal(t, newP.ID, newID)
	require.Equal(t, newP.Name, p.Name)
	require.Equal(t, newP.Age, p.Age)
	require.Equal(t, newP.External, p.External)

	log.Printf("TestCreatePatient: %v", newP)
}
