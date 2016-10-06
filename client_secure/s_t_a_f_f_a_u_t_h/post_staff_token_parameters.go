package s_t_a_f_f_a_u_t_h

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

// NewPostStaffTokenParams creates a new PostStaffTokenParams object
// with the default values initialized.
func NewPostStaffTokenParams() *PostStaffTokenParams {
	var ()
	return &PostStaffTokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostStaffTokenParamsWithTimeout creates a new PostStaffTokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostStaffTokenParamsWithTimeout(timeout time.Duration) *PostStaffTokenParams {
	var ()
	return &PostStaffTokenParams{

		timeout: timeout,
	}
}

/*PostStaffTokenParams contains all the parameters to send to the API endpoint
for the post staff token operation typically these are written to a http.Request
*/
type PostStaffTokenParams struct {

	/*Body*/
	Body *models_secure.TokenRequestBody

	timeout time.Duration
}

// WithBody adds the body to the post staff token params
func (o *PostStaffTokenParams) WithBody(Body *models_secure.TokenRequestBody) *PostStaffTokenParams {
	o.Body = Body
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostStaffTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models_secure.TokenRequestBody)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
