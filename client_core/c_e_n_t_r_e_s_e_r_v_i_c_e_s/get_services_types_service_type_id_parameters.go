package c_e_n_t_r_e_s_e_r_v_i_c_e_s

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

// NewGetServicesTypesServiceTypeIDParams creates a new GetServicesTypesServiceTypeIDParams object
// with the default values initialized.
func NewGetServicesTypesServiceTypeIDParams() *GetServicesTypesServiceTypeIDParams {
	var ()
	return &GetServicesTypesServiceTypeIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetServicesTypesServiceTypeIDParamsWithTimeout creates a new GetServicesTypesServiceTypeIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetServicesTypesServiceTypeIDParamsWithTimeout(timeout time.Duration) *GetServicesTypesServiceTypeIDParams {
	var ()
	return &GetServicesTypesServiceTypeIDParams{

		timeout: timeout,
	}
}

/*GetServicesTypesServiceTypeIDParams contains all the parameters to send to the API endpoint
for the get services types service type ID operation typically these are written to a http.Request
*/
type GetServicesTypesServiceTypeIDParams struct {

	/*ServiceTypeID
	  Service-type identifier. Request the service-type with service_type_id.

	*/
	ServiceTypeID int64

	timeout time.Duration
}

// WithServiceTypeID adds the serviceTypeId to the get services types service type ID params
func (o *GetServicesTypesServiceTypeIDParams) WithServiceTypeID(ServiceTypeID int64) *GetServicesTypesServiceTypeIDParams {
	o.ServiceTypeID = ServiceTypeID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetServicesTypesServiceTypeIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param service_type_id
	if err := r.SetPathParam("service_type_id", swag.FormatInt64(o.ServiceTypeID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
