package memory

import (
	"log"
	"testing"

	"github.com/rkabanov/service/store"
	"github.com/stretchr/testify/require"
)

var testMemStore *Store

func init() {
	testMemStore = NewStore(
		[]store.DoctorRecord{
			{ID: "11", Name: "Dr. John", Email: "john@yopmail.com", Role: "radiologist", Speciality: "neurology"},
			{ID: "22", Name: "Dr. Mary", Email: "mary@yopmail.com", Role: "admin", Speciality: "admin"},
		},
		[]store.PatientRecord{
			{ID: "1", Name: "John", Age: 30, External: false},
			{ID: "2", Name: "Mary", Age: 25, External: true},
		},
	)
}

func TestNextDoctorID(t *testing.T) {
	id, err := testMemStore.NextDoctorID()
	require.NoError(t, err)
	require.NotEmpty(t, id)
	require.Equal(t, id, "23")
}

func TestGetDoctor(t *testing.T) {
	d, err := testMemStore.GetDoctor("11")
	require.NoError(t, err)
	require.NotEmpty(t, d)
	require.Equal(t, "11", d.ID)
	require.Equal(t, "Dr. John", d.Name)
	require.Equal(t, "john@yopmail.com", d.Email)
	require.Equal(t, "radiologist", d.Role)
	require.Equal(t, "neurology", d.Speciality)
}

func TestGetDoctors(t *testing.T) {
	list, err := testMemStore.GetDoctors()
	require.NoError(t, err)
	require.NotEmpty(t, list)
	require.Equal(t, 2, len(list))
	require.Equal(t, "11", list[0].ID)
	require.Equal(t, "22", list[1].ID)
}

func TestCreateDoctor(t *testing.T) {
	d := store.DoctorRecord{
		ID:         "",
		Name:       "Dr. Charles",
		Email:      "charles@yopmail.com",
		Role:       "radiologist",
		Speciality: "neurology",
	}

	newID, err := testMemStore.CreateDoctor(d)
	require.NoError(t, err)
	require.NotEmpty(t, newID)

	var newDoc store.DoctorRecord
	newDoc, err = testMemStore.GetDoctor(newID)
	require.NoError(t, err)
	require.NotEmpty(t, newDoc)
	require.Equal(t, newDoc.ID, newID)
	require.Equal(t, newDoc.Name, d.Name)
	require.Equal(t, newDoc.Email, d.Email)
	require.Equal(t, newDoc.Role, d.Role)
	require.Equal(t, newDoc.Speciality, d.Speciality)

	log.Printf("TestCreateDoctor: docStore: %v", testMemStore)
}

func TestNextPatientID(t *testing.T) {
	id, err := testMemStore.NextPatientID()
	require.NoError(t, err)
	require.NotEmpty(t, id)
	require.Equal(t, id, "3")
}

func TestGetPatient(t *testing.T) {
	p, err := testMemStore.GetPatient("1")
	require.NoError(t, err)
	require.NotEmpty(t, p)
	require.Equal(t, "1", p.ID)
	require.Equal(t, "John", p.Name)
	require.Equal(t, 30, p.Age)
	require.Equal(t, false, p.External)
}

func TestGetPatients(t *testing.T) {
	list, err := testMemStore.GetPatients()
	require.NoError(t, err)
	require.NotEmpty(t, list)
	require.Equal(t, 2, len(list))
	require.Equal(t, "1", list[0].ID)
	require.Equal(t, "2", list[1].ID)
}

func TestCreatePatient(t *testing.T) {
	p := store.PatientRecord{
		ID:       "",
		Name:     "Charles",
		Age:      21,
		External: true,
	}

	newID, err := testMemStore.CreatePatient(p)
	require.NoError(t, err)
	require.NotEmpty(t, newID)

	var newPat store.PatientRecord
	newPat, err = testMemStore.GetPatient(newID)
	require.NoError(t, err)
	require.NotEmpty(t, newPat)
	require.Equal(t, newPat.ID, newID)
	require.Equal(t, newPat.Name, p.Name)
	require.Equal(t, newPat.Age, p.Age)
	require.Equal(t, newPat.External, p.External)

	log.Printf("TestCreatePatient: %v", newPat)
}
