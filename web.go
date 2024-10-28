package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

type App interface {
	GetPatient(PatientID) (Patient, error)
	GetPatients() ([]Patient, error)
	CreatePatient(Patient) (PatientID, error)
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
	var id PatientID = PatientID(r.FormValue("id"))
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

func (wa *WebAPI) CreatePatient(w http.ResponseWriter, r *http.Request) {
	var p Patient
	var err error

	// Here we check only that args have correct types.
	p.ID = PatientID(r.FormValue("id"))
	p.Name = r.FormValue("name")

	p.Age, err = strconv.Atoi(r.FormValue("age"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p.External, err = strconv.ParseBool(r.FormValue("external"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p.ID, err = wa.app.CreatePatient(p)
	if err != nil {
		if errors.Is(err, ErrorInvalidPatientData) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	j, err := json.Marshal(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(j)
}
