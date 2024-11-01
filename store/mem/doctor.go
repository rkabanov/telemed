package mem

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/rkabanov/service/store"
)

func (ms *MemStore) GetDoctor(id string) (store.DoctorRecord, error) {
	d, ok := ms.doctors[id]
	if !ok {
		return store.DoctorRecord{}, store.ErrorDoctorNotFound
	}
	return d, nil
}

func (ms *MemStore) GetDoctors() ([]store.DoctorRecord, error) {
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

func (ms *MemStore) CreateDoctor(d store.DoctorRecord) (string, error) {
	log.Println("MemStore.CreateDoctor")
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

func digitizeDoctorID(id string) int {
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

func (ms *MemStore) NextDoctorID() (string, error) {
	// Make numeric representation of IDs.
	maxid := 0
	for key := range ms.doctors {
		numID := digitizeDoctorID(key)
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
