package postgres

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/rkabanov/service/store"
)

func (ps *Store) Print() {
	// for _, d := range ps.data {
	// 	log.Printf("print: %v", d)
	// }
	log.Println("Store.Print - TODO")
}

func (ps *Store) GetDoctor(id string) (store.DoctorRecord, error) {
	query := "select id, name, email, role, speciality from doctors where id=$1 limit 1"
	row := ps.db.QueryRow(query, id)
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

func (ps *Store) GetDoctors() ([]store.DoctorRecord, error) {
	query := "select id, name, email, role, speciality from doctors order by id"

	// arg.Limit, arg.Offset - TODO
	rows, err := ps.db.Query(query)
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

func (ps *Store) CreateDoctor(d store.DoctorRecord) (string, error) {
	log.Println("Store.CreateDoctor")

	var err error
	d.ID, err = ps.NextDoctorID()
	if err != nil {
		return "", err
	}

	query := `insert into doctors(id, name, email, role, speciality)
	values ($1, $2, $3, $4, $5) returning id`
	row := ps.db.QueryRow(query, d.ID, d.Name, d.Email, d.Role, d.Speciality)
	err = row.Scan(&d.ID)
	return d.ID, err
}

func (ps *Store) NextDoctorID() (string, error) {
	doctors, err := ps.GetDoctors()
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
