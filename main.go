package main

import (
	"log"
	"net/http"

	"github.com/rkabanov/service/app"
	"github.com/rkabanov/service/store"
	"github.com/rkabanov/service/web"
)

var buildDate string

func main() {
	log.Println("start, build:", buildDate)

	patStore := store.NewPatientStore([]store.PatientRecord{
		{ID: "1001", Name: "Evelynn Lang", Age: 20, External: false},
		{ID: "1002", Name: "Wells Maldonado", Age: 21, External: true},
		{ID: "1003", Name: "Elaina Davis", Age: 22, External: false},
		{ID: "1004", Name: "Lucas Wright", Age: 23, External: true},
		{ID: "1005", Name: "Lily Contreras", Age: 24, External: false},
	})
	patStore.Print()

	docStore := store.NewDoctorStore([]store.DoctorRecord{
		{ID: "9001", Name: "Dr. Paul", Email: "paul@yopmail.com", Role: "radiologist", Speciality: "dermatology"},
		{ID: "9002", Name: "Dr. Smith", Email: "smith@yopmail.com", Role: "admin", Speciality: ""},
		{ID: "9003", Name: "Dr. Tucker", Email: "tucker@yopmail.com", Role: "nurse", Speciality: ""},
	})
	docStore.Print()

	store := store.NewAppStore(patStore, docStore) // Common store - depends on individual stores.

	app := app.NewApp(store) // Business logic - depends on common store (or on all individual stores?).

	web := web.NewWebAPI(app) // Web API for the app - depends on business logic.

	http.HandleFunc("/patient", web.HandlePatient) // GET and POST
	http.HandleFunc("/patients", web.GetPatients)
	http.HandleFunc("/doctor", web.HandleDoctor) // GET and POST
	http.HandleFunc("/doctors", web.GetDoctors)
	log.Println(http.ListenAndServe(":8180", nil))

	log.Println("finish")
}
