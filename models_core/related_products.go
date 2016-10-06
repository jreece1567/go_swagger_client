package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*RelatedProducts Related products

swagger:model relatedProducts
*/
type RelatedProducts struct {

	/* List of related products

	Required: true
	*/
	Related []*Product `json:"related"`

	/* List of similar products

	Required: true
	*/
	Similar []*Product `json:"similar"`
}

// Validate validates this related products
func (m *RelatedProducts) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRelated(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSimilar(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RelatedProducts) validateRelated(formats strfmt.Registry) error {

	if err := validate.Required("related", "body", m.Related); err != nil {
		return err
	}

	for i := 0; i < len(m.Related); i++ {

		if swag.IsZero(m.Related[i]) { // not required
			continue
		}

		if m.Related[i] != nil {

			if err := m.Related[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *RelatedProducts) validateSimilar(formats strfmt.Registry) error {

	if err := validate.Required("similar", "body", m.Similar); err != nil {
		return err
	}

	for i := 0; i < len(m.Similar); i++ {

		if swag.IsZero(m.Similar[i]) { // not required
			continue
		}

		if m.Similar[i] != nil {

			if err := m.Similar[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}
