package models_secure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*TransactionInstance Transaction

swagger:model transactionInstance
*/
type TransactionInstance struct {

	/* Amount charged, in cents.

	Required: true
	*/
	Amount *int64 `json:"amount"`

	/* Date-time the transaction was created.

	Required: true
	*/
	CreatedAt *strfmt.DateTime `json:"created_at"`

	/* Currency of the amount being charged. ISO-4217 format.

	Required: true
	*/
	CurrencyCode *string `json:"currency_code"`

	/* Description of product or service sold.
	 */
	Description string `json:"description,omitempty"`

	/* Fields that a gateway defines for a specific purpose but is not implemented by all gateways

	Required: true
	*/
	GatewaySpecificFields *GatewaySpecificFieldsForAuthorization `json:"gateway_specific_fields"`

	/* Additional fields that a gateway may choose to provide.

	Required: true
	*/
	GatewaySpecificResponseFields *GatewaySpecificResponseFields `json:"gateway_specific_response_fields"`

	/* The gateway's corresponding transaction id.

	Required: true
	*/
	GatewayTransactionID *string `json:"gateway_transaction_id"`

	/* Customer IPv4 address.
	 */
	IP string `json:"ip,omitempty"`

	/* Location of merchant.
	 */
	MerchantLocationDescriptor string `json:"merchant_location_descriptor,omitempty"`

	/* Name of merchant.
	 */
	MerchantNameDescriptor string `json:"merchant_name_descriptor,omitempty"`

	/* A brief description of the results of the transaction.

	Required: true
	*/
	Message *string `json:"message"`

	/* Value is true if transaction was made on test gateway

	Required: true
	*/
	OnTestGateway *bool `json:"on_test_gateway"`

	/* A merchant specified tracking number.
	 */
	OrderID string `json:"order_id,omitempty"`

	/* payment method

	Required: true
	*/
	PaymentMethod *PaymentMethodInstance `json:"payment_method"`

	/* response

	Required: true
	*/
	Response *PaymentMethodVerificationResponseInstance `json:"response"`

	/* The status of the transaction.

	Required: true
	*/
	Status *string `json:"status"`

	/* Value is true if transaction was successful.

	Required: true
	*/
	Succeeded *bool `json:"succeeded"`

	/* Transaction identifier. Identifier to retrieve this transaction.

	Required: true
	*/
	TransactionID *string `json:"transaction_id"`

	/* The type of transaction, e.g., Authorization, Capture, Credit

	Required: true
	*/
	TransactionType *string `json:"transaction_type"`
}

// Validate validates this transaction instance
func (m *TransactionInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
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

	if err := m.validateGatewaySpecificResponseFields(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGatewayTransactionID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOnTestGateway(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePaymentMethod(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateResponse(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSucceeded(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTransactionID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTransactionType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransactionInstance) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *TransactionInstance) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *TransactionInstance) validateCurrencyCode(formats strfmt.Registry) error {

	if err := validate.Required("currency_code", "body", m.CurrencyCode); err != nil {
		return err
	}

	return nil
}

func (m *TransactionInstance) validateGatewaySpecificFields(formats strfmt.Registry) error {

	if err := validate.Required("gateway_specific_fields", "body", m.GatewaySpecificFields); err != nil {
		return err
	}

	if m.GatewaySpecificFields != nil {

		if err := m.GatewaySpecificFields.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *TransactionInstance) validateGatewaySpecificResponseFields(formats strfmt.Registry) error {

	if err := validate.Required("gateway_specific_response_fields", "body", m.GatewaySpecificResponseFields); err != nil {
		return err
	}

	if m.GatewaySpecificResponseFields != nil {

		if err := m.GatewaySpecificResponseFields.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *TransactionInstance) validateGatewayTransactionID(formats strfmt.Registry) error {

	if err := validate.Required("gateway_transaction_id", "body", m.GatewayTransactionID); err != nil {
		return err
	}

	return nil
}

func (m *TransactionInstance) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

func (m *TransactionInstance) validateOnTestGateway(formats strfmt.Registry) error {

	if err := validate.Required("on_test_gateway", "body", m.OnTestGateway); err != nil {
		return err
	}

	return nil
}

func (m *TransactionInstance) validatePaymentMethod(formats strfmt.Registry) error {

	if err := validate.Required("payment_method", "body", m.PaymentMethod); err != nil {
		return err
	}

	if m.PaymentMethod != nil {

		if err := m.PaymentMethod.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *TransactionInstance) validateResponse(formats strfmt.Registry) error {

	if err := validate.Required("response", "body", m.Response); err != nil {
		return err
	}

	if m.Response != nil {

		if err := m.Response.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

var transactionInstanceTypeStatusPropEnum []interface{}

// prop value enum
func (m *TransactionInstance) validateStatusEnum(path, location string, value string) error {
	if transactionInstanceTypeStatusPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["succeeded","failed","gateway_processing_failed","gateway_processing_result_unknown"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			transactionInstanceTypeStatusPropEnum = append(transactionInstanceTypeStatusPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, transactionInstanceTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TransactionInstance) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *TransactionInstance) validateSucceeded(formats strfmt.Registry) error {

	if err := validate.Required("succeeded", "body", m.Succeeded); err != nil {
		return err
	}

	return nil
}

func (m *TransactionInstance) validateTransactionID(formats strfmt.Registry) error {

	if err := validate.Required("transaction_id", "body", m.TransactionID); err != nil {
		return err
	}

	return nil
}

var transactionInstanceTypeTransactionTypePropEnum []interface{}

// prop value enum
func (m *TransactionInstance) validateTransactionTypeEnum(path, location string, value string) error {
	if transactionInstanceTypeTransactionTypePropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Authorization","Capture","Credit","Verification","Void","Purchase"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			transactionInstanceTypeTransactionTypePropEnum = append(transactionInstanceTypeTransactionTypePropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, transactionInstanceTypeTransactionTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TransactionInstance) validateTransactionType(formats strfmt.Registry) error {

	if err := validate.Required("transaction_type", "body", m.TransactionType); err != nil {
		return err
	}

	// value enum
	if err := m.validateTransactionTypeEnum("transaction_type", "body", *m.TransactionType); err != nil {
		return err
	}

	return nil
}
