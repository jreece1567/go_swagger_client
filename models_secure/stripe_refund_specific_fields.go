package models_secure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*StripeRefundSpecificFields stripe refund specific fields

swagger:model stripeRefundSpecificFields
*/
type StripeRefundSpecificFields struct {

	/* This is an optional parameter you can specify when you run a refund on Stripe. It indicates whether the Stripe Connect application_fee should be refunded when refunding the charge.
	 */
	RefundApplicationFee bool `json:"refund_application_fee,omitempty"`

	/* This is an optional parameter you can specify when you run a refund.
	 */
	ReverseTransfer bool `json:"reverse_transfer,omitempty"`
}

// Validate validates this stripe refund specific fields
func (m *StripeRefundSpecificFields) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}