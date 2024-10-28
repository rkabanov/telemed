package main

import (
	"fmt"
	"log"
	"slices"
)

type MemStore struct {
	data map[PatientID]Patient
}

var ErrorPatientNotFound = fmt.Errorf("patient not found")

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

func (ms *MemStore) GetPatient(id PatientID) (*Patient, error) {
	p, ok := ms.data[id]
	if !ok {
		return nil, ErrorPatientNotFound
	}
	log.Printf("MemStore.GetPatient: %v", p)
	return &p, nil
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
