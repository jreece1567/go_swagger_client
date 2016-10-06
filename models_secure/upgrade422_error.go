package models_secure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*Upgrade422Error upgrade422 error

swagger:model upgrade422Error
*/
type Upgrade422Error struct {

	/* Password and password confirmation do not match, or account is not a partial account
	 */
	Base ValidationError `json:"base,omitempty"`
}

// Validate validates this upgrade422 error
func (m *Upgrade422Error) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}