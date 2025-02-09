package memory

import (
	"fmt"
	"log"
	"slices"
	"strconv"

	"github.com/rkabanov/telemed/store"
)

func (ms *Store) GetPatient(id string) (store.PatientRecord, error) {
	p, ok := ms.patients[id]
	if !ok {
		return store.PatientRecord{}, store.ErrorPatientNotFound
	}
	return p, nil
}

func (ms *Store) GetPatients() ([]store.PatientRecord, error) {
	keys := make([]string, len(ms.patients))
	i := 0
	for key := range ms.patients {
		keys[i] = key
		i++
	}

	slices.Sort(keys)

	result := make([]store.PatientRecord, len(ms.patients))
	for i, key := range keys {
		result[i] = ms.patients[key]
	}

	return result, nil
}

func (ms *Store) CreatePatient(p store.PatientRecord) (string, error) {
	log.Println("Store.CreatePatient")
	if p.ID != "" {
		return "", fmt.Errorf("%w: ID", store.ErrorInvalidPatientData)
	}
	if p.Name == "" {
		return "", fmt.Errorf("%w: Name", store.ErrorInvalidPatientData)
	}
	if p.Age <= 0 {
		return "", fmt.Errorf("%w: Age", store.ErrorInvalidPatientData)
	}

	var err error
	p.ID, err = ms.NextPatientID()
	if err != nil {
		return "", err
	}

	ms.patients[p.ID] = p
	return p.ID, nil
}

func (ms *Store) NextPatientID() (string, error) {
	// Make numeric representation of IDs.
	maxid := 0
	for key := range ms.patients {
		numID := store.ExtractNumberFromString(key)
		if numID > maxid {
			maxid = numID
		}
	}

	// make the next ID - increase the number by 1
	maxid++

	// convert the number to string
	newID := strconv.Itoa(maxid)
	return newID, nil
}
