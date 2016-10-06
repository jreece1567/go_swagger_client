package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*Locale422Error locale422 error

swagger:model locale422Error
*/
type Locale422Error struct {

	/* A list of errors that do not apply to an attribute.
	 */
	Base ValidationError `json:"base,omitempty"`

	/* description
	 */
	Description ValidationError `json:"description,omitempty"`

	/* name
	 */
	Name ValidationError `json:"name,omitempty"`

	/* path
	 */
	Path ValidationError `json:"path,omitempty"`

	/* slug
	 */
	Slug ValidationError `json:"slug,omitempty"`

	/* westfield locale
	 */
	WestfieldLocale ValidationError `json:"westfield_locale,omitempty"`
}

// Validate validates this locale422 error
func (m *Locale422Error) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
