package models_secure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*TokenBody Token parameter

swagger:model tokenBody
*/
type TokenBody struct {

	/* Access token

	Required: true
	*/
	Token *string `json:"token"`
}

// Validate validates this token body
func (m *TokenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateToken(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TokenBody) validateToken(formats strfmt.Registry) error {

	if err := validate.Required("token", "body", m.Token); err != nil {
		return err
	}

	return nil
}
