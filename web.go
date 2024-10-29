package main

type App interface {
	GetPatient(PatientID) (Patient, error)
	GetPatients() ([]Patient, error)
	CreatePatient(Patient) (PatientID, error)

	GetDoctor(DoctorID) (Doctor, error)
	GetDoctors() ([]Doctor, error)
	CreateDoctor(Doctor) (DoctorID, error)
}

type WebAPI struct {
	app App
}

func NewWebAPI(a App) *WebAPI {
	return &WebAPI{
		app: a,
	}
}
