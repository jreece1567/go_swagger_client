package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*Directory422Error directory422 error

swagger:model directory422Error
*/
type Directory422Error struct {

	/* name
	 */
	Name ValidationError `json:"name,omitempty"`
}

// Validate validates this directory422 error
func (m *Directory422Error) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
