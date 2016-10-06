package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*EmbeddedCurations Embedded curation

swagger:model embeddedCurations
*/
type EmbeddedCurations struct {

	/* curations

	Required: true
	*/
	Curations []*EmbeddedCuration `json:"curations"`
}

// Validate validates this embedded curations
func (m *EmbeddedCurations) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurations(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmbeddedCurations) validateCurations(formats strfmt.Registry) error {

	if err := validate.Required("curations", "body", m.Curations); err != nil {
		return err
	}

	for i := 0; i < len(m.Curations); i++ {

		if swag.IsZero(m.Curations[i]) { // not required
			continue
		}

		if m.Curations[i] != nil {

			if err := m.Curations[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}
