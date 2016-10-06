package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*ServiceHourBody Service opening hours

swagger:model serviceHourBody
*/
type ServiceHourBody struct {

	/* Name of this specific hour
	 */
	Label string `json:"label,omitempty"`

	/* Type of hour
	 */
	Name string `json:"name,omitempty"`

	/* Order in which this hour should be displayed
	 */
	SortOrder int64 `json:"sort_order,omitempty"`

	/* The value of the hour
	 */
	Value string `json:"value,omitempty"`
}

// Validate validates this service hour body
func (m *ServiceHourBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}