package models_secure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*TransactionAuthorizeOrPurchaseBody Data to authorize a transaction.

swagger:model transactionAuthorizeOrPurchaseBody
*/
type TransactionAuthorizeOrPurchaseBody struct {

	/* Amount to be charged, in cents.

	Required: true
	*/
	Amount *int64 `json:"amount"`

	/* Currency of the amount being charged. ISO-4217 format.

	Required: true
	*/
	CurrencyCode *string `json:"currency_code"`

	/* Description of product or service sold.
	 */
	Description string `json:"description,omitempty"`

	/* Customer email address.
	 */
	Email string `json:"email,omitempty"`

	/* gateway specific fields
	 */
	GatewaySpecificFields *GatewaySpecificFieldsForAuthorization `json:"gateway_specific_fields,omitempty"`

	/* Customer IPv4 address.
	 */
	IP string `json:"ip,omitempty"`

	/* Code of the merchant that will receive the amount.

	Required: true
	*/
	MerchantCode *string `json:"merchant_code"`

	/* Custom location description that may end up on bank statement, depending on the gateway.
	 */
	MerchantLocationDescriptor string `json:"merchant_location_descriptor,omitempty"`

	/* Custom name description that may end up on bank statement, depending on the gateway.
	 */
	MerchantNameDescriptor string `json:"merchant_name_descriptor,omitempty"`

	/* A merchant specified tracking number.
	 */
	OrderID string `json:"order_id,omitempty"`

	/* Token of the payment method that should be charged.

	Required: true
	*/
	PaymentMethodToken *string `json:"payment_method_token"`
}

// Validate validates this transaction authorize or purchase body
func (m *TransactionAuthorizeOrPurchaseBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCurrencyCode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGatewaySpecificFields(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMerchantCode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePaymentMethodToken(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransactionAuthorizeOrPurchaseBody) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *TransactionAuthorizeOrPurchaseBody) validateCurrencyCode(formats strfmt.Registry) error {

	if err := validate.Required("currency_code", "body", m.CurrencyCode); err != nil {
		return err
	}

	return nil
}

func (m *TransactionAuthorizeOrPurchaseBody) validateGatewaySpecificFields(formats strfmt.Registry) error {

	if swag.IsZero(m.GatewaySpecificFields) { // not required
		return nil
	}

	if m.GatewaySpecificFields != nil {

		if err := m.GatewaySpecificFields.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *TransactionAuthorizeOrPurchaseBody) validateMerchantCode(formats strfmt.Registry) error {

	if err := validate.Required("merchant_code", "body", m.MerchantCode); err != nil {
		return err
	}

	return nil
}

func (m *TransactionAuthorizeOrPurchaseBody) validatePaymentMethodToken(formats strfmt.Registry) error {

	if err := validate.Required("payment_method_token", "body", m.PaymentMethodToken); err != nil {
		return err
	}

	return nil
}
