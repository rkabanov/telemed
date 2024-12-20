package memory

import (
	// "errors"
	"log"

	"github.com/rkabanov/telemed/store"
)

type Store struct {
	doctors  map[string]store.DoctorRecord
	patients map[string]store.PatientRecord
}

func NewStore(doctors []store.DoctorRecord, patients []store.PatientRecord) *Store {
	store := Store{
		doctors:  make(map[string]store.DoctorRecord, len(doctors)),
		patients: make(map[string]store.PatientRecord, len(patients)),
	}
	for _, d := range doctors {
		store.doctors[d.ID] = d
	}
	for _, p := range patients {
		store.patients[p.ID] = p
	}
	return &store

}

func (store *Store) Print() {
	for _, p := range store.patients {
		log.Printf("Store.patients: %v", p)
	}
	for _, d := range store.doctors {
		log.Printf("Store.doctors: %v", d)
	}
}

// // extractNumberFromID is an utility function.
// func extractNumberFromID(id string) int {
// 	// filter digits
// 	var sb strings.Builder
// 	for _, r := range string(id) {
// 		if r >= '0' && r <= '9' {
// 			sb.WriteRune(r)
// 		}
// 	}
// 	digits := sb.String()
// 	if len(digits) == 0 {
// 		return 0
// 	}

// 	// convert to number
// 	num, err := strconv.Atoi(digits)
// 	if err != nil {
// 		return 0
// 	}

// 	return num
// }
