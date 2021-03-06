package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*MovieClassification Movie classification information

swagger:model movieClassification
*/
type MovieClassification struct {

	/* Organisation who rated the movie

	Required: true
	*/
	Body *string `json:"body"`

	/* Information as to why the movie has the rating it does

	Required: true
	*/
	Detail *string `json:"detail"`

	/* Movie rating code

	Required: true
	*/
	Rating *string `json:"rating"`
}

// Validate validates this movie classification
func (m *MovieClassification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBody(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDetail(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRating(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MovieClassification) validateBody(formats strfmt.Registry) error {

	if err := validate.Required("body", "body", m.Body); err != nil {
		return err
	}

	return nil
}

func (m *MovieClassification) validateDetail(formats strfmt.Registry) error {

	if err := validate.Required("detail", "body", m.Detail); err != nil {
		return err
	}

	return nil
}

func (m *MovieClassification) validateRating(formats strfmt.Registry) error {

	if err := validate.Required("rating", "body", m.Rating); err != nil {
		return err
	}

	return nil
}
