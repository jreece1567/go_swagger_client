package models_secure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*ParkingSignup422Response Response for invalid data on parking signup.

swagger:model parkingSignup422Response
*/
type ParkingSignup422Response struct {

	/* data

	Required: true
	*/
	Data EmptyObject `json:"data"`

	/* errors

	Required: true
	*/
	Errors *ParkingSignup422Error `json:"errors"`

	/* meta

	Required: true
	*/
	Meta *MetaResponse `json:"meta"`
}

// Validate validates this parking signup422 response
func (m *ParkingSignup422Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateErrors(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMeta(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ParkingSignup422Response) validateData(formats strfmt.Registry) error {

	return nil
}

func (m *ParkingSignup422Response) validateErrors(formats strfmt.Registry) error {

	if err := validate.Required("errors", "body", m.Errors); err != nil {
		return err
	}

	if m.Errors != nil {

		if err := m.Errors.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *ParkingSignup422Response) validateMeta(formats strfmt.Registry) error {

	if err := validate.Required("meta", "body", m.Meta); err != nil {
		return err
	}

	if m.Meta != nil {

		if err := m.Meta.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
