package store

import (
	"errors"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

type PatientStore struct {
	data map[string]PatientRecord
}

var ErrorPatientNotFound = errors.New("patient not found")
var ErrorNextPatientID = errors.New("failed to get next patient ID")

// Errors for fields validation.
var ErrorInvalidPatientData = errors.New("invalid patient data")

func NewPatientStore(list []PatientRecord) *PatientStore {
	store := PatientStore{
		data: make(map[string]PatientRecord, len(list)),
	}
	for _, p := range list {
		store.data[p.ID] = p
	}

	return &store
}

func (store *PatientStore) Print() {
	for _, p := range store.data {
		log.Printf("print: %v", p)
	}
}

func (store *PatientStore) GetPatient(id string) (PatientRecord, error) {
	p, ok := store.data[id]
	if !ok {
		return PatientRecord{}, ErrorPatientNotFound
	}
	return p, nil
}

func (store *PatientStore) GetPatients() ([]PatientRecord, error) {
	keys := make([]string, len(store.data))
	i := 0
	for key := range store.data {
		keys[i] = key
		i++
	}

	slices.Sort(keys)

	result := make([]PatientRecord, len(store.data))
	for i, key := range keys {
		result[i] = store.data[key]
	}

	return result, nil
}

func (store *PatientStore) CreatePatient(p PatientRecord) (string, error) {
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
	p.ID, err = store.NextPatientID()
	if err != nil {
		return "", err
	}

	store.data[p.ID] = p
	return p.ID, nil
}

func extractNumberFromID(id string) int {
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

func (store *PatientStore) NextPatientID() (string, error) {
	// Make numeric representation of IDs.
	maxid := 0
	for key := range store.data {
		numID := extractNumberFromID(key)
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
