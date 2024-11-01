package store

import (
	"errors"
	"log"
	"slices"
	"strconv"
	"strings"
)

var ErrorDoctorNotFound = errors.New("doctor not found")
var ErrorNextDoctorID = errors.New("failed to get next docotor ID")

type DoctorStore struct {
	data map[string]DoctorRecord
}

func NewDoctorStore(list []DoctorRecord) *DoctorStore {
	store := DoctorStore{
		data: make(map[string]DoctorRecord, len(list)),
	}
	for _, d := range list {
		store.data[d.ID] = d
	}

	return &store
}

func (store *DoctorStore) Print() {
	for _, d := range store.data {
		log.Printf("print: %v", d)
	}
}

func (store *DoctorStore) GetDoctor(id string) (DoctorRecord, error) {
	d, ok := store.data[id]
	if !ok {
		return DoctorRecord{}, ErrorDoctorNotFound
	}
	return d, nil
}

func (store *DoctorStore) GetDoctors() ([]DoctorRecord, error) {
	log.Println("GetDoctors: len(store.data)=", len(store.data))
	keys := make([]string, len(store.data))
	i := 0
	for key := range store.data {
		keys[i] = key
		i++
	}

	slices.Sort(keys)

	result := make([]DoctorRecord, len(store.data))
	for i, key := range keys {
		result[i] = store.data[key]
	}

	return result, nil
}

func (store *DoctorStore) CreateDoctor(d DoctorRecord) (string, error) {
	log.Println("DoctorStore.CreateDoctor")

	var err error
	d.ID, err = store.NextDoctorID()
	if err != nil {
		return "", err
	}

	store.data[d.ID] = d
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

func (store *DoctorStore) NextDoctorID() (string, error) {
	// Make numeric representation of IDs.
	maxid := 0
	for key := range store.data {
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
