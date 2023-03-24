package model

import (
	"time"
)

var Prescs []CreatePrescInput

var Expiration string

var Drugs []string

type CreatePrescInput struct {
	PreId      string    `json:"pre-id"`
	Drugs      []string  `json:"drugs"`
	Expiration time.Time `json:"expiration"`
}

var PrescriptionID string

var P CreatePrescInput
