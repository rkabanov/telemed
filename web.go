package main

import (
	"encoding/json"
	"net/http"
)

type App interface {
	GetPatient(id PatientID) (*Patient, error)
	GetPatients() ([]Patient, error)
}

type WebAPI struct {
	app App
}

func NewWebAPI(a App) *WebAPI {
	return &WebAPI{
		app: a,
	}
}

func (wa *WebAPI) GetPatient(w http.ResponseWriter, r *http.Request) {
	var id PatientID
	id = PatientID(r.FormValue("id"))
	p, err := wa.app.GetPatient(id)
	if err == ErrorPatientNotFound {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	j, err := json.Marshal(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// result := fmt.Sprintf("PATIENT: %v", p)
	w.Write(j)
}

func (wa *WebAPI) GetPatients(w http.ResponseWriter, r *http.Request) {
	list, err := wa.app.GetPatients()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	j, err := json.Marshal(list)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// result := fmt.Sprintf("LIST: %v", list)
	w.Write(j)
}
