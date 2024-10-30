package main

import (
	"errors"
	"log"
	"slices"
	"strconv"
	"strings"
)

var ErrorDoctorNotFound = errors.New("doctor not found")
var ErrorNextDoctorID = errors.New("failed to get next docotor ID")

// Errors for fields validation.
var ErrorInvalidDoctorData = errors.New("invalid doctor data")

type DoctorStore struct {
	data map[DoctorID]Doctor
}

func NewDoctorStore(list []Doctor) *DoctorStore {
	ms := DoctorStore{
		data: make(map[DoctorID]Doctor, len(list)),
	}
	for _, d := range list {
		ms.data[d.ID] = d
	}

	return &ms
}

func (ms *DoctorStore) Print() {
	for _, d := range ms.data {
		log.Printf("print: %v", d)
	}
}

func (ms *DoctorStore) GetDoctor(id DoctorID) (Doctor, error) {
	d, ok := ms.data[id]
	if !ok {
		return Doctor{}, ErrorDoctorNotFound
	}
	return d, nil
}

func (ms *DoctorStore) GetDoctors() ([]Doctor, error) {
	keys := make([]DoctorID, len(ms.data))
	i := 0
	for key := range ms.data {
		keys[i] = key
		i++
	}

	slices.Sort(keys)

	result := make([]Doctor, len(ms.data))
	for i, key := range keys {
		result[i] = ms.data[key]
	}

	return result, nil
}

func (ms *DoctorStore) CreateDoctor(d Doctor) (DoctorID, error) {
	log.Println("DoctorStore.CreateDoctor")
	// if d.ID != "" {
	// 	return "", fmt.Errorf("%w: ID", ErrorInvalidDoctorData)
	// }
	// if d.Name == "" {
	// 	return "", fmt.Errorf("%w: Name", ErrorInvalidDoctorData)
	// }
	// if d.Email == "" {
	// 	return "", fmt.Errorf("%w: Email", ErrorInvalidDoctorData)
	// }
	// if d.Speciality == "" {
	// 	return "", fmt.Errorf("%w: Speciality", ErrorInvalidDoctorData)
	// }

	var err error
	d.ID, err = ms.NextDoctorID()
	if err != nil {
		return "", err
	}

	ms.data[d.ID] = d
	return d.ID, nil
}

func digitizeDoctorID(id DoctorID) int {
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

func (ms *DoctorStore) NextDoctorID() (DoctorID, error) {
	// Make numeric representation of IDs.
	// numIDs := make([]int, len(ms.data))
	maxid := 0
	for key := range ms.data {
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
	return DoctorID(newID), nil
}
