package models_secure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

/*Staff401Error Unauthorized.

swagger:model staff401Error
*/
type Staff401Error struct {

	/* List of errors with this attribute.
	 */
	AccessToken []string `json:"access_token,omitempty"`

	/* List of errors with this attribute.
	 */
	XAPISecret []string `json:"x_api_secret,omitempty"`
}

// Validate validates this staff401 error
func (m *Staff401Error) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessToken(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateXAPISecret(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Staff401Error) validateAccessToken(formats strfmt.Registry) error {

	if swag.IsZero(m.AccessToken) { // not required
		return nil
	}

	return nil
}

func (m *Staff401Error) validateXAPISecret(formats strfmt.Registry) error {

	if swag.IsZero(m.XAPISecret) { // not required
		return nil
	}

	return nil
}
