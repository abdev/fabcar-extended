package actions

import (
	"fmt"

	"github.com/abdev/fabcar-extended/web_app/models"
	"github.com/gobuffalo/buffalo"
	"github.com/pkg/errors"
)

// UploadFile default implementation.
func UploadFile(c buffalo.Context) error {
	c.Set("energy", &models.Energy{})
	return c.Render(200, r.HTML("energy_data/upload_file.html"))
}

// ProcessUploadFile handles the csv file processing
func ProcessUploadFile(c buffalo.Context) error {
	energy := &models.Energy{}

	if err := c.Bind(energy); err != nil {
		return errors.WithStack(err)
	}

	fmt.Println("The energy data is", energy.EnergyFile)

	/*content, err := energy.ReadFile()
	if err != nil {
		return errors.WithStack(err)
	}

	fmt.Println("The file content is", string(content))*/
	fmt.Println("The options are", energy.Option["test1"])

	if err := energy.SaveFile(); err != nil {
		return errors.WithStack(err)
	}

	c.Set("energy", energy)

	// If there are no errors set a success message
	c.Flash().Add("success", "The file was uploaded successfully")

	return c.Redirect(302, "/energy_data/uploadFile")

	//return c.Render(201, r.HTML("energy_data/upload_file.html"))
}
