package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*OccasionUpdateBody Data to update an occasion.

swagger:model occasionUpdateBody
*/
type OccasionUpdateBody struct {

	/* Occasion name.
	 */
	Name string `json:"name,omitempty"`
}

// Validate validates this occasion update body
func (m *OccasionUpdateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
