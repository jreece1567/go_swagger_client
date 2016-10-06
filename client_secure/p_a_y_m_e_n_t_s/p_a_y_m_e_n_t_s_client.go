package p_a_y_m_e_n_t_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new p a y m e n t s API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for p a y m e n t s API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeletePaymentsGatewaysGatewayID deletes a gateway

Request deletion of a gateway.
*/
func (a *Client) DeletePaymentsGatewaysGatewayID(params *DeletePaymentsGatewaysGatewayIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeletePaymentsGatewaysGatewayIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePaymentsGatewaysGatewayIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeletePaymentsGatewaysGatewayID",
		Method:             "DELETE",
		PathPattern:        "/payments/gateways/{gateway_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeletePaymentsGatewaysGatewayIDReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeletePaymentsGatewaysGatewayIDNoContent), nil
}

/*
DeletePaymentsMerchantsMerchantID deletes a merchant

Request deletion of a merchant.
*/
func (a *Client) DeletePaymentsMerchantsMerchantID(params *DeletePaymentsMerchantsMerchantIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeletePaymentsMerchantsMerchantIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePaymentsMerchantsMerchantIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeletePaymentsMerchantsMerchantID",
		Method:             "DELETE",
		PathPattern:        "/payments/merchants/{merchant_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeletePaymentsMerchantsMerchantIDReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeletePaymentsMerchantsMerchantIDNoContent), nil
}

/*
DeletePaymentsPaymentMethodsPaymentMethodToken redacts a payment method

Request a redaction of a payment method. Redacting a payment method does not delete it. Its transaction history is maintained as your permanent record . However, all sensitive information (e.g. PAN, CVV) is removed immediately and permanently. A redacted payment method can no longer be used for transactions since the account information will have been disposed of.
*/
func (a *Client) DeletePaymentsPaymentMethodsPaymentMethodToken(params *DeletePaymentsPaymentMethodsPaymentMethodTokenParams, authInfo runtime.ClientAuthInfoWriter) (*DeletePaymentsPaymentMethodsPaymentMethodTokenNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePaymentsPaymentMethodsPaymentMethodTokenParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeletePaymentsPaymentMethodsPaymentMethodToken",
		Method:             "DELETE",
		PathPattern:        "/payments/payment_methods/{payment_method_token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeletePaymentsPaymentMethodsPaymentMethodTokenReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeletePaymentsPaymentMethodsPaymentMethodTokenNoContent), nil
}

/*
GetPaymentsGateways lists gateways

Request an array of gateways.
*/
func (a *Client) GetPaymentsGateways(params *GetPaymentsGatewaysParams, authInfo runtime.ClientAuthInfoWriter) (*GetPaymentsGatewaysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPaymentsGatewaysParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentsGateways",
		Method:             "GET",
		PathPattern:        "/payments/gateways",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPaymentsGatewaysReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentsGatewaysOK), nil
}

/*
GetPaymentsGatewaysGatewayID retrieves a gateway

Request the details of a gateway.
*/
func (a *Client) GetPaymentsGatewaysGatewayID(params *GetPaymentsGatewaysGatewayIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetPaymentsGatewaysGatewayIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPaymentsGatewaysGatewayIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentsGatewaysGatewayID",
		Method:             "GET",
		PathPattern:        "/payments/gateways/{gateway_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPaymentsGatewaysGatewayIDReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentsGatewaysGatewayIDOK), nil
}

/*
GetPaymentsMerchants lists merchants

Request an array of merchants.
*/
func (a *Client) GetPaymentsMerchants(params *GetPaymentsMerchantsParams, authInfo runtime.ClientAuthInfoWriter) (*GetPaymentsMerchantsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPaymentsMerchantsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentsMerchants",
		Method:             "GET",
		PathPattern:        "/payments/merchants",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPaymentsMerchantsReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentsMerchantsOK), nil
}

/*
GetPaymentsMerchantsMerchantID retrieves a merchant

Request the details of a merchant.
*/
func (a *Client) GetPaymentsMerchantsMerchantID(params *GetPaymentsMerchantsMerchantIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetPaymentsMerchantsMerchantIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPaymentsMerchantsMerchantIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentsMerchantsMerchantID",
		Method:             "GET",
		PathPattern:        "/payments/merchants/{merchant_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPaymentsMerchantsMerchantIDReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentsMerchantsMerchantIDOK), nil
}

/*
GetPaymentsPaymentMethodsPaymentMethodToken retrieves a payment method

Request the details of a payment method.
*/
func (a *Client) GetPaymentsPaymentMethodsPaymentMethodToken(params *GetPaymentsPaymentMethodsPaymentMethodTokenParams, authInfo runtime.ClientAuthInfoWriter) (*GetPaymentsPaymentMethodsPaymentMethodTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPaymentsPaymentMethodsPaymentMethodTokenParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentsPaymentMethodsPaymentMethodToken",
		Method:             "GET",
		PathPattern:        "/payments/payment_methods/{payment_method_token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPaymentsPaymentMethodsPaymentMethodTokenReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentsPaymentMethodsPaymentMethodTokenOK), nil
}

/*
GetPaymentsStripeMerchantsStripeAccountID retrieves a stripe merchant

Request the details of a Stripe merchant.
*/
func (a *Client) GetPaymentsStripeMerchantsStripeAccountID(params *GetPaymentsStripeMerchantsStripeAccountIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetPaymentsStripeMerchantsStripeAccountIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPaymentsStripeMerchantsStripeAccountIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentsStripeMerchantsStripeAccountID",
		Method:             "GET",
		PathPattern:        "/payments/stripe_merchants/{stripe_account_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPaymentsStripeMerchantsStripeAccountIDReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentsStripeMerchantsStripeAccountIDOK), nil
}

/*
GetPaymentsTransactionsTransactionID retrieves a transaction

Request the details of a single transaction.
*/
func (a *Client) GetPaymentsTransactionsTransactionID(params *GetPaymentsTransactionsTransactionIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetPaymentsTransactionsTransactionIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPaymentsTransactionsTransactionIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentsTransactionsTransactionID",
		Method:             "GET",
		PathPattern:        "/payments/transactions/{transaction_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPaymentsTransactionsTransactionIDReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentsTransactionsTransactionIDOK), nil
}

/*
PatchPaymentsGatewaysGatewayID updates a gateway

Request update of a gateway from JSON data in the request body.
*/
func (a *Client) PatchPaymentsGatewaysGatewayID(params *PatchPaymentsGatewaysGatewayIDParams, authInfo runtime.ClientAuthInfoWriter) (*PatchPaymentsGatewaysGatewayIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchPaymentsGatewaysGatewayIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchPaymentsGatewaysGatewayID",
		Method:             "PATCH",
		PathPattern:        "/payments/gateways/{gateway_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchPaymentsGatewaysGatewayIDReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchPaymentsGatewaysGatewayIDNoContent), nil
}

/*
PatchPaymentsMerchantsMerchantID updates a merchant

Request update of a merchant from JSON data in the request body.
*/
func (a *Client) PatchPaymentsMerchantsMerchantID(params *PatchPaymentsMerchantsMerchantIDParams, authInfo runtime.ClientAuthInfoWriter) (*PatchPaymentsMerchantsMerchantIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchPaymentsMerchantsMerchantIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchPaymentsMerchantsMerchantID",
		Method:             "PATCH",
		PathPattern:        "/payments/merchants/{merchant_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchPaymentsMerchantsMerchantIDReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchPaymentsMerchantsMerchantIDNoContent), nil
}

/*
PatchPaymentsPaymentMethodsPaymentMethodToken updates payment method details

Request an update of the details associated with a payment method.
*/
func (a *Client) PatchPaymentsPaymentMethodsPaymentMethodToken(params *PatchPaymentsPaymentMethodsPaymentMethodTokenParams, authInfo runtime.ClientAuthInfoWriter) (*PatchPaymentsPaymentMethodsPaymentMethodTokenNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchPaymentsPaymentMethodsPaymentMethodTokenParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchPaymentsPaymentMethodsPaymentMethodToken",
		Method:             "PATCH",
		PathPattern:        "/payments/payment_methods/{payment_method_token}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchPaymentsPaymentMethodsPaymentMethodTokenReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchPaymentsPaymentMethodsPaymentMethodTokenNoContent), nil
}

/*
PatchPaymentsStripeMerchantsStripeAccountID updates a stripe merchant

Request update of a Stripe Merchant from JSON data in the request body.
*/
func (a *Client) PatchPaymentsStripeMerchantsStripeAccountID(params *PatchPaymentsStripeMerchantsStripeAccountIDParams, authInfo runtime.ClientAuthInfoWriter) (*PatchPaymentsStripeMerchantsStripeAccountIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchPaymentsStripeMerchantsStripeAccountIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchPaymentsStripeMerchantsStripeAccountID",
		Method:             "PATCH",
		PathPattern:        "/payments/stripe_merchants/{stripe_account_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchPaymentsStripeMerchantsStripeAccountIDReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchPaymentsStripeMerchantsStripeAccountIDNoContent), nil
}

/*
PostPaymentsGateways creates a gateway

Request creation of a gateway from JSON data in the request body.
*/
func (a *Client) PostPaymentsGateways(params *PostPaymentsGatewaysParams, authInfo runtime.ClientAuthInfoWriter) (*PostPaymentsGatewaysCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPaymentsGatewaysParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostPaymentsGateways",
		Method:             "POST",
		PathPattern:        "/payments/gateways",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostPaymentsGatewaysReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostPaymentsGatewaysCreated), nil
}

/*
PostPaymentsMerchants creates a merchant

Request creation of a merchant from JSON data in the request body.
*/
func (a *Client) PostPaymentsMerchants(params *PostPaymentsMerchantsParams, authInfo runtime.ClientAuthInfoWriter) (*PostPaymentsMerchantsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPaymentsMerchantsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostPaymentsMerchants",
		Method:             "POST",
		PathPattern:        "/payments/merchants",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostPaymentsMerchantsReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostPaymentsMerchantsCreated), nil
}

/*
PostPaymentsPaymentMethodsPaymentMethodTokenVerify verifies and retain a payment method

Request verification of a payment method checking if it is in good standing. This call will also retain the payment method
*/
func (a *Client) PostPaymentsPaymentMethodsPaymentMethodTokenVerify(params *PostPaymentsPaymentMethodsPaymentMethodTokenVerifyParams, authInfo runtime.ClientAuthInfoWriter) (*PostPaymentsPaymentMethodsPaymentMethodTokenVerifyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPaymentsPaymentMethodsPaymentMethodTokenVerifyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostPaymentsPaymentMethodsPaymentMethodTokenVerify",
		Method:             "POST",
		PathPattern:        "/payments/payment_methods/{payment_method_token}/verify",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostPaymentsPaymentMethodsPaymentMethodTokenVerifyReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostPaymentsPaymentMethodsPaymentMethodTokenVerifyOK), nil
}

/*
PostPaymentsStripeEvents creates a new stripe event in response to a webhook

Request the creation of a new stripe webhook event.
*/
func (a *Client) PostPaymentsStripeEvents(params *PostPaymentsStripeEventsParams, authInfo runtime.ClientAuthInfoWriter) (*PostPaymentsStripeEventsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPaymentsStripeEventsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostPaymentsStripeEvents",
		Method:             "POST",
		PathPattern:        "/payments/stripe_events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostPaymentsStripeEventsReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostPaymentsStripeEventsNoContent), nil
}

/*
PostPaymentsTransactionsAuthorize authorizes a payment method

Request the authorization of a payment method from JSON data in the request body. This transaction is used to authorize a payment amount, and such payment method can be subsequently captured or voided in order to complete the transaction.
*/
func (a *Client) PostPaymentsTransactionsAuthorize(params *PostPaymentsTransactionsAuthorizeParams, authInfo runtime.ClientAuthInfoWriter) (*PostPaymentsTransactionsAuthorizeCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPaymentsTransactionsAuthorizeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostPaymentsTransactionsAuthorize",
		Method:             "POST",
		PathPattern:        "/payments/transactions/authorize",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostPaymentsTransactionsAuthorizeReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostPaymentsTransactionsAuthorizeCreated), nil
}

/*
PostPaymentsTransactionsPurchase charges a payment method

Request the purchase of a payment method from JSON data in the request body. This transaction is used to authorize and capture a payment amount.
*/
func (a *Client) PostPaymentsTransactionsPurchase(params *PostPaymentsTransactionsPurchaseParams, authInfo runtime.ClientAuthInfoWriter) (*PostPaymentsTransactionsPurchaseCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPaymentsTransactionsPurchaseParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostPaymentsTransactionsPurchase",
		Method:             "POST",
		PathPattern:        "/payments/transactions/purchase",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostPaymentsTransactionsPurchaseReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostPaymentsTransactionsPurchaseCreated), nil
}

/*
PostPaymentsTransactionsTransactionIDCapture captures a transaction

Request the capture of funds previously reserved by an authorization.
*/
func (a *Client) PostPaymentsTransactionsTransactionIDCapture(params *PostPaymentsTransactionsTransactionIDCaptureParams, authInfo runtime.ClientAuthInfoWriter) (*PostPaymentsTransactionsTransactionIDCaptureOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPaymentsTransactionsTransactionIDCaptureParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostPaymentsTransactionsTransactionIDCapture",
		Method:             "POST",
		PathPattern:        "/payments/transactions/{transaction_id}/capture",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostPaymentsTransactionsTransactionIDCaptureReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostPaymentsTransactionsTransactionIDCaptureOK), nil
}

/*
PostPaymentsTransactionsTransactionIDRefund refunds a transaction

Request the refund of a transaction. A refund is like a void, except it actually reverses a charge instead of just canceling a charge that hasn't yet been made.
*/
func (a *Client) PostPaymentsTransactionsTransactionIDRefund(params *PostPaymentsTransactionsTransactionIDRefundParams, authInfo runtime.ClientAuthInfoWriter) (*PostPaymentsTransactionsTransactionIDRefundOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPaymentsTransactionsTransactionIDRefundParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostPaymentsTransactionsTransactionIDRefund",
		Method:             "POST",
		PathPattern:        "/payments/transactions/{transaction_id}/refund",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostPaymentsTransactionsTransactionIDRefundReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostPaymentsTransactionsTransactionIDRefundOK), nil
}

/*
PostPaymentsTransactionsTransactionIDVoid voids a transaction

Request voiding (canceling) a transaction previously authorized.
*/
func (a *Client) PostPaymentsTransactionsTransactionIDVoid(params *PostPaymentsTransactionsTransactionIDVoidParams, authInfo runtime.ClientAuthInfoWriter) (*PostPaymentsTransactionsTransactionIDVoidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPaymentsTransactionsTransactionIDVoidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostPaymentsTransactionsTransactionIDVoid",
		Method:             "POST",
		PathPattern:        "/payments/transactions/{transaction_id}/void",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostPaymentsTransactionsTransactionIDVoidReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostPaymentsTransactionsTransactionIDVoidOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}