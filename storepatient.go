package main

import (
	"errors"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

type PatientStore struct {
	data map[PatientID]Patient
}

var ErrorPatientNotFound = errors.New("patient not found")
var ErrorNextPatientID = errors.New("failed to get next patient ID")

// Errors for fields validation.
var ErrorInvalidPatientData = errors.New("invalid patient data")

func NewPatientStore(list []Patient) *PatientStore {
	ms := PatientStore{
		data: make(map[PatientID]Patient, len(list)),
	}
	for _, p := range list {
		ms.data[p.ID] = p
	}

	return &ms
}

func (ms *PatientStore) Print() {
	for _, p := range ms.data {
		log.Printf("print: %v", p)
	}
}

func (ms *PatientStore) GetPatient(id PatientID) (Patient, error) {
	p, ok := ms.data[id]
	if !ok {
		return Patient{}, ErrorPatientNotFound
	}
	return p, nil
}

func (ms *PatientStore) GetPatients() ([]Patient, error) {
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

func (ms *PatientStore) CreatePatient(p Patient) (PatientID, error) {
	log.Println("PatientStore.CreatePatient")
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

func (ms *PatientStore) NextPatientID() (PatientID, error) {
	// Make numeric representation of IDs.
	maxid := 0
	for key := range ms.data {
		numID := extractNumberFromID(key)
		if numID > maxid {
			maxid = numID
		}
	}

	// make the next ID - increase the number by 1
	maxid++

	// convert the number to string
	newID := strconv.Itoa(maxid)
	return PatientID(newID), nil
}
