package store

type PatientRecord struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	External bool   `json:"external"`
}

type DoctorRecord struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Role       string `json:"role"`
	Speciality string `json:"speciality"`
}
