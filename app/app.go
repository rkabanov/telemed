package app

import (
	"github.com/rkabanov/service/store"
)

type Store interface {
	GetPatient(string) (store.PatientRecord, error)
	GetPatients() ([]store.PatientRecord, error)
	CreatePatient(store.PatientRecord) (string, error)

	GetDoctor(string) (store.DoctorRecord, error)
	GetDoctors() ([]store.DoctorRecord, error)
	CreateDoctor(store.DoctorRecord) (string, error)

	Print()
}

type MedicalApp struct {
	store Store
}

func NewApp(s Store) *MedicalApp {
	return &MedicalApp{
		store: s,
	}
}
