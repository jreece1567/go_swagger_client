package s_t_a_f_f_a_u_t_h

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetStaffTokenInfoParams creates a new GetStaffTokenInfoParams object
// with the default values initialized.
func NewGetStaffTokenInfoParams() *GetStaffTokenInfoParams {
	var ()
	return &GetStaffTokenInfoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStaffTokenInfoParamsWithTimeout creates a new GetStaffTokenInfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStaffTokenInfoParamsWithTimeout(timeout time.Duration) *GetStaffTokenInfoParams {
	var ()
	return &GetStaffTokenInfoParams{

		timeout: timeout,
	}
}

/*GetStaffTokenInfoParams contains all the parameters to send to the API endpoint
for the get staff token info operation typically these are written to a http.Request
*/
type GetStaffTokenInfoParams struct {

	/*Token*/
	Token string

	timeout time.Duration
}

// WithToken adds the token to the get staff token info params
func (o *GetStaffTokenInfoParams) WithToken(Token string) *GetStaffTokenInfoParams {
	o.Token = Token
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetStaffTokenInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// query param token
	qrToken := o.Token
	qToken := qrToken
	if qToken != "" {
		if err := r.SetQueryParam("token", qToken); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
