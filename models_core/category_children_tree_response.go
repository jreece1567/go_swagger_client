package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*CategoryChildrenTreeResponse List category tree response.

swagger:model categoryChildrenTreeResponse
*/
type CategoryChildrenTreeResponse struct {

	/* data

	Required: true
	*/
	Data *CategoryChildrenInstance `json:"data"`

	/* errors

	Required: true
	*/
	Errors EmptyObject `json:"errors"`

	/* meta

	Required: true
	*/
	Meta *MetaResponse `json:"meta"`
}

// Validate validates this category children tree response
func (m *CategoryChildrenTreeResponse) Validate(formats strfmt.Registry) error {
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

func (m *CategoryChildrenTreeResponse) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	if m.Data != nil {

		if err := m.Data.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *CategoryChildrenTreeResponse) validateErrors(formats strfmt.Registry) error {

	return nil
}

func (m *CategoryChildrenTreeResponse) validateMeta(formats strfmt.Registry) error {

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
