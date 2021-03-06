package p_a_y_m_e_n_t_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPaymentsGatewaysGatewayIDParams creates a new GetPaymentsGatewaysGatewayIDParams object
// with the default values initialized.
func NewGetPaymentsGatewaysGatewayIDParams() *GetPaymentsGatewaysGatewayIDParams {
	var ()
	return &GetPaymentsGatewaysGatewayIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPaymentsGatewaysGatewayIDParamsWithTimeout creates a new GetPaymentsGatewaysGatewayIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPaymentsGatewaysGatewayIDParamsWithTimeout(timeout time.Duration) *GetPaymentsGatewaysGatewayIDParams {
	var ()
	return &GetPaymentsGatewaysGatewayIDParams{

		timeout: timeout,
	}
}

/*GetPaymentsGatewaysGatewayIDParams contains all the parameters to send to the API endpoint
for the get payments gateways gateway ID operation typically these are written to a http.Request
*/
type GetPaymentsGatewaysGatewayIDParams struct {

	/*Authorization
	  Access token

	*/
	Authorization string
	/*Fields
	  Fields in response. Array that lists the fields requested. If empty or not supplied, all attributes will be returned.

	*/
	Fields []string
	/*GatewayID
	  Gateway identifier.

	*/
	GatewayID int64

	timeout time.Duration
}

// WithAuthorization adds the authorization to the get payments gateways gateway ID params
func (o *GetPaymentsGatewaysGatewayIDParams) WithAuthorization(Authorization string) *GetPaymentsGatewaysGatewayIDParams {
	o.Authorization = Authorization
	return o
}

// WithFields adds the fields to the get payments gateways gateway ID params
func (o *GetPaymentsGatewaysGatewayIDParams) WithFields(Fields []string) *GetPaymentsGatewaysGatewayIDParams {
	o.Fields = Fields
	return o
}

// WithGatewayID adds the gatewayId to the get payments gateways gateway ID params
func (o *GetPaymentsGatewaysGatewayIDParams) WithGatewayID(GatewayID int64) *GetPaymentsGatewaysGatewayIDParams {
	o.GatewayID = GatewayID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetPaymentsGatewaysGatewayIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	valuesFields := o.Fields

	joinedFields := swag.JoinByFormat(valuesFields, "csv")
	// query array param fields
	if err := r.SetQueryParam("fields", joinedFields...); err != nil {
		return err
	}

	// path param gateway_id
	if err := r.SetPathParam("gateway_id", swag.FormatInt64(o.GatewayID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
