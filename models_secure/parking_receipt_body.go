package models_secure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*ParkingReceiptBody parking receipt body

swagger:model parkingReceiptBody
*/
type ParkingReceiptBody struct {

	/* Parking session identifier. Used to retrieve the data for the receipt. Session must belong to the person.

	Required: true
	*/
	ParkingSessionID *string `json:"parking_session_id"`
}

// Validate validates this parking receipt body
func (m *ParkingReceiptBody) Validate(formats strfmt.Registry) error {
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

func (m *ParkingReceiptBody) validateParkingSessionID(formats strfmt.Registry) error {

	if err := validate.Required("parking_session_id", "body", m.ParkingSessionID); err != nil {
		return err
	}

	return nil
}
