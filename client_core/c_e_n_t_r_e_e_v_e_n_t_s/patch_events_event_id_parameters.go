package c_e_n_t_r_e_e_v_e_n_t_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"restclient/models_core"
)

// NewPatchEventsEventIDParams creates a new PatchEventsEventIDParams object
// with the default values initialized.
func NewPatchEventsEventIDParams() *PatchEventsEventIDParams {
	var ()
	return &PatchEventsEventIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchEventsEventIDParamsWithTimeout creates a new PatchEventsEventIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchEventsEventIDParamsWithTimeout(timeout time.Duration) *PatchEventsEventIDParams {
	var ()
	return &PatchEventsEventIDParams{

		timeout: timeout,
	}
}

/*PatchEventsEventIDParams contains all the parameters to send to the API endpoint
for the patch events event ID operation typically these are written to a http.Request
*/
type PatchEventsEventIDParams struct {

	/*Authorization
	  Staff Access token.

	*/
	Authorization string
	/*EventData
	  Data to update an event.

	*/
	EventData *models_core.EventUpdateBody
	/*EventID
	  Event identifier. Request the event with event_id.

	*/
	EventID int64

	timeout time.Duration
}

// WithAuthorization adds the authorization to the patch events event ID params
func (o *PatchEventsEventIDParams) WithAuthorization(Authorization string) *PatchEventsEventIDParams {
	o.Authorization = Authorization
	return o
}

// WithEventData adds the eventData to the patch events event ID params
func (o *PatchEventsEventIDParams) WithEventData(EventData *models_core.EventUpdateBody) *PatchEventsEventIDParams {
	o.EventData = EventData
	return o
}

// WithEventID adds the eventId to the patch events event ID params
func (o *PatchEventsEventIDParams) WithEventID(EventID int64) *PatchEventsEventIDParams {
	o.EventID = EventID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PatchEventsEventIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.EventData == nil {
		o.EventData = new(models_core.EventUpdateBody)
	}

	if err := r.SetBodyParam(o.EventData); err != nil {
		return err
	}

	// path param event_id
	if err := r.SetPathParam("event_id", swag.FormatInt64(o.EventID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
