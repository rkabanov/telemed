package web

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/rkabanov/service/app"
)

func (wa *WebAPI) HandlePatient(w http.ResponseWriter, r *http.Request) {
	log.Println("HandlePatient: GET or POST?")
	switch r.Method {
	case http.MethodGet:
		log.Println("HandlePatient GET")
		wa.GetPatient(w, r)
	case http.MethodPost:
		log.Println("HandlePatient POST")
		wa.CreatePatient(w, r)
	case http.MethodPut:
		// Update an existing record.
	case http.MethodDelete:
		// Remove the record.
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (wa *WebAPI) GetPatient(w http.ResponseWriter, r *http.Request) {
	var id app.PatientID = app.PatientID(r.FormValue("id"))
	p, err := wa.app.GetPatient(id)
	if err == app.ErrorPatientNotFound {
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
	var p app.Patient
	var err error

	// Here we check only that args have correct types.
	p.ID = app.PatientID(r.FormValue("id"))
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
		if errors.Is(err, app.ErrorInvalidPatientData) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("CreatePatient: %v", p)

	j, err := json.Marshal(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(j)
}
