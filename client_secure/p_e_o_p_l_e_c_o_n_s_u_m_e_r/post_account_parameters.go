package p_e_o_p_l_e_c_o_n_s_u_m_e_r

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

// NewPostAccountParams creates a new PostAccountParams object
// with the default values initialized.
func NewPostAccountParams() *PostAccountParams {
	var ()
	return &PostAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAccountParamsWithTimeout creates a new PostAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAccountParamsWithTimeout(timeout time.Duration) *PostAccountParams {
	var ()
	return &PostAccountParams{

		timeout: timeout,
	}
}

/*PostAccountParams contains all the parameters to send to the API endpoint
for the post account operation typically these are written to a http.Request
*/
type PostAccountParams struct {

	/*Body*/
	Body *models_secure.NewAccount

	timeout time.Duration
}

// WithBody adds the body to the post account params
func (o *PostAccountParams) WithBody(Body *models_secure.NewAccount) *PostAccountParams {
	o.Body = Body
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models_secure.NewAccount)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}