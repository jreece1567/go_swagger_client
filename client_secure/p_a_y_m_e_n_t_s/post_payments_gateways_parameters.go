package p_a_y_m_e_n_t_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"restclient/models_secure"
)

// NewPostPaymentsGatewaysParams creates a new PostPaymentsGatewaysParams object
// with the default values initialized.
func NewPostPaymentsGatewaysParams() *PostPaymentsGatewaysParams {
	var ()
	return &PostPaymentsGatewaysParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPaymentsGatewaysParamsWithTimeout creates a new PostPaymentsGatewaysParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPaymentsGatewaysParamsWithTimeout(timeout time.Duration) *PostPaymentsGatewaysParams {
	var ()
	return &PostPaymentsGatewaysParams{

		timeout: timeout,
	}
}

/*PostPaymentsGatewaysParams contains all the parameters to send to the API endpoint
for the post payments gateways operation typically these are written to a http.Request
*/
type PostPaymentsGatewaysParams struct {

	/*Authorization
	  Access token

	*/
	Authorization string
	/*GatewayData
	  Data to create a gateway.

	*/
	GatewayData *models_secure.GatewayCreateBody

	timeout time.Duration
}

// WithAuthorization adds the authorization to the post payments gateways params
func (o *PostPaymentsGatewaysParams) WithAuthorization(Authorization string) *PostPaymentsGatewaysParams {
	o.Authorization = Authorization
	return o
}

// WithGatewayData adds the gatewayData to the post payments gateways params
func (o *PostPaymentsGatewaysParams) WithGatewayData(GatewayData *models_secure.GatewayCreateBody) *PostPaymentsGatewaysParams {
	o.GatewayData = GatewayData
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostPaymentsGatewaysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.GatewayData == nil {
		o.GatewayData = new(models_secure.GatewayCreateBody)
	}

	if err := r.SetBodyParam(o.GatewayData); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
