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

// NewGetCentresZonesZoneIDParams creates a new GetCentresZonesZoneIDParams object
// with the default values initialized.
func NewGetCentresZonesZoneIDParams() *GetCentresZonesZoneIDParams {
	var ()
	return &GetCentresZonesZoneIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCentresZonesZoneIDParamsWithTimeout creates a new GetCentresZonesZoneIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCentresZonesZoneIDParamsWithTimeout(timeout time.Duration) *GetCentresZonesZoneIDParams {
	var ()
	return &GetCentresZonesZoneIDParams{

		timeout: timeout,
	}
}

/*GetCentresZonesZoneIDParams contains all the parameters to send to the API endpoint
for the get centres zones zone ID operation typically these are written to a http.Request
*/
type GetCentresZonesZoneIDParams struct {

	/*Fields
	  Fields in response. Array that lists the fields requested. If empty or not supplied, all attributes will be returned.

	*/
	Fields []string
	/*ZoneID
	  Zone identifier. Request the centre zone with zone_id.

	*/
	ZoneID int64

	timeout time.Duration
}

// WithFields adds the fields to the get centres zones zone ID params
func (o *GetCentresZonesZoneIDParams) WithFields(Fields []string) *GetCentresZonesZoneIDParams {
	o.Fields = Fields
	return o
}

// WithZoneID adds the zoneId to the get centres zones zone ID params
func (o *GetCentresZonesZoneIDParams) WithZoneID(ZoneID int64) *GetCentresZonesZoneIDParams {
	o.ZoneID = ZoneID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetCentresZonesZoneIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	valuesFields := o.Fields

	joinedFields := swag.JoinByFormat(valuesFields, "csv")
	// query array param fields
	if err := r.SetQueryParam("fields", joinedFields...); err != nil {
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