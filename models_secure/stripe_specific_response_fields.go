package models_secure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*StripeSpecificResponseFields stripe specific response fields

swagger:model stripeSpecificResponseFields
*/
type StripeSpecificResponseFields struct {

	/* A response from Stripe may include a card_funding response field that indicates what type of funding source was used to complete the transaction, such as credit or debit.
	 */
	CardFunding string `json:"card_funding,omitempty"`
}

// Validate validates this stripe specific response fields
func (m *StripeSpecificResponseFields) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
