package models_secure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*UnauthorizedError unauthorized error

swagger:model unauthorizedError
*/
type UnauthorizedError struct {

	/* access token
	 */
	AccessToken ValidationError `json:"access_token,omitempty"`

	/* A list of errors that do not apply to an attribute.
	 */
	Base ValidationError `json:"base,omitempty"`

	/* credentials
	 */
	Credentials ValidationError `json:"credentials,omitempty"`
}

// Validate validates this unauthorized error
func (m *UnauthorizedError) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
