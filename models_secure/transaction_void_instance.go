package models_secure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*TransactionVoidInstance Transaction void.

swagger:model transactionVoidInstance
*/
type TransactionVoidInstance struct {

	/* Date and time of origination.

	Required: true
	*/
	CreatedAt *strfmt.DateTime `json:"created_at"`

	/* Description of the product or service rendered.

	Required: true
	*/
	Description *string `json:"description"`

	/* Fields that a gateway defines for a specific purpose but is not implemented by all gateways

	Required: true
	*/
	GatewaySpecificFields interface{} `json:"gateway_specific_fields"`

	/* A unique string generated by Spreedly to identify a gateway.

	Required: true
	*/
	GatewayToken *string `json:"gateway_token"`

	/* The gateway’s corresponding transaction id.

	Required: true
	*/
	GatewayTransactionID *string `json:"gateway_transaction_id"`

	/* IP address of the customer making the purchase.

	Required: true
	*/
	IP *string `json:"ip"`

	/* Location of merchant.

	Required: true
	*/
	MerchantLocationDescriptor *string `json:"merchant_location_descriptor"`

	/* Name of merchant.

	Required: true
	*/
	MerchantNameDescriptor *string `json:"merchant_name_descriptor"`

	/* A brief description of the results of the transaction.

	Required: true
	*/
	Message *string `json:"message"`

	/* Value is true if transaction was made on test gateway

	Required: true
	*/
	OnTestGateway *bool `json:"on_test_gateway"`

	/* A tracking number that you can declare.

	Required: true
	*/
	OrderID *string `json:"order_id"`

	/* response

	Required: true
	*/
	Response *PaymentMethodVerificationResponseInstance `json:"response"`

	/* The status of the transaction

	Required: true
	*/
	Status *string `json:"status"`

	/* Value is true if transaction was successful.

	Required: true
	*/
	Succeeded *bool `json:"succeeded"`

	/* Transaction identifier. A unique string generated by Spreedly to identify a transaction.

	Required: true
	*/
	TransactionID *string `json:"transaction_id"`

	/* The type of transaction, e.g., Authorization, Capture, Credit

	Required: true
	*/
	TransactionType *string `json:"transaction_type"`
}

// Validate validates this transaction void instance
func (m *TransactionVoidInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGatewaySpecificFields(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGatewayToken(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGatewayTransactionID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIP(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMerchantLocationDescriptor(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMerchantNameDescriptor(formats); err != nil {
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

	if err := m.validateOrderID(formats); err != nil {
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

func (m *TransactionVoidInstance) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *TransactionVoidInstance) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *TransactionVoidInstance) validateGatewaySpecificFields(formats strfmt.Registry) error {

	return nil
}

func (m *TransactionVoidInstance) validateGatewayToken(formats strfmt.Registry) error {

	if err := validate.Required("gateway_token", "body", m.GatewayToken); err != nil {
		return err
	}

	return nil
}

func (m *TransactionVoidInstance) validateGatewayTransactionID(formats strfmt.Registry) error {

	if err := validate.Required("gateway_transaction_id", "body", m.GatewayTransactionID); err != nil {
		return err
	}

	return nil
}

func (m *TransactionVoidInstance) validateIP(formats strfmt.Registry) error {

	if err := validate.Required("ip", "body", m.IP); err != nil {
		return err
	}

	return nil
}

func (m *TransactionVoidInstance) validateMerchantLocationDescriptor(formats strfmt.Registry) error {

	if err := validate.Required("merchant_location_descriptor", "body", m.MerchantLocationDescriptor); err != nil {
		return err
	}

	return nil
}

func (m *TransactionVoidInstance) validateMerchantNameDescriptor(formats strfmt.Registry) error {

	if err := validate.Required("merchant_name_descriptor", "body", m.MerchantNameDescriptor); err != nil {
		return err
	}

	return nil
}

func (m *TransactionVoidInstance) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

func (m *TransactionVoidInstance) validateOnTestGateway(formats strfmt.Registry) error {

	if err := validate.Required("on_test_gateway", "body", m.OnTestGateway); err != nil {
		return err
	}

	return nil
}

func (m *TransactionVoidInstance) validateOrderID(formats strfmt.Registry) error {

	if err := validate.Required("order_id", "body", m.OrderID); err != nil {
		return err
	}

	return nil
}

func (m *TransactionVoidInstance) validateResponse(formats strfmt.Registry) error {

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

func (m *TransactionVoidInstance) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *TransactionVoidInstance) validateSucceeded(formats strfmt.Registry) error {

	if err := validate.Required("succeeded", "body", m.Succeeded); err != nil {
		return err
	}

	return nil
}

func (m *TransactionVoidInstance) validateTransactionID(formats strfmt.Registry) error {

	if err := validate.Required("transaction_id", "body", m.TransactionID); err != nil {
		return err
	}

	return nil
}

var transactionVoidInstanceTypeTransactionTypePropEnum []interface{}

// prop value enum
func (m *TransactionVoidInstance) validateTransactionTypeEnum(path, location string, value string) error {
	if transactionVoidInstanceTypeTransactionTypePropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Authorization","Capture","Credit","Verification","Void","Purchase"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			transactionVoidInstanceTypeTransactionTypePropEnum = append(transactionVoidInstanceTypeTransactionTypePropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, transactionVoidInstanceTypeTransactionTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TransactionVoidInstance) validateTransactionType(formats strfmt.Registry) error {

	if err := validate.Required("transaction_type", "body", m.TransactionType); err != nil {
		return err
	}

	// value enum
	if err := m.validateTransactionTypeEnum("transaction_type", "body", *m.TransactionType); err != nil {
		return err
	}

	return nil
}
