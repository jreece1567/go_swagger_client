package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*Schedule422Error schedule422 error

swagger:model schedule422Error
*/
type Schedule422Error struct {

	/* A list of errors that do not apply to an attribute.
	 */
	Base ValidationError `json:"base,omitempty"`

	/* ends at
	 */
	EndsAt ValidationError `json:"ends_at,omitempty"`

	/* events
	 */
	Events ValidationError `json:"events,omitempty"`

	/* name
	 */
	Name ValidationError `json:"name,omitempty"`

	/* participant
	 */
	Participant ValidationError `json:"participant,omitempty"`

	/* starts at
	 */
	StartsAt ValidationError `json:"starts_at,omitempty"`

	/* time zone
	 */
	TimeZone ValidationError `json:"time_zone,omitempty"`

	/* type
	 */
	Type ValidationError `json:"type,omitempty"`
}

// Validate validates this schedule422 error
func (m *Schedule422Error) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
