package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*CurationLinks curation links

swagger:model curationLinks
*/
type CurationLinks struct {

	/* Link to the image associated with this curation

	Required: true
	*/
	Image *GenericLink `json:"image"`

	/* Canonical URL for this curation

	Required: true
	*/
	Self *GenericLink `json:"self"`
}

// Validate validates this curation links
func (m *CurationLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CurationLinks) validateImage(formats strfmt.Registry) error {

	if err := validate.Required("image", "body", m.Image); err != nil {
		return err
	}

	if m.Image != nil {

		if err := m.Image.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *CurationLinks) validateSelf(formats strfmt.Registry) error {

	if err := validate.Required("self", "body", m.Self); err != nil {
		return err
	}

	if m.Self != nil {

		if err := m.Self.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
