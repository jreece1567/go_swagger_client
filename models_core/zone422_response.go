package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*Zone422Response zone422 response

swagger:model zone422Response
*/
type Zone422Response struct {

	/* data

	Required: true
	*/
	Data EmptyObject `json:"data"`

	/* errors

	Required: true
	*/
	Errors *Zone422Error `json:"errors"`

	/* meta

	Required: true
	*/
	Meta *MetaResponse `json:"meta"`
}

// Validate validates this zone422 response
func (m *Zone422Response) Validate(formats strfmt.Registry) error {
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

func (m *Zone422Response) validateData(formats strfmt.Registry) error {

	return nil
}

func (m *Zone422Response) validateErrors(formats strfmt.Registry) error {

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

func (m *Zone422Response) validateMeta(formats strfmt.Registry) error {

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