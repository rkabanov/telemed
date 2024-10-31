package app

import (
	"errors"
	"fmt"

	"github.com/rkabanov/service/store"
)

type PatientID string

type Patient struct {
	ID       PatientID `json:"id"`
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	External bool      `json:"external"`
}

// Application level errors.
var ErrorInvalidPatientData = errors.New("invalid patient data")
var ErrorPatientNotFound = errors.New("patient not found")

func (app *MedicalApp) GetPatient(id PatientID) (Patient, error) {
	p, err := app.store.GetPatient(string(id))
	if err != nil {
		return Patient{}, fmt.Errorf("MedicalApp.GetPatient failed: %w", err)
	}

	pat := Patient{
		ID:       PatientID(p.ID),
		Name:     p.Name,
		Age:      p.Age,
		External: p.External,
	}
	return pat, nil
}

func (app *MedicalApp) GetPatients() ([]Patient, error) {
	list, err := app.store.GetPatients()
	if err != nil {
		return []Patient{}, fmt.Errorf("MedicalApp.GetPatients failed: %w", err)
	}

	result := make([]Patient, len(list))
	for i, p := range list {
		result[i] = Patient{
			ID:       PatientID(p.ID),
			Name:     p.Name,
			Age:      p.Age,
			External: p.External,
		}
	}

	return result, nil
}

func (app *MedicalApp) CreatePatient(p Patient) (PatientID, error) {
	id, err := app.store.CreatePatient(store.PatientRecord{
		Name:     string(p.Name),
		Age:      p.Age,
		External: p.External,
	})
	if err != nil {
		return PatientID(""), fmt.Errorf("MedicalApp.CreatePatient failed: %w", err)
	}

	return PatientID(id), nil
}
