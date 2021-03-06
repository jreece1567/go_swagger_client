package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*Occasion422Error occasion422 error

swagger:model occasion422Error
*/
type Occasion422Error struct {

	/* A list of errors that do not apply to an attribute.
	 */
	Base ValidationError `json:"base,omitempty"`

	/* centre id
	 */
	CentreID ValidationError `json:"centre_id,omitempty"`

	/* ends at
	 */
	EndsAt ValidationError `json:"ends_at,omitempty"`

	/* name
	 */
	Name ValidationError `json:"name,omitempty"`

	/* starts at
	 */
	StartsAt ValidationError `json:"starts_at,omitempty"`
}

// Validate validates this occasion422 error
func (m *Occasion422Error) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
