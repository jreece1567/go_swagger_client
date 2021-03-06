package c_e_n_t_r_e_p_a_r_k_i_n_g_r_a_t_e_s

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

// NewPatchCentresParkingsParkingIDParams creates a new PatchCentresParkingsParkingIDParams object
// with the default values initialized.
func NewPatchCentresParkingsParkingIDParams() *PatchCentresParkingsParkingIDParams {
	var ()
	return &PatchCentresParkingsParkingIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchCentresParkingsParkingIDParamsWithTimeout creates a new PatchCentresParkingsParkingIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchCentresParkingsParkingIDParamsWithTimeout(timeout time.Duration) *PatchCentresParkingsParkingIDParams {
	var ()
	return &PatchCentresParkingsParkingIDParams{

		timeout: timeout,
	}
}

/*PatchCentresParkingsParkingIDParams contains all the parameters to send to the API endpoint
for the patch centres parkings parking ID operation typically these are written to a http.Request
*/
type PatchCentresParkingsParkingIDParams struct {

	/*Authorization
	  Staff Access token.

	*/
	Authorization string
	/*ParkingData
	  Data to update a centre parking rates.

	*/
	ParkingData *models_core.ParkingUpdateBody
	/*ParkingID
	  Parking identifier. Request the centre parking rates with parking_id.

	*/
	ParkingID int64

	timeout time.Duration
}

// WithAuthorization adds the authorization to the patch centres parkings parking ID params
func (o *PatchCentresParkingsParkingIDParams) WithAuthorization(Authorization string) *PatchCentresParkingsParkingIDParams {
	o.Authorization = Authorization
	return o
}

// WithParkingData adds the parkingData to the patch centres parkings parking ID params
func (o *PatchCentresParkingsParkingIDParams) WithParkingData(ParkingData *models_core.ParkingUpdateBody) *PatchCentresParkingsParkingIDParams {
	o.ParkingData = ParkingData
	return o
}

// WithParkingID adds the parkingId to the patch centres parkings parking ID params
func (o *PatchCentresParkingsParkingIDParams) WithParkingID(ParkingID int64) *PatchCentresParkingsParkingIDParams {
	o.ParkingID = ParkingID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PatchCentresParkingsParkingIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.ParkingData == nil {
		o.ParkingData = new(models_core.ParkingUpdateBody)
	}

	if err := r.SetBodyParam(o.ParkingData); err != nil {
		return err
	}

	// path param parking_id
	if err := r.SetPathParam("parking_id", swag.FormatInt64(o.ParkingID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
