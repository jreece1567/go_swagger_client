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

// NewPostPaymentsStripeEventsParams creates a new PostPaymentsStripeEventsParams object
// with the default values initialized.
func NewPostPaymentsStripeEventsParams() *PostPaymentsStripeEventsParams {
	var ()
	return &PostPaymentsStripeEventsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPaymentsStripeEventsParamsWithTimeout creates a new PostPaymentsStripeEventsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPaymentsStripeEventsParamsWithTimeout(timeout time.Duration) *PostPaymentsStripeEventsParams {
	var ()
	return &PostPaymentsStripeEventsParams{

		timeout: timeout,
	}
}

/*PostPaymentsStripeEventsParams contains all the parameters to send to the API endpoint
for the post payments stripe events operation typically these are written to a http.Request
*/
type PostPaymentsStripeEventsParams struct {

	/*StripeEventData
	  Data to create a Stripe event.

	*/
	StripeEventData *models_secure.StripeEventBody

	timeout time.Duration
}

// WithStripeEventData adds the stripeEventData to the post payments stripe events params
func (o *PostPaymentsStripeEventsParams) WithStripeEventData(StripeEventData *models_secure.StripeEventBody) *PostPaymentsStripeEventsParams {
	o.StripeEventData = StripeEventData
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostPaymentsStripeEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.StripeEventData == nil {
		o.StripeEventData = new(models_secure.StripeEventBody)
	}

	if err := r.SetBodyParam(o.StripeEventData); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}