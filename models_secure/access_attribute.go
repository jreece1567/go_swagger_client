package models_secure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

/*AccessAttribute Access Attributes

swagger:model accessAttribute
*/
type AccessAttribute struct {

	/* Centres array. (Only for centre role)
	 */
	Centres []string `json:"centres,omitempty"`

	/* Retailer identifier. (Only Retailer role)
	 */
	RetailerID int64 `json:"retailer_id,omitempty"`
}

// Validate validates this access attribute
func (m *AccessAttribute) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCentres(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccessAttribute) validateCentres(formats strfmt.Registry) error {

	if swag.IsZero(m.Centres) { // not required
		return nil
	}

	return nil
}
