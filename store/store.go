package store

import (
	"errors"
	"strconv"
	"strings"
)

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

// Store specific errors.
var ErrorNextDoctorID = errors.New("failed to get next docotor ID")
var ErrorNextPatientID = errors.New("failed to get next patient ID")

var ErrorPatientNotFound = errors.New("patient not found")
var ErrorDoctorNotFound = errors.New("doctor not found")

// Errors for fields validation.
var ErrorInvalidPatientData = errors.New("invalid patient data")
var ErrorInvalidDoctorData = errors.New("invalid doctor data")

// extractNumberFromID is an utility function.
func ExtractNumberFromString(id string) int {
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
