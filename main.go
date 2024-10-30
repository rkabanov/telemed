package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("start")

	patStore := NewPatientStore([]Patient{
		{ID: "1001", Name: "Ann", Age: 20, External: false},
		{ID: "1002", Name: "Bob", Age: 21, External: true},
		{ID: "1003", Name: "Chris", Age: 22, External: false},
		{ID: "1004", Name: "Donna", Age: 23, External: true},
		{ID: "1005", Name: "Emma", Age: 24, External: false},
	})
	patStore.Print()

	docStore := NewDoctorStore([]Doctor{
		{ID: "9001", Name: "Dr. Paul", Email: "paul@yopmail.com", Role: "radiologist", Speciality: "dermatology"},
		{ID: "9002", Name: "Dr. Smith", Email: "smith@yopmail.com", Role: "admin", Speciality: ""},
		{ID: "9003", Name: "Dr. Tucker", Email: "tucker@yopmail.com", Role: "nurse", Speciality: ""},
	})
	docStore.Print()

	store := NewAppStore(patStore, docStore) // Common store - depends on individual stores.

	app := NewApp(store) // Business logic - depends on common store (or on all individual stores?).

	web := NewWebAPI(app) // Web API for the app - depends on business logic.

	http.HandleFunc("/patient", web.HandlePatient) // GET and POST
	http.HandleFunc("/patients", web.GetPatients)
	http.HandleFunc("/doctor", web.HandleDoctor) // GET and POST
	http.HandleFunc("/doctors", web.GetDoctors)
	log.Println(http.ListenAndServe(":8180", nil))

	log.Println("finish")
}
