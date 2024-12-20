package memory

import (
	"fmt"
	"log"
	"slices"
	"strconv"

	"github.com/rkabanov/telemed/store"
)

func (ms *Store) GetDoctor(id string) (store.DoctorRecord, error) {
	log.Println("GetDoctor")
	d, ok := ms.doctors[id]
	if !ok {
		return store.DoctorRecord{}, store.ErrorDoctorNotFound
	}
	return d, nil
}

func (ms *Store) GetDoctors() ([]store.DoctorRecord, error) {
	log.Println("GetDoctors: len(store.doctors)=", len(ms.doctors))
	keys := make([]string, len(ms.doctors))
	i := 0
	for key := range ms.doctors {
		keys[i] = key
		i++
	}

	slices.Sort(keys)

	result := make([]store.DoctorRecord, len(ms.doctors))
	for i, key := range keys {
		result[i] = ms.doctors[key]
	}

	return result, nil
}

func (ms *Store) CreateDoctor(d store.DoctorRecord) (string, error) {
	log.Println("Store.CreateDoctor")
	if d.ID != "" {
		return "", fmt.Errorf("%w: ID", store.ErrorInvalidDoctorData)
	}
	if d.Name == "" {
		return "", fmt.Errorf("%w: Name", store.ErrorInvalidDoctorData)
	}
	if d.Email == "" {
		return "", fmt.Errorf("%w: Email", store.ErrorInvalidDoctorData)
	}
	if d.Role == "" {
		return "", fmt.Errorf("%w: Role", store.ErrorInvalidDoctorData)
	}

	var err error
	d.ID, err = ms.NextDoctorID()
	if err != nil {
		return "", err
	}

	ms.doctors[d.ID] = d
	return d.ID, nil
}

func (ms *Store) NextDoctorID() (string, error) {
	// Make numeric representation of IDs.
	maxid := 0
	for key := range ms.doctors {
		numID := store.ExtractNumberFromString(key)
		if numID > maxid {
			maxid = numID
		}
	}

	// make the next ID - increase the number by 1
	maxid++

	// convert the number to string
	newID := strconv.Itoa(maxid)
	log.Println("NextDoctorID: newID=", newID)
	return newID, nil
}
