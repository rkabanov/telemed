package pgstore

import (
	"log"
	"strconv"
	"strings"

	"github.com/rkabanov/service/store"
)

// var ErrorDoctorNotFound = errors.New("doctor not found")
// var ErrorNextDoctorID = errors.New("failed to get next docotor ID")

func (ps *PostgresStore) Print() {
	// for _, d := range ps.data {
	// 	log.Printf("print: %v", d)
	// }
	log.Println("PostgresStore.Print")
}

func (ps *PostgresStore) GetDoctor(id string) (store.DoctorRecord, error) {
	query := "select id, name, email, role, speciality from doctors where id=$1 limit 1"
	row := ps.db.QueryRow(query)
	var d store.DoctorRecord
	err := row.Scan(
		&d.ID,
		&d.Name,
		&d.Email,
		&d.Role,
		&d.Speciality,
	)
	return d, err
}

func (ps *PostgresStore) GetDoctors() ([]store.DoctorRecord, error) {
	query := "select id, name, email, role, speciality from doctors where id=$1 order by id"

	// arg.Limit, arg.Offset
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

func (ps *PostgresStore) CreateDoctor(d store.DoctorRecord) (string, error) {
	log.Println("PostgresStore.CreateDoctor")

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

func (ps *PostgresStore) NextDoctorID() (string, error) {
	doctors, err := ps.GetDoctors()
	if err != nil {
		return "", store.ErrorNextDoctorID
	}

	// Make numeric representation of IDs.
	maxid := 0
	for _, d := range doctors {
		numID := digitizeDoctorID(d.ID)
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
