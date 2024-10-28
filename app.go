package main

type PatientID string

type Patient struct {
	ID       PatientID `json:"id"`
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	External bool      `json:"external"`
}

type Store interface {
	GetPatient(PatientID) (Patient, error)
	GetPatients() ([]Patient, error)
	CreatePatient(Patient) (PatientID, error)
}

type PatientApp struct {
	store Store
}

func NewApp(s Store) *PatientApp {
	return &PatientApp{
		store: s,
	}
}

func (app *PatientApp) GetPatient(id PatientID) (Patient, error) {
	return app.store.GetPatient(id)
}

func (app *PatientApp) GetPatients() ([]Patient, error) {
	return app.store.GetPatients()
}

func (app *PatientApp) CreatePatient(p Patient) (PatientID, error) {
	return app.store.CreatePatient(p)
}
