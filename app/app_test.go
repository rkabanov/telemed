package app

import (
	"log"
	"testing"

	"github.com/rkabanov/service/store"
	"github.com/stretchr/testify/require"
)

var docStore *store.DoctorStore
var patStore *store.PatientStore

func init() {
	docStore = store.NewDoctorStore([]store.DoctorRecord{
		{ID: "301", Name: "Dr. Gibbs", Email: "gibbs@yopmail.com", Role: "radiologist", Speciality: "cardiology"},
		{ID: "302", Name: "Dr. Evans", Email: "evans@yopmail.com", Role: "admin", Speciality: ""},
	})

	patStore = store.NewPatientStore([]store.PatientRecord{
		{ID: "30001", Name: "Nick", Age: 53, External: false},
		{ID: "30002", Name: "Albert", Age: 42, External: false},
		{ID: "30003", Name: "William", Age: 38, External: true},
	})
}

func TestApp(t *testing.T) {
	app := store.NewAppStore(patStore, docStore)
	require.NotEmpty(t, app)

	doctors, err := app.GetDoctors()
	log.Println(doctors)
	require.NoError(t, err)
	require.Equal(t, len(doctors), 2)

	patients, err := app.GetPatients()
	require.NoError(t, err)
	require.Equal(t, len(patients), 3)
}
