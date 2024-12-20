package app

import (
	"log"
	"testing"

	"github.com/rkabanov/telemed/store"
	"github.com/rkabanov/telemed/store/memory"
	"github.com/stretchr/testify/require"
)

var testMemStore *memory.Store

func init() {
	testMemStore = memory.NewStore([]store.DoctorRecord{
		{ID: "11", Name: "Dr. John", Email: "john@yopmail.com", Role: "radiologist", Speciality: "neurology"},
		{ID: "22", Name: "Dr. Mary", Email: "mary@yopmail.com", Role: "admin", Speciality: "admin"},
	},
		[]store.PatientRecord{
			{ID: "1", Name: "John", Age: 30, External: false},
			{ID: "2", Name: "Mary", Age: 25, External: true},
		})
}
func TestApp(t *testing.T) {
	app := NewApp(testMemStore)
	require.NotEmpty(t, app)

	doctors, err := app.GetDoctors()
	log.Println(doctors)
	require.NoError(t, err)
	require.Equal(t, len(doctors), 2)

	patients, err := app.GetPatients()
	require.NoError(t, err)
	require.Equal(t, len(patients), 2)
}
