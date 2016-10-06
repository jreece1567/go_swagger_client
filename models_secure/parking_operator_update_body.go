package models_secure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*ParkingOperatorUpdateBody Data to update a Car Parking Business Operator.

swagger:model parkingOperatorUpdateBody
*/
type ParkingOperatorUpdateBody struct {

	/* Full address for the Car Parking Business Operator.
	 */
	Address string `json:"address,omitempty"`

	/* Code used to process payments for this operator.
	 */
	MerchantCode string `json:"merchant_code,omitempty"`

	/* Name of the Car Parking Business Operator
	 */
	Name string `json:"name,omitempty"`

	/* Bonus time in seconds given to users to avoid complaints about busy service.
	 */
	ParkingTimeBonus int64 `json:"parking_time_bonus,omitempty"`

	/* Phone number of the Car Parking Business Operator
	 */
	PhoneNumber string `json:"phone_number,omitempty"`

	/* Tax identifier of the Car Parking Business Operator
	 */
	TaxID string `json:"tax_id,omitempty"`
}

// Validate validates this parking operator update body
func (m *ParkingOperatorUpdateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
