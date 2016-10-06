package k_i_o_s_k

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetKioskKioskCentresParams creates a new GetKioskKioskCentresParams object
// with the default values initialized.
func NewGetKioskKioskCentresParams() *GetKioskKioskCentresParams {

	return &GetKioskKioskCentresParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetKioskKioskCentresParamsWithTimeout creates a new GetKioskKioskCentresParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetKioskKioskCentresParamsWithTimeout(timeout time.Duration) *GetKioskKioskCentresParams {

	return &GetKioskKioskCentresParams{

		timeout: timeout,
	}
}

/*GetKioskKioskCentresParams contains all the parameters to send to the API endpoint
for the get kiosk kiosk centres operation typically these are written to a http.Request
*/
type GetKioskKioskCentresParams struct {
	timeout time.Duration
}

// WriteToRequest writes these params to a swagger request
func (o *GetKioskKioskCentresParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}