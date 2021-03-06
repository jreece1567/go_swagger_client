package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*Occurrence Range of times when the event occurs

swagger:model occurrence
*/
type Occurrence struct {

	/* Date and time the occurrence of the event finishes

	Required: true
	*/
	FinishesAt *strfmt.DateTime `json:"finishes_at"`

	/* Date and time the occurrence of the event starts

	Required: true
	*/
	StartsAt *strfmt.DateTime `json:"starts_at"`
}

// Validate validates this occurrence
func (m *Occurrence) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFinishesAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStartsAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Occurrence) validateFinishesAt(formats strfmt.Registry) error {

	if err := validate.Required("finishes_at", "body", m.FinishesAt); err != nil {
		return err
	}

	return nil
}

func (m *Occurrence) validateStartsAt(formats strfmt.Registry) error {

	if err := validate.Required("starts_at", "body", m.StartsAt); err != nil {
		return err
	}

	return nil
}
