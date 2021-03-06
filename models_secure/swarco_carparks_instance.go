package models_secure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*SwarcoCarparksInstance Data to update a Swarco car park.

swagger:model swarcoCarparksInstance
*/
type SwarcoCarparksInstance struct {

	/* l o c a t i o n

	Required: true
	*/
	LOCATION *SwarcoLocationInstance `json:"LOCATION"`
}

// Validate validates this swarco carparks instance
func (m *SwarcoCarparksInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLOCATION(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SwarcoCarparksInstance) validateLOCATION(formats strfmt.Registry) error {

	if err := validate.Required("LOCATION", "body", m.LOCATION); err != nil {
		return err
	}

	if m.LOCATION != nil {

		if err := m.LOCATION.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
