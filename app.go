package main

import (
	"fmt"
	"log"
)

type PatientID string

type Patient struct {
	ID       PatientID `json:"id"`
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	External bool      `json:"external"`
}

type DoctorID string

type Doctor struct {
	ID         DoctorID `json:"id"`
	Name       string   `json:"name"`
	Email      string   `json:"email"`
	Role       string   `json:"role"`
	Speciality string   `json:"speciality"`
}

var DoctorRoles = [...]string{"radiologist", "technician", "nurse", "admin"}
var DoctorSpecialities = []string{"general", "dermatology", "neurology", "cardiology", ""}

type Store interface {
	GetPatient(PatientID) (Patient, error)
	GetPatients() ([]Patient, error)
	CreatePatient(Patient) (PatientID, error)

	GetDoctor(DoctorID) (Doctor, error)
	GetDoctors() ([]Doctor, error)
	CreateDoctor(Doctor) (DoctorID, error)
}

type MedicalApp struct {
	store Store
}

func NewApp(s Store) *MedicalApp {
	return &MedicalApp{
		store: s,
	}
}

func (app *MedicalApp) GetPatient(id PatientID) (Patient, error) {
	return app.store.GetPatient(id)
}

func (app *MedicalApp) GetPatients() ([]Patient, error) {
	return app.store.GetPatients()
}

func (app *MedicalApp) CreatePatient(p Patient) (PatientID, error) {
	return app.store.CreatePatient(p)
}

func (app *MedicalApp) GetDoctor(id DoctorID) (Doctor, error) {
	log.Printf("MedicalApp.GetDoctor %v", id)
	return app.store.GetDoctor(id)
}

func (app *MedicalApp) GetDoctors() ([]Doctor, error) {
	log.Printf("MedicalApp.GetDoctors")
	return app.store.GetDoctors()
}

func (app *MedicalApp) CreateDoctor(d Doctor) (DoctorID, error) {
	log.Printf("MedicalApp.CreateDoctor: %+v", d)
	if d.ID != "" {
		return "", fmt.Errorf("%w: ID=[%v]", ErrorInvalidDoctorData, d.ID)
	}
	if d.Name == "" {
		return "", fmt.Errorf("%w: Name=[%v]", ErrorInvalidDoctorData, d.Name)
	}
	if d.Email == "" {
		return "", fmt.Errorf("%w: Email=[%v]", ErrorInvalidDoctorData, d.Email)
	}
	if d.Role == "" || !ValidDoctorRole(d.Role) {
		return "", fmt.Errorf("%w: Role=[%v]", ErrorInvalidDoctorData, d.Role)
	}
	if d.Speciality != "" && !ValidDoctorSpeciality(d.Speciality) {
		return "", fmt.Errorf("%w: Speciality=[%v]", ErrorInvalidDoctorData, d.Speciality)
	}
	if d.Role == "admin" && d.Speciality != "" {
		return "", fmt.Errorf("%w: Role=admin, Speciality=[%v]", ErrorInvalidDoctorData, d.Speciality)
	}

	log.Printf(">> MedicalApp.CreateDoctor: %+v, call store.CreateDoctor", d)
	return app.store.CreateDoctor(d)
}

func ValidDoctorRole(role string) bool {
	for _, r := range DoctorRoles {
		if role == r {
			return true
		}
	}
	return false
}

func ValidDoctorSpeciality(spec string) bool {
	for _, r := range DoctorSpecialities {
		if spec == r {
			return true
		}
	}
	return false
}
