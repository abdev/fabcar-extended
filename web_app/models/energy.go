package models

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/gobuffalo/buffalo/binding"
	"github.com/gobuffalo/uuid"
	"github.com/pkg/errors"
)

type Energy struct {
	ID         uuid.UUID         `json:"ID" db:"id"`
	CreatedAt  time.Time         `json:"-" db:"created_at"`
	UpdatedAt  time.Time         `json:"-" db:"updated_at"`
	Name       string            `json:"Name" db:"name"`
	Option     map[string]string `json:"Options" form:"Option"`
	EnergyFile binding.File      `db:"-" form:"EnergyFile"`
}

func (e *Energy) ReadFile() ([]byte, error) {
	content, err := ioutil.ReadAll(e.EnergyFile.File)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return content, nil
}

func (e *Energy) SaveFile() error {
	fmt.Println("We are in after create", e.EnergyFile.Valid())

	if !(e.EnergyFile.Valid()) {
		return nil
	}

	dir := filepath.Join("../", "docs")

	fmt.Println("the dir is", dir)

	// create dir if not exists
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.Mkdir(dir, 0755); err != nil {
			return errors.WithStack(err)
		}
	}

	f, err := os.Create(filepath.Join(dir, e.EnergyFile.Filename))

	if err != nil {
		return errors.WithStack(err)
	}

	defer f.Close()

	_, err = io.Copy(f, e.EnergyFile)

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
