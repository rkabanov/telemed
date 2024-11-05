package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/rkabanov/service/store"
)

func (ps *Store) GetPatient(id string) (store.PatientRecord, error) {
	log.Println("Store.GetPatient")
	query := "select id, name, age, external from patient where id=$1"
	var p store.PatientRecord
	row := ps.db.QueryRow(query, id)
	err := row.Scan(&p.ID,
		&p.Name,
		&p.Age,
		&p.External,
	)

	if err == sql.ErrNoRows {
		return p, store.ErrorPatientNotFound
	}

	if err != nil {
		return p, fmt.Errorf("postgres.GetPatient error: %w", err)
	}

	return p, nil
}

func (ps *Store) GetPatients() ([]store.PatientRecord, error) {
	log.Println("Store.GetPatients")
	query := "select id, name, age, external from patients order by id"
	rows, err := ps.db.Query(query)
	if err != nil {
		return []store.PatientRecord{}, err
	}
	defer rows.Close()

	var result []store.PatientRecord
	for rows.Next() {
		var p store.PatientRecord
		if err := rows.Scan(&p.ID,
			&p.Name,
			&p.Age,
			&p.External,
		); err != nil {
			return result, err
		}
		result = append(result, p)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (ps *Store) CreatePatient(p store.PatientRecord) (string, error) {
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

	// var err error
	// p.ID, err = ms.NextPatientID()
	// if err != nil {
	// 	return "", err
	// }

	query := "insert into patients(id, name, age, external) values ($1, $2, $3, $4) returning id"
	row := ps.db.QueryRow(query, p.ID, p.Name, p.Age, p.External)
	err := row.Scan(&p.ID)
	if err != nil {
		return "", err
	}

	return p.ID, nil
}

// func digitizePatientID(id string) int {
// 	// filter digits
// 	var sb strings.Builder
// 	for _, r := range string(id) {
// 		if r >= '0' && r <= '9' {
// 			sb.WriteRune(r)
// 		}
// 	}
// 	digits := sb.String()
// 	if len(digits) == 0 {
// 		return 0
// 	}

// 	// convert to number
// 	num, err := strconv.Atoi(digits)
// 	if err != nil {
// 		return 0
// 	}

// 	return num
// }

func (ps *Store) NextPatientID() (string, error) {
	doctors, err := ps.GetPatients()
	if err != nil {
		return "", store.ErrorNextPatientID
	}

	// Make numeric representation of IDs.
	maxid := 0
	for _, d := range doctors {
		numID := store.ExtractNumberFromString(d.ID)
		if numID > maxid {
			maxid = numID
		}
	}

	// make the next ID - increase the number by 1
	maxid++

	// convert the number to string
	newID := strconv.Itoa(maxid)
	log.Println("NextPatientID: newID=", newID)
	return newID, nil
}
