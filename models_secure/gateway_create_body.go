package models_secure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*GatewayCreateBody Data to create a gateway.

swagger:model gatewayCreateBody
*/
type GatewayCreateBody struct {

	/* Name of the gateway.

	Required: true
	*/
	Name *string `json:"name"`

	/* Token used to process transactions on this gateway.

	Required: true
	*/
	Token *string `json:"token"`
}

// Validate validates this gateway create body
func (m *GatewayCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateToken(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GatewayCreateBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *GatewayCreateBody) validateToken(formats strfmt.Registry) error {

	if err := validate.Required("token", "body", m.Token); err != nil {
		return err
	}

	return nil
}