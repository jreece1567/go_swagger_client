package models_secure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*UpdatePassword422Error update password422 error

swagger:model updatePassword422Error
*/
type UpdatePassword422Error struct {

	/* current password
	 */
	CurrentPassword ValidationError `json:"current_password,omitempty"`

	/* new password
	 */
	NewPassword ValidationError `json:"new_password,omitempty"`
}

// Validate validates this update password422 error
func (m *UpdatePassword422Error) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}