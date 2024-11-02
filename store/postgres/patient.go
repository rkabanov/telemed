package postgres

import (
	"log"
	"strconv"
	"strings"

	"github.com/rkabanov/service/store"
)

func (ps *Store) GetPatient(id string) (store.PatientRecord, error) {
	log.Println("Store.GetPatient")
	var p store.PatientRecord
	return p, nil
}

func (ps *Store) GetPatients() ([]store.PatientRecord, error) {
	log.Println("Store.GetPatients")
	var result []store.PatientRecord
	return result, nil
}

func (ps *Store) CreatePatient(p store.PatientRecord) (string, error) {
	log.Println("Store.CreatePatient")
	// if p.ID != "" {
	// 	return "", fmt.Errorf("%w: ID", store.ErrorInvalidPatientData)
	// }
	// if p.Name == "" {
	// 	return "", fmt.Errorf("%w: Name", store.ErrorInvalidPatientData)
	// }
	// if p.Age <= 0 {
	// 	return "", fmt.Errorf("%w: Age", store.ErrorInvalidPatientData)
	// }

	// var err error
	// p.ID, err = ms.NextPatientID()
	// if err != nil {
	// 	return "", err
	// }

	return p.ID, nil
}

func digitizePatientID(id string) int {
	// filter digits
	var sb strings.Builder
	for _, r := range string(id) {
		if r >= '0' && r <= '9' {
			sb.WriteRune(r)
		}
	}
	digits := sb.String()
	if len(digits) == 0 {
		return 0
	}

	// convert to number
	num, err := strconv.Atoi(digits)
	if err != nil {
		return 0
	}

	return num
}

func (ps *Store) NextPatientID() (string, error) {
	doctors, err := ps.GetPatients()
	if err != nil {
		return "", store.ErrorNextPatientID
	}

	// Make numeric representation of IDs.
	maxid := 0
	for _, d := range doctors {
		numID := digitizePatientID(d.ID)
		if numID > maxid {
			maxid = numID
		}
	}

	// make the next ID - increase the number by 1
	maxid++

	// convert the number to string
	newID := strconv.Itoa(maxid)
	log.Println("NextPatientID: newID=", newID)
	return newID, nil
}
