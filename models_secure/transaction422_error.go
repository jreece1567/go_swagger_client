package models_secure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*Transaction422Error transaction422 error

swagger:model transaction422Error
*/
type Transaction422Error struct {

	/* amount
	 */
	Amount ValidationError `json:"amount,omitempty"`

	/* A list of errors that do not apply to an attribute.
	 */
	Base ValidationError `json:"base,omitempty"`

	/* merchant code
	 */
	MerchantCode ValidationError `json:"merchant_code,omitempty"`

	/* payment method token
	 */
	PaymentMethodToken ValidationError `json:"payment_method_token,omitempty"`
}

// Validate validates this transaction422 error
func (m *Transaction422Error) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
