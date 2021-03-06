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

// NewPostParkingIntegrationsSwarcoCentreIDGaragesParams creates a new PostParkingIntegrationsSwarcoCentreIDGaragesParams object
// with the default values initialized.
func NewPostParkingIntegrationsSwarcoCentreIDGaragesParams() *PostParkingIntegrationsSwarcoCentreIDGaragesParams {
	var ()
	return &PostParkingIntegrationsSwarcoCentreIDGaragesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostParkingIntegrationsSwarcoCentreIDGaragesParamsWithTimeout creates a new PostParkingIntegrationsSwarcoCentreIDGaragesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostParkingIntegrationsSwarcoCentreIDGaragesParamsWithTimeout(timeout time.Duration) *PostParkingIntegrationsSwarcoCentreIDGaragesParams {
	var ()
	return &PostParkingIntegrationsSwarcoCentreIDGaragesParams{

		timeout: timeout,
	}
}

/*PostParkingIntegrationsSwarcoCentreIDGaragesParams contains all the parameters to send to the API endpoint
for the post parking integrations swarco centre ID garages operation typically these are written to a http.Request
*/
type PostParkingIntegrationsSwarcoCentreIDGaragesParams struct {

	/*CentreID
	  Centre identifier. Request the parking garage details with centre_id. Lowercase code name for a specific centre.

	*/
	CentreID string
	/*GarageData
	  Data to update a parking garage.

	*/
	GarageData *models_secure.SwarcoGaragesUpdateBody

	timeout time.Duration
}

// WithCentreID adds the centreId to the post parking integrations swarco centre ID garages params
func (o *PostParkingIntegrationsSwarcoCentreIDGaragesParams) WithCentreID(CentreID string) *PostParkingIntegrationsSwarcoCentreIDGaragesParams {
	o.CentreID = CentreID
	return o
}

// WithGarageData adds the garageData to the post parking integrations swarco centre ID garages params
func (o *PostParkingIntegrationsSwarcoCentreIDGaragesParams) WithGarageData(GarageData *models_secure.SwarcoGaragesUpdateBody) *PostParkingIntegrationsSwarcoCentreIDGaragesParams {
	o.GarageData = GarageData
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostParkingIntegrationsSwarcoCentreIDGaragesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param centre_id
	if err := r.SetPathParam("centre_id", o.CentreID); err != nil {
		return err
	}

	if o.GarageData == nil {
		o.GarageData = new(models_secure.SwarcoGaragesUpdateBody)
	}

	if err := r.SetBodyParam(o.GarageData); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
