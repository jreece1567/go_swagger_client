package models_secure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*ParkingSessionRefundBody Data to process a parking session refund.

swagger:model parkingSessionRefundBody
*/
type ParkingSessionRefundBody struct {

	/* Amount to be refunded
	 */
	Amount int64 `json:"amount,omitempty"`

	/* Reason for the session refund
	 */
	Reason string `json:"reason,omitempty"`

	/* Identifier of the staff user that requested the refund
	 */
	RequesterID string `json:"requester_id,omitempty"`
}

// Validate validates this parking session refund body
func (m *ParkingSessionRefundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}