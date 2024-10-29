package main

import "fmt"

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
	return app.store.GetDoctor(id)
}

func (app *MedicalApp) GetDoctors() ([]Doctor, error) {
	return app.store.GetDoctors()
}

func (app *MedicalApp) CreateDoctor(d Doctor) (DoctorID, error) {
	if d.ID != "" {
		return "", fmt.Errorf("%w: ID", ErrorInvalidDoctorData)
	}
	if d.Name == "" {
		return "", fmt.Errorf("%w: Name", ErrorInvalidDoctorData)
	}
	if d.Email == "" {
		return "", fmt.Errorf("%w: Email", ErrorInvalidDoctorData)
	}
	if d.Role == "" || !ValidDoctorRole(d.Role) {
		return "", fmt.Errorf("%w: Role", ErrorInvalidDoctorData)
	}
	if d.Speciality != "" && !ValidDoctorSpeciality(d.Speciality) {
		return "", fmt.Errorf("%w: Speciality", ErrorInvalidDoctorData)
	}
	if d.Role == "admin" && d.Speciality != "" {
		return "", fmt.Errorf("%w: Role, Speciality", ErrorInvalidDoctorData)
	}

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
