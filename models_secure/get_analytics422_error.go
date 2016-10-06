package models_secure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*GetAnalytics422Error get analytics422 error

swagger:model getAnalytics422Error
*/
type GetAnalytics422Error struct {

	/* Owner identifier is missing or invalid
	 */
	Base ValidationError `json:"base,omitempty"`
}

// Validate validates this get analytics422 error
func (m *GetAnalytics422Error) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
