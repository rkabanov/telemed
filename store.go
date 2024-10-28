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

func (ms *MemStore) NextPatientID() (PatientID, error) {
	log.Println("NextPatientID")
	var maxid PatientID = ""
	for key := range ms.data {
		if key > PatientID(maxid) {
			maxid = key
		}
	}
	if maxid == "" {
		maxid = "0"
	}
	log.Println("NextPatientID: maxid=", maxid)

	// filter digits
	var sb strings.Builder
	for i, r := range string(maxid) {
		log.Println("NextPatientID: i=", i, ", rune=", r)
		if r >= '0' && r <= '9' {
			sb.WriteRune(r)
			// digits = append(digits, byte(ch))
		}
	}
	digits := sb.String()
	if len(digits) == 0 {
		digits = "0"
	}
	log.Println("NextPatientID: digits=", string(digits))

	// convert to number
	num, err := strconv.Atoi(digits)
	if err != nil {
		log.Println("NextPatientID: ERROR:", err)
		return "", ErrorNextPatientID
	}
	log.Println("NextPatientID: num=", num)

	// make the next ID - increase the number by 1
	num++

	// convert the number to string
	newID := strconv.Itoa(num)
	log.Println("NextPatientID: newID=", newID)
	return PatientID(newID), nil
}
