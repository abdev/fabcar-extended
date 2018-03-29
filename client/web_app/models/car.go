package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Car struct {
	ID        uuid.UUID `json:"ID" db:"id"`
	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt time.Time `json:"-" db:"updated_at"`
	Make      string    `json:"Make" db:"make"`
	Model     string    `json:"Model" db:"model"`
	Colour    string    `json:"Colour" db:"colour"`
	UserID    uuid.UUID `json:"-" db:"user_id"`
	Owner     User      `belongs_to:"user"`
}

// String is not required by pop and may be deleted
func (c Car) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Cars is not required by pop and may be deleted
type Cars []Car

// String is not required by pop and may be deleted
func (c Cars) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (c *Car) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: c.Make, Name: "Make"},
		&validators.StringIsPresent{Field: c.Model, Name: "Model"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (c *Car) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (c *Car) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
