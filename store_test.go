package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var testMemStore *MemStore

// type Patient struct {
// 	ID       PatientID `json:"id"`
// 	Name     string    `json:"name"`
// 	Age      int       `json:"age"`
// 	External bool      `json:"external"`
// }

func init() {
	testMemStore = NewMemStore([]Patient{
		{ID: "1", Name: "John", Age: 30, External: false},
		{ID: "2", Name: "Mary", Age: 25, External: true},
	})
}

func TestGetPatient(t *testing.T) {
	p, err := testMemStore.GetPatient("1")
	require.NoError(t, err)
	require.NotEmpty(t, p)
	require.Equal(t, PatientID("1"), p.ID)
	require.Equal(t, "John", p.Name)
	require.Equal(t, 30, p.Age)
	require.Equal(t, false, p.External)
}

func TestGetPatients(t *testing.T) {
	list, err := testMemStore.GetPatients()
	require.NoError(t, err)
	require.NotEmpty(t, list)
	require.Equal(t, 2, len(list))
	require.Equal(t, PatientID("1"), list[0].ID)
	require.Equal(t, PatientID("2"), list[1].ID)
}
