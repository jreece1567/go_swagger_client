package models_secure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*ParkingRefundBody parking refund body

swagger:model parkingRefundBody
*/
type ParkingRefundBody struct {

	/* Amount to be refunded.
	 */
	Amount int64 `json:"amount,omitempty"`

	/* Parking session identifier of the session to be refunded. Session must belong to the person.

	Required: true
	*/
	ParkingSessionID *string `json:"parking_session_id"`

	/* Reason of the refund.
	 */
	Reason string `json:"reason,omitempty"`
}

// Validate validates this parking refund body
func (m *ParkingRefundBody) Validate(formats strfmt.Registry) error {
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

func (m *ParkingRefundBody) validateParkingSessionID(formats strfmt.Registry) error {

	if err := validate.Required("parking_session_id", "body", m.ParkingSessionID); err != nil {
		return err
	}

	return nil
}
