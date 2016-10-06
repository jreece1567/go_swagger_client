package p_a_y_m_e_n_t_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPaymentsPaymentMethodsPaymentMethodTokenParams creates a new GetPaymentsPaymentMethodsPaymentMethodTokenParams object
// with the default values initialized.
func NewGetPaymentsPaymentMethodsPaymentMethodTokenParams() *GetPaymentsPaymentMethodsPaymentMethodTokenParams {
	var ()
	return &GetPaymentsPaymentMethodsPaymentMethodTokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPaymentsPaymentMethodsPaymentMethodTokenParamsWithTimeout creates a new GetPaymentsPaymentMethodsPaymentMethodTokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPaymentsPaymentMethodsPaymentMethodTokenParamsWithTimeout(timeout time.Duration) *GetPaymentsPaymentMethodsPaymentMethodTokenParams {
	var ()
	return &GetPaymentsPaymentMethodsPaymentMethodTokenParams{

		timeout: timeout,
	}
}

/*GetPaymentsPaymentMethodsPaymentMethodTokenParams contains all the parameters to send to the API endpoint
for the get payments payment methods payment method token operation typically these are written to a http.Request
*/
type GetPaymentsPaymentMethodsPaymentMethodTokenParams struct {

	/*PaymentMethodToken
	  Payment method token.

	*/
	PaymentMethodToken string

	timeout time.Duration
}

// WithPaymentMethodToken adds the paymentMethodToken to the get payments payment methods payment method token params
func (o *GetPaymentsPaymentMethodsPaymentMethodTokenParams) WithPaymentMethodToken(PaymentMethodToken string) *GetPaymentsPaymentMethodsPaymentMethodTokenParams {
	o.PaymentMethodToken = PaymentMethodToken
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetPaymentsPaymentMethodsPaymentMethodTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param payment_method_token
	if err := r.SetPathParam("payment_method_token", o.PaymentMethodToken); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
