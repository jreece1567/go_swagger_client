package models_secure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*Payment422Error List of resources not found.

swagger:model payment422Error
*/
type Payment422Error struct {

	/* A list of errors that do not apply to an attribute.
	 */
	Base ValidationError `json:"base,omitempty"`
}

// Validate validates this payment422 error
func (m *Payment422Error) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}