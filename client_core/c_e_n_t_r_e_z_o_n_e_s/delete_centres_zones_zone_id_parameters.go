package c_e_n_t_r_e_z_o_n_e_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteCentresZonesZoneIDParams creates a new DeleteCentresZonesZoneIDParams object
// with the default values initialized.
func NewDeleteCentresZonesZoneIDParams() *DeleteCentresZonesZoneIDParams {
	var ()
	return &DeleteCentresZonesZoneIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCentresZonesZoneIDParamsWithTimeout creates a new DeleteCentresZonesZoneIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCentresZonesZoneIDParamsWithTimeout(timeout time.Duration) *DeleteCentresZonesZoneIDParams {
	var ()
	return &DeleteCentresZonesZoneIDParams{

		timeout: timeout,
	}
}

/*DeleteCentresZonesZoneIDParams contains all the parameters to send to the API endpoint
for the delete centres zones zone ID operation typically these are written to a http.Request
*/
type DeleteCentresZonesZoneIDParams struct {

	/*Authorization
	  Staff Access token.

	*/
	Authorization string
	/*ZoneID
	  Zone identifier. Request to delete the centre zone with zone_id.

	*/
	ZoneID int64

	timeout time.Duration
}

// WithAuthorization adds the authorization to the delete centres zones zone ID params
func (o *DeleteCentresZonesZoneIDParams) WithAuthorization(Authorization string) *DeleteCentresZonesZoneIDParams {
	o.Authorization = Authorization
	return o
}

// WithZoneID adds the zoneId to the delete centres zones zone ID params
func (o *DeleteCentresZonesZoneIDParams) WithZoneID(ZoneID int64) *DeleteCentresZonesZoneIDParams {
	o.ZoneID = ZoneID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCentresZonesZoneIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param zone_id
	if err := r.SetPathParam("zone_id", swag.FormatInt64(o.ZoneID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}