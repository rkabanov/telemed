package postgres

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/rkabanov/telemed/store"
)

func (s *Store) Print() {
	doctors, err := s.GetDoctors()
	if err != nil {
		log.Println("Print: failed to get dostors:", err)
	}
	for i, d := range doctors {
		log.Printf("\tdoctor [%v]: %v\n", i, d)
	}

	patients, err := s.GetPatients()
	if err != nil {
		log.Println("Print: failed to get patients:", err)
	}
	for i, p := range patients {
		log.Printf("\tpatient [%v]: %v\n", i, p)
	}
}

func (s *Store) GetDoctor(id string) (store.DoctorRecord, error) {
	query := "select id, name, email, role, speciality from doctors where id=$1 limit 1"
	row := s.db.QueryRow(query, id)
	var d store.DoctorRecord
	err := row.Scan(
		&d.ID,
		&d.Name,
		&d.Email,
		&d.Role,
		&d.Speciality,
	)

	if err == sql.ErrNoRows {
		return store.DoctorRecord{}, store.ErrorDoctorNotFound
	}

	return d, err
}

func (s *Store) GetDoctors() ([]store.DoctorRecord, error) {
	query := "select id, name, email, role, speciality from doctors order by id"

	// arg.Limit, arg.Offset - TODO
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []store.DoctorRecord
	for rows.Next() {
		var i store.DoctorRecord
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Email,
			&i.Role,
			&i.Speciality,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (s *Store) CreateDoctor(d store.DoctorRecord) (string, error) {
	log.Println("Store.CreateDoctor")

	var err error
	d.ID, err = s.NextDoctorID()
	if err != nil {
		return "", err
	}

	query := `insert into doctors(id, name, email, role, speciality)
	values ($1, $2, $3, $4, $5) returning id`
	row := s.db.QueryRow(query, d.ID, d.Name, d.Email, d.Role, d.Speciality)
	err = row.Scan(&d.ID)
	return d.ID, err
}

func (s *Store) NextDoctorID() (string, error) {
	doctors, err := s.GetDoctors()
	if err != nil {
		return "", store.ErrorNextDoctorID
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
	log.Println("NextDoctorID: newID=", newID)
	return newID, nil
}
