package structs

import (
	"github.com/szmulinho/prescription/pkg/model"
	"time"
)

var Prescs []model.CreatePrescInput

var PreID string

var Expiration time.Time

var Drugs []string

type CreatePrescInput struct {
	PreId      string    `json:"pre-id"`
	Drugs      []string  `json:"drugs"`
	Expiration time.Time `json:"expiration"`
}
