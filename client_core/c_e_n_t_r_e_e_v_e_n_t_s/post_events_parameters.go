package c_e_n_t_r_e_e_v_e_n_t_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"restclient/models_core"
)

// NewPostEventsParams creates a new PostEventsParams object
// with the default values initialized.
func NewPostEventsParams() *PostEventsParams {
	var ()
	return &PostEventsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostEventsParamsWithTimeout creates a new PostEventsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostEventsParamsWithTimeout(timeout time.Duration) *PostEventsParams {
	var ()
	return &PostEventsParams{

		timeout: timeout,
	}
}

/*PostEventsParams contains all the parameters to send to the API endpoint
for the post events operation typically these are written to a http.Request
*/
type PostEventsParams struct {

	/*Authorization
	  Staff Access token.

	*/
	Authorization string
	/*EventData
	  Data to create an event.

	*/
	EventData *models_core.EventCreateBody

	timeout time.Duration
}

// WithAuthorization adds the authorization to the post events params
func (o *PostEventsParams) WithAuthorization(Authorization string) *PostEventsParams {
	o.Authorization = Authorization
	return o
}

// WithEventData adds the eventData to the post events params
func (o *PostEventsParams) WithEventData(EventData *models_core.EventCreateBody) *PostEventsParams {
	o.EventData = EventData
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.EventData == nil {
		o.EventData = new(models_core.EventCreateBody)
	}

	if err := r.SetBodyParam(o.EventData); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
