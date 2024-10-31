package web

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/rkabanov/service/app"
)

func (wa *WebAPI) HandleDoctor(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleDoctor: GET or POST?")
	switch r.Method {
	case http.MethodGet:
		log.Println("HandleDoctor GET")
		wa.GetDoctor(w, r)
	case http.MethodPost:
		log.Println("HandleDoctor POST")
		wa.CreateDoctor(w, r)
	case http.MethodPut:
		// Update an existing record.
	case http.MethodDelete:
		// Remove the record.
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (wa *WebAPI) GetDoctor(w http.ResponseWriter, r *http.Request) {
	var id app.DoctorID = app.DoctorID(r.FormValue("id"))
	d, err := wa.app.GetDoctor(id)
	if err == app.ErrorDoctorNotFound {
		log.Printf("ERROR: %v", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err != nil {
		log.Printf("ERROR: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	j, err := json.Marshal(d)
	if err != nil {
		log.Printf("ERROR: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(j)
}

func (wa *WebAPI) GetDoctors(w http.ResponseWriter, r *http.Request) {
	list, err := wa.app.GetDoctors()
	if err != nil {
		log.Printf("ERROR: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	j, err := json.Marshal(list)
	if err != nil {
		log.Printf("ERROR: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(j)
}

func (wa *WebAPI) CreateDoctor(w http.ResponseWriter, r *http.Request) {
	log.Printf("WebAPI.CreateDoctor")
	var d app.Doctor
	var err error

	// Here we check only that args have correct types.
	d.ID = app.DoctorID(r.FormValue("id"))
	d.Name = r.FormValue("name")
	d.Email = r.FormValue("email")
	d.Role = r.FormValue("role")
	d.Speciality = r.FormValue("speciality")

	log.Printf("WebAPI.CreateDoctor: call app.CreateDoctor, d=%v", d)
	d.ID, err = wa.app.CreateDoctor(d) // store method is called!
	if err != nil {
		var status int
		if errors.Is(err, app.ErrorInvalidDoctorData) {
			status = http.StatusBadRequest
		} else {
			status = http.StatusInternalServerError
		}

		log.Printf("ERROR: %v", err)
		w.WriteHeader(status)
		return
	}
	log.Printf("WebAPI.CreateDoctor: result: %v", d)

	j, err := json.Marshal(d)
	if err != nil {
		log.Printf("ERROR: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(j)
}
