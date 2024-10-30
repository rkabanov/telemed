package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var docStore *DoctorStore
var patStore *PatientStore

func init() {
	docStore = NewDoctorStore([]Doctor{
		{ID: "301", Name: "Dr. Gibbs", Email: "gibbs@yopmail.com", Role: "radiologist", Speciality: "cardiology"},
		{ID: "301", Name: "Dr. Evans", Email: "evans@yopmail.com", Role: "admin", Speciality: ""},
	})

	patStore = NewPatientStore([]Patient{
		{ID: "30001", Name: "Nick", Age: 53, External: false},
		{ID: "30002", Name: "Albert", Age: 42, External: false},
		{ID: "30003", Name: "William", Age: 38, External: true},
	})
}

func TestApp(t *testing.T) {
	app := NewAppStore(patStore, docStore)
	require.NotEmpty(t, app)

	doctors, err := app.GetDoctors()
	require.NoError(t, err)
	require.Len(t, doctors, len(docStore.data))

	patients, err := app.GetPatients()
	require.NoError(t, err)
	require.Len(t, patients, len(patStore.data))
}
