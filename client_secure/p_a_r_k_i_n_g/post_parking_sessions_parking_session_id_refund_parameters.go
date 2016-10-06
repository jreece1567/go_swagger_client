package p_a_r_k_i_n_g

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"restclient/models_secure"
)

// NewPostParkingSessionsParkingSessionIDRefundParams creates a new PostParkingSessionsParkingSessionIDRefundParams object
// with the default values initialized.
func NewPostParkingSessionsParkingSessionIDRefundParams() *PostParkingSessionsParkingSessionIDRefundParams {
	var ()
	return &PostParkingSessionsParkingSessionIDRefundParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostParkingSessionsParkingSessionIDRefundParamsWithTimeout creates a new PostParkingSessionsParkingSessionIDRefundParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostParkingSessionsParkingSessionIDRefundParamsWithTimeout(timeout time.Duration) *PostParkingSessionsParkingSessionIDRefundParams {
	var ()
	return &PostParkingSessionsParkingSessionIDRefundParams{

		timeout: timeout,
	}
}

/*PostParkingSessionsParkingSessionIDRefundParams contains all the parameters to send to the API endpoint
for the post parking sessions parking session ID refund operation typically these are written to a http.Request
*/
type PostParkingSessionsParkingSessionIDRefundParams struct {

	/*Data
	  Data to process the refund of a parking session.

	*/
	Data *models_secure.ParkingSessionRefundBody
	/*ParkingSessionID
	  Parking session identifier.

	*/
	ParkingSessionID int64

	timeout time.Duration
}

// WithData adds the data to the post parking sessions parking session ID refund params
func (o *PostParkingSessionsParkingSessionIDRefundParams) WithData(Data *models_secure.ParkingSessionRefundBody) *PostParkingSessionsParkingSessionIDRefundParams {
	o.Data = Data
	return o
}

// WithParkingSessionID adds the parkingSessionId to the post parking sessions parking session ID refund params
func (o *PostParkingSessionsParkingSessionIDRefundParams) WithParkingSessionID(ParkingSessionID int64) *PostParkingSessionsParkingSessionIDRefundParams {
	o.ParkingSessionID = ParkingSessionID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostParkingSessionsParkingSessionIDRefundParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Data == nil {
		o.Data = new(models_secure.ParkingSessionRefundBody)
	}

	if err := r.SetBodyParam(o.Data); err != nil {
		return err
	}

	// path param parking_session_id
	if err := r.SetPathParam("parking_session_id", swag.FormatInt64(o.ParkingSessionID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
