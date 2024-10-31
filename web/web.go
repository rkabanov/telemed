package web

import (
	"github.com/rkabanov/service/app"
)

type App interface {
	GetPatient(app.PatientID) (app.Patient, error)
	GetPatients() ([]app.Patient, error)
	CreatePatient(app.Patient) (app.PatientID, error)

	GetDoctor(app.DoctorID) (app.Doctor, error)
	GetDoctors() ([]app.Doctor, error)
	CreateDoctor(app.Doctor) (app.DoctorID, error)
}

type WebAPI struct {
	app App
}

func NewWebAPI(a App) *WebAPI {
	return &WebAPI{
		app: a,
	}
}
