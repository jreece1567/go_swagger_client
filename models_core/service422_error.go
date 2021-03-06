package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*Service422Error service422 error

swagger:model service422Error
*/
type Service422Error struct {

	/* active
	 */
	Active ValidationError `json:"active,omitempty"`

	/* A list of errors that do not apply to an attribute.
	 */
	Base ValidationError `json:"base,omitempty"`

	/* centre id
	 */
	CentreID ValidationError `json:"centre_id,omitempty"`

	/* description
	 */
	Description ValidationError `json:"description,omitempty"`

	/* email
	 */
	Email ValidationError `json:"email,omitempty"`

	/* image ref
	 */
	ImageRef ValidationError `json:"image_ref,omitempty"`

	/* long title
	 */
	LongTitle ValidationError `json:"long_title,omitempty"`

	/* phone number
	 */
	PhoneNumber ValidationError `json:"phone_number,omitempty"`

	/* rates
	 */
	Rates ValidationError `json:"rates,omitempty"`

	/* service type
	 */
	ServiceType ValidationError `json:"service_type,omitempty"`

	/* short title
	 */
	ShortTitle ValidationError `json:"short_title,omitempty"`

	/* sort order
	 */
	SortOrder ValidationError `json:"sort_order,omitempty"`

	/* url
	 */
	URL ValidationError `json:"url,omitempty"`
}

// Validate validates this service422 error
func (m *Service422Error) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
