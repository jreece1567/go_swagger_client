package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*Deal422Error deal422 error

swagger:model deal422Error
*/
type Deal422Error struct {

	/* A list of errors that do not apply to an attribute.
	 */
	Base ValidationError `json:"base,omitempty"`

	/* ends at
	 */
	EndsAt ValidationError `json:"ends_at,omitempty"`

	/* published at
	 */
	PublishedAt ValidationError `json:"published_at,omitempty"`

	/* starts at
	 */
	StartsAt ValidationError `json:"starts_at,omitempty"`

	/* stores
	 */
	Stores ValidationError `json:"stores,omitempty"`

	/* terms and conditions
	 */
	TermsAndConditions ValidationError `json:"terms_and_conditions,omitempty"`

	/* title
	 */
	Title ValidationError `json:"title,omitempty"`
}

// Validate validates this deal422 error
func (m *Deal422Error) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
