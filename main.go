package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("start")

	store := NewMemStore([]Patient{
		{ID: "1001", Name: "Ann", Age: 20, External: false},
		{ID: "1002", Name: "Bob", Age: 21, External: true},
		{ID: "1003", Name: "Chris", Age: 22, External: false},
		{ID: "1004", Name: "Donna", Age: 23, External: true},
		{ID: "1005", Name: "Emma", Age: 24, External: false},
	})
	store.Print()

	app := NewApp(store)

	web := NewWebAPI(app)

	http.HandleFunc("/patient", web.GetPatient)
	http.HandleFunc("/patients", web.GetPatients)
	log.Println(http.ListenAndServe(":8180", nil))

	log.Println("finish")
}
