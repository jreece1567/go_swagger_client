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

// NewPatchAccountNewslettersSubscribeParams creates a new PatchAccountNewslettersSubscribeParams object
// with the default values initialized.
func NewPatchAccountNewslettersSubscribeParams() *PatchAccountNewslettersSubscribeParams {
	var ()
	return &PatchAccountNewslettersSubscribeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAccountNewslettersSubscribeParamsWithTimeout creates a new PatchAccountNewslettersSubscribeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchAccountNewslettersSubscribeParamsWithTimeout(timeout time.Duration) *PatchAccountNewslettersSubscribeParams {
	var ()
	return &PatchAccountNewslettersSubscribeParams{

		timeout: timeout,
	}
}

/*PatchAccountNewslettersSubscribeParams contains all the parameters to send to the API endpoint
for the patch account newsletters subscribe operation typically these are written to a http.Request
*/
type PatchAccountNewslettersSubscribeParams struct {

	/*NewsletterManageData
	  Newsletter and account details.

	*/
	NewsletterManageData *models_secure.NewslettersManageBody

	timeout time.Duration
}

// WithNewsletterManageData adds the newsletterManageData to the patch account newsletters subscribe params
func (o *PatchAccountNewslettersSubscribeParams) WithNewsletterManageData(NewsletterManageData *models_secure.NewslettersManageBody) *PatchAccountNewslettersSubscribeParams {
	o.NewsletterManageData = NewsletterManageData
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAccountNewslettersSubscribeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.NewsletterManageData == nil {
		o.NewsletterManageData = new(models_secure.NewslettersManageBody)
	}

	if err := r.SetBodyParam(o.NewsletterManageData); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
