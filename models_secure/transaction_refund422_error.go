package models_secure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*TransactionRefund422Error transaction refund422 error

swagger:model transactionRefund422Error
*/
type TransactionRefund422Error struct {

	/* amount
	 */
	Amount ValidationError `json:"amount,omitempty"`

	/* A list of errors that do not apply to an attribute.
	 */
	Base ValidationError `json:"base,omitempty"`
}

// Validate validates this transaction refund422 error
func (m *TransactionRefund422Error) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}