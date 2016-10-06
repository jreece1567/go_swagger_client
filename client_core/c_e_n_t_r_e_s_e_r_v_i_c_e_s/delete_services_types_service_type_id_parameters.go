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

// NewDeleteServicesTypesServiceTypeIDParams creates a new DeleteServicesTypesServiceTypeIDParams object
// with the default values initialized.
func NewDeleteServicesTypesServiceTypeIDParams() *DeleteServicesTypesServiceTypeIDParams {
	var ()
	return &DeleteServicesTypesServiceTypeIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteServicesTypesServiceTypeIDParamsWithTimeout creates a new DeleteServicesTypesServiceTypeIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteServicesTypesServiceTypeIDParamsWithTimeout(timeout time.Duration) *DeleteServicesTypesServiceTypeIDParams {
	var ()
	return &DeleteServicesTypesServiceTypeIDParams{

		timeout: timeout,
	}
}

/*DeleteServicesTypesServiceTypeIDParams contains all the parameters to send to the API endpoint
for the delete services types service type ID operation typically these are written to a http.Request
*/
type DeleteServicesTypesServiceTypeIDParams struct {

	/*Authorization
	  Staff Access token.

	*/
	Authorization string
	/*ServiceTypeID
	  Service-type identifier. Request to delete the service-type with service_type_id.

	*/
	ServiceTypeID int64

	timeout time.Duration
}

// WithAuthorization adds the authorization to the delete services types service type ID params
func (o *DeleteServicesTypesServiceTypeIDParams) WithAuthorization(Authorization string) *DeleteServicesTypesServiceTypeIDParams {
	o.Authorization = Authorization
	return o
}

// WithServiceTypeID adds the serviceTypeId to the delete services types service type ID params
func (o *DeleteServicesTypesServiceTypeIDParams) WithServiceTypeID(ServiceTypeID int64) *DeleteServicesTypesServiceTypeIDParams {
	o.ServiceTypeID = ServiceTypeID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteServicesTypesServiceTypeIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param service_type_id
	if err := r.SetPathParam("service_type_id", swag.FormatInt64(o.ServiceTypeID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}