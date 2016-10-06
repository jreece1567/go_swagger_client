package models_secure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*Gateway422Error gateway422 error

swagger:model gateway422Error
*/
type Gateway422Error struct {

	/* A list of errors that do not apply to an attribute.
	 */
	Base ValidationError `json:"base,omitempty"`

	/* name
	 */
	Name ValidationError `json:"name,omitempty"`

	/* token
	 */
	Token ValidationError `json:"token,omitempty"`
}

// Validate validates this gateway422 error
func (m *Gateway422Error) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}