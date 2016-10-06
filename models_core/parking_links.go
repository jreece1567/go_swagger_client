package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*ParkingLinks parking links

swagger:model parkingLinks
*/
type ParkingLinks struct {

	/* URI to pdf ref file

	Required: true
	*/
	Pdf *GenericLink `json:"pdf"`
}

// Validate validates this parking links
func (m *ParkingLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePdf(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ParkingLinks) validatePdf(formats strfmt.Registry) error {

	if err := validate.Required("pdf", "body", m.Pdf); err != nil {
		return err
	}

	if m.Pdf != nil {

		if err := m.Pdf.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
