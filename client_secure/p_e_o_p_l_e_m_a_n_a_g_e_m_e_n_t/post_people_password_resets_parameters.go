package p_e_o_p_l_e_m_a_n_a_g_e_m_e_n_t

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

// NewPostPeoplePasswordResetsParams creates a new PostPeoplePasswordResetsParams object
// with the default values initialized.
func NewPostPeoplePasswordResetsParams() *PostPeoplePasswordResetsParams {
	var ()
	return &PostPeoplePasswordResetsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPeoplePasswordResetsParamsWithTimeout creates a new PostPeoplePasswordResetsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPeoplePasswordResetsParamsWithTimeout(timeout time.Duration) *PostPeoplePasswordResetsParams {
	var ()
	return &PostPeoplePasswordResetsParams{

		timeout: timeout,
	}
}

/*PostPeoplePasswordResetsParams contains all the parameters to send to the API endpoint
for the post people password resets operation typically these are written to a http.Request
*/
type PostPeoplePasswordResetsParams struct {

	/*Body*/
	Body *models_secure.InitiatePasswordResetBody

	timeout time.Duration
}

// WithBody adds the body to the post people password resets params
func (o *PostPeoplePasswordResetsParams) WithBody(Body *models_secure.InitiatePasswordResetBody) *PostPeoplePasswordResetsParams {
	o.Body = Body
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostPeoplePasswordResetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models_secure.InitiatePasswordResetBody)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
