package main

type PatientID string

type Patient struct {
	ID       PatientID `json:"id"`
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	External bool      `json:"external"`
}

type Store interface {
	GetPatient(PatientID) (*Patient, error)
	GetPatients() ([]Patient, error)
	// CreatePatient(*Patient) (*Patient, error)
}

type PatientApp struct {
	store Store
}

func NewApp(s Store) *PatientApp {
	return &PatientApp{
		store: s,
	}
}

func (app *PatientApp) GetPatient(id PatientID) (*Patient, error) {
	p, err := app.store.GetPatient(id)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (app *PatientApp) GetPatients() ([]Patient, error) {
	res, err := app.store.GetPatients()
	if err != nil {
		return nil, err
	}
	return res, nil
}
