package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type CarLog struct {
	ID            uuid.UUID `json:"id" db:"id"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
	CarID         uuid.UUID `json:"car_id" db:"car_id"`
	TransactionID string    `json:"transaction_id" db:"transaction_id"`
	BlockID       string    `json:"block_id" db:"block_id"`
	Payload       string    `json:"payload" db:"payload"`
	Data          string    `json:"datum" db:"datum"`
	SubmittedOn   time.Time `json:"submitted_on" db:"submitted_on"`
	UpdatedOn     time.Time `json:"updated_on" db:"updated_on"`
}

// String is not required by pop and may be deleted
func (c CarLog) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// CarLogs is not required by pop and may be deleted
type CarLogs []CarLog

// String is not required by pop and may be deleted
func (c CarLogs) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (c *CarLog) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: c.TransactionID, Name: "TransactionID"},
		&validators.StringIsPresent{Field: c.BlockID, Name: "BlockID"},
		&validators.StringIsPresent{Field: c.Payload, Name: "Payload"},
		&validators.StringIsPresent{Field: c.Data, Name: "Data"},
		&validators.TimeIsPresent{Field: c.SubmittedOn, Name: "SubmittedOn"},
		&validators.TimeIsPresent{Field: c.UpdatedOn, Name: "UpdatedOn"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (c *CarLog) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (c *CarLog) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
