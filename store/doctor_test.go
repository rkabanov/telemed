package store

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

var testDoctorStore *DoctorStore

func init() {
	testDoctorStore = NewDoctorStore([]DoctorRecord{
		{ID: "11", Name: "Dr. John", Email: "john@yopmail.com", Role: "radiologist", Speciality: "neurology"},
		{ID: "22", Name: "Dr. Mary", Email: "mary@yopmail.com", Role: "admin", Speciality: "admin"},
	})
}

func TestNextDoctorID(t *testing.T) {
	id, err := testDoctorStore.NextDoctorID()
	require.NoError(t, err)
	require.NotEmpty(t, id)
	require.Equal(t, id, "23")
}

func TestGetDoctor(t *testing.T) {
	d, err := testDoctorStore.GetDoctor("11")
	require.NoError(t, err)
	require.NotEmpty(t, d)
	require.Equal(t, "11", d.ID)
	require.Equal(t, "Dr. John", d.Name)
	require.Equal(t, "john@yopmail.com", d.Email)
	require.Equal(t, "radiologist", d.Role)
	require.Equal(t, "neurology", d.Speciality)
}

func TestGetDoctors(t *testing.T) {
	list, err := testDoctorStore.GetDoctors()
	require.NoError(t, err)
	require.NotEmpty(t, list)
	require.Equal(t, 2, len(list))
	require.Equal(t, "11", list[0].ID)
	require.Equal(t, "22", list[1].ID)
}

func TestCreateDoctor(t *testing.T) {
	d := DoctorRecord{
		ID:         "",
		Name:       "Dr. Charles",
		Email:      "charles@yopmail.com",
		Role:       "radiologist",
		Speciality: "neurology",
	}

	newID, err := testDoctorStore.CreateDoctor(d)
	require.NoError(t, err)
	require.NotEmpty(t, newID)

	var newDoc DoctorRecord
	newDoc, err = testDoctorStore.GetDoctor(newID)
	require.NoError(t, err)
	require.NotEmpty(t, newDoc)
	require.Equal(t, newDoc.ID, newID)
	require.Equal(t, newDoc.Name, d.Name)
	require.Equal(t, newDoc.Email, d.Email)
	require.Equal(t, newDoc.Role, d.Role)
	require.Equal(t, newDoc.Speciality, d.Speciality)

	log.Printf("TestCreateDoctor: docStore: %v", testDoctorStore)
}
