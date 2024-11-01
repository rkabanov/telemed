package store

import (
	"errors"
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
