package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*CurationCreateResponse Information about the new curation

swagger:model curationCreateResponse
*/
type CurationCreateResponse struct {

	/* data

	Required: true
	*/
	Data *CurationInstance `json:"data"`

	/* The `errors` component of the response

	Required: true
	*/
	Errors EmptyObject `json:"errors"`

	/* The `meta` component of the response

	Required: true
	*/
	Meta *MetaResponse `json:"meta"`
}

// Validate validates this curation create response
func (m *CurationCreateResponse) Validate(formats strfmt.Registry) error {
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

func (m *CurationCreateResponse) validateData(formats strfmt.Registry) error {

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

func (m *CurationCreateResponse) validateErrors(formats strfmt.Registry) error {

	return nil
}

func (m *CurationCreateResponse) validateMeta(formats strfmt.Registry) error {

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