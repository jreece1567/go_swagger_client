package models_secure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

/*GatewaySpecificFieldsForAuthorization Specific fields passed to gateway when processing transactions for this merchant.

swagger:model gatewaySpecificFieldsForAuthorization
*/
type GatewaySpecificFieldsForAuthorization struct {

	/* stripe
	 */
	Stripe *StripeAuthorizationSpecificFields `json:"stripe,omitempty"`
}

// Validate validates this gateway specific fields for authorization
func (m *GatewaySpecificFieldsForAuthorization) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStripe(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GatewaySpecificFieldsForAuthorization) validateStripe(formats strfmt.Registry) error {

	if swag.IsZero(m.Stripe) { // not required
		return nil
	}

	if m.Stripe != nil {

		if err := m.Stripe.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
