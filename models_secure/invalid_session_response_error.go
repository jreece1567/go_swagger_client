package models_secure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*InvalidSessionResponseError Invalid session identifier for person

swagger:model invalidSessionResponseError
*/
type InvalidSessionResponseError struct {

	/* parking session id

	Required: true
	*/
	ParkingSessionID []string `json:"parking_session_id"`
}

// Validate validates this invalid session response error
func (m *InvalidSessionResponseError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateParkingSessionID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InvalidSessionResponseError) validateParkingSessionID(formats strfmt.Registry) error {

	if err := validate.Required("parking_session_id", "body", m.ParkingSessionID); err != nil {
		return err
	}

	return nil
}
