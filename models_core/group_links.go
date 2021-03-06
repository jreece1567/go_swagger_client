package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*GroupLinks Links relating to the group.

swagger:model groupLinks
*/
type GroupLinks struct {

	/* Canonical link to this group.

	Required: true
	*/
	Self *GenericLink `json:"self"`
}

// Validate validates this group links
func (m *GroupLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupLinks) validateSelf(formats strfmt.Registry) error {

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
