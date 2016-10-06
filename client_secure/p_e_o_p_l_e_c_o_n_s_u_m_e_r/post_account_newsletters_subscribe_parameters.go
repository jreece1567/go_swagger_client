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

// NewPostAccountNewslettersSubscribeParams creates a new PostAccountNewslettersSubscribeParams object
// with the default values initialized.
func NewPostAccountNewslettersSubscribeParams() *PostAccountNewslettersSubscribeParams {
	var ()
	return &PostAccountNewslettersSubscribeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAccountNewslettersSubscribeParamsWithTimeout creates a new PostAccountNewslettersSubscribeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAccountNewslettersSubscribeParamsWithTimeout(timeout time.Duration) *PostAccountNewslettersSubscribeParams {
	var ()
	return &PostAccountNewslettersSubscribeParams{

		timeout: timeout,
	}
}

/*PostAccountNewslettersSubscribeParams contains all the parameters to send to the API endpoint
for the post account newsletters subscribe operation typically these are written to a http.Request
*/
type PostAccountNewslettersSubscribeParams struct {

	/*NewsletterSubscribeData
	  Data to create a newsletter subscription.

	*/
	NewsletterSubscribeData *models_secure.NewslettersSubscribeBody

	timeout time.Duration
}

// WithNewsletterSubscribeData adds the newsletterSubscribeData to the post account newsletters subscribe params
func (o *PostAccountNewslettersSubscribeParams) WithNewsletterSubscribeData(NewsletterSubscribeData *models_secure.NewslettersSubscribeBody) *PostAccountNewslettersSubscribeParams {
	o.NewsletterSubscribeData = NewsletterSubscribeData
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostAccountNewslettersSubscribeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.NewsletterSubscribeData == nil {
		o.NewsletterSubscribeData = new(models_secure.NewslettersSubscribeBody)
	}

	if err := r.SetBodyParam(o.NewsletterSubscribeData); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}