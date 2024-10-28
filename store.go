package main

import (
	"errors"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

type MemStore struct {
	data map[PatientID]Patient
}

var ErrorPatientNotFound = errors.New("patient not found")
var ErrorNextPatientID = errors.New("failed to get next patient ID")

// Errors for fields validation.
var ErrorInvalidPatientData = errors.New("invalid patient data")

func NewMemStore(list []Patient) *MemStore {
	ms := MemStore{
		data: make(map[PatientID]Patient, len(list)),
	}
	for _, p := range list {
		log.Printf("insert %v", p)
		ms.data[p.ID] = p
	}

	return &ms
}

func (ms *MemStore) Print() {
	for _, p := range ms.data {
		log.Printf("print: %v", p)
	}
}

func (ms *MemStore) GetPatient(id PatientID) (Patient, error) {
	p, ok := ms.data[id]
	if !ok {
		return Patient{}, ErrorPatientNotFound
	}
	log.Printf("MemStore.GetPatient: %v", p)
	return p, nil
}

func (ms *MemStore) GetPatients() ([]Patient, error) {
	keys := make([]PatientID, len(ms.data))
	i := 0
	for key := range ms.data {
		keys[i] = key
		i++
	}

	slices.Sort(keys)

	result := make([]Patient, len(ms.data))
	for i, key := range keys {
		result[i] = ms.data[key]
	}

	return result, nil
}

func (ms *MemStore) CreatePatient(p Patient) (PatientID, error) {
	if p.ID != "" {
		return "", fmt.Errorf("%w: ID", ErrorInvalidPatientData)
	}
	if p.Name == "" {
		return "", fmt.Errorf("%w: Name", ErrorInvalidPatientData)
	}
	if p.Age <= 0 {
		return "", fmt.Errorf("%w: Age", ErrorInvalidPatientData)
	}

	var err error
	p.ID, err = ms.NextPatientID()
	if err != nil {
		return "", err
	}

	ms.data[p.ID] = p
	return p.ID, nil
}

func extractNumberFromID(id PatientID) int {
	// filter digits
	var sb strings.Builder
	for i, r := range string(id) {
		log.Println("extractNumberFromID: i=", i, ", rune=", r)
		if r >= '0' && r <= '9' {
			sb.WriteRune(r)
		}
	}
	digits := sb.String()
	if len(digits) == 0 {
		return 0
	}
	log.Println("extractNumberFromID: digits=", string(digits))

	// convert to number
	num, err := strconv.Atoi(digits)
	if err != nil {
		log.Println("NextPatientID: ERROR:", err)
		return 0
	}

	return num
}

func (ms *MemStore) NextPatientID() (PatientID, error) {
	log.Println("NextPatientID")
	// Make numeric representation of IDs.
	// numIDs := make([]int, len(ms.data))
	maxid := 0
	for key := range ms.data {
		numID := extractNumberFromID(key)
		if numID > maxid {
			maxid = numID
		}
	}
	log.Println("NextPatientID: maxid=", maxid)

	// make the next ID - increase the number by 1
	maxid++

	// convert the number to string
	newID := strconv.Itoa(maxid)
	log.Println("NextPatientID: newID=", newID)
	return PatientID(newID), nil
}
