package p_a_r_k_i_n_g

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

// NewPostParkingActivitiesParams creates a new PostParkingActivitiesParams object
// with the default values initialized.
func NewPostParkingActivitiesParams() *PostParkingActivitiesParams {
	var ()
	return &PostParkingActivitiesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostParkingActivitiesParamsWithTimeout creates a new PostParkingActivitiesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostParkingActivitiesParamsWithTimeout(timeout time.Duration) *PostParkingActivitiesParams {
	var ()
	return &PostParkingActivitiesParams{

		timeout: timeout,
	}
}

/*PostParkingActivitiesParams contains all the parameters to send to the API endpoint
for the post parking activities operation typically these are written to a http.Request
*/
type PostParkingActivitiesParams struct {

	/*Authorization
	  Staff Access Application token

	*/
	Authorization string
	/*ParkingActivityData
	  Data to create a parking activity.

	*/
	ParkingActivityData *models_secure.ParkingActivityCreateBody

	timeout time.Duration
}

// WithAuthorization adds the authorization to the post parking activities params
func (o *PostParkingActivitiesParams) WithAuthorization(Authorization string) *PostParkingActivitiesParams {
	o.Authorization = Authorization
	return o
}

// WithParkingActivityData adds the parkingActivityData to the post parking activities params
func (o *PostParkingActivitiesParams) WithParkingActivityData(ParkingActivityData *models_secure.ParkingActivityCreateBody) *PostParkingActivitiesParams {
	o.ParkingActivityData = ParkingActivityData
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostParkingActivitiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.ParkingActivityData == nil {
		o.ParkingActivityData = new(models_secure.ParkingActivityCreateBody)
	}

	if err := r.SetBodyParam(o.ParkingActivityData); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
