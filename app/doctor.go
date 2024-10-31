package app

import (
	"errors"
	"fmt"
	"log"

	"github.com/rkabanov/service/store"
)

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

// Application level errors.
var ErrorDoctorNotFound = errors.New("doctor not found")
var ErrorInvalidDoctorData = errors.New("invalid doctor data")

func (app *MedicalApp) GetDoctor(id DoctorID) (Doctor, error) {
	log.Printf("MedicalApp.GetDoctor %v", id)
	d, err := app.store.GetDoctor(string(id))
	if err != nil {
		return Doctor{}, fmt.Errorf("MedicalApp.GetDoctor failed: %w", err)
	}

	return Doctor{
		ID:         DoctorID(d.ID),
		Name:       d.Name,
		Email:      d.Email,
		Role:       d.Role,
		Speciality: d.Speciality,
	}, nil
}

func (app *MedicalApp) GetDoctors() ([]Doctor, error) {
	log.Printf("MedicalApp.GetDoctors")
	list, err := app.store.GetDoctors()
	if err != nil {
		return []Doctor{}, fmt.Errorf("MedicalApp.GetDoctors failed: %w", err)
	}

	result := make([]Doctor, len(list))
	for i, d := range list {
		result[i] = Doctor{
			ID:         DoctorID(d.ID),
			Name:       d.Name,
			Email:      d.Email,
			Role:       d.Role,
			Speciality: d.Speciality,
		}
	}

	return result, nil
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
	id, err := app.store.CreateDoctor(store.DoctorRecord{
		Name:       d.Name,
		Email:      d.Email,
		Role:       d.Role,
		Speciality: d.Speciality,
	})
	if err != nil {
		return "", fmt.Errorf("MedicalApp.CreateDoctor failed: %w", err)
	}

	return DoctorID(id), nil
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
