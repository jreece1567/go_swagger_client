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

	"restclient/models_core"
)

// NewPatchServicesTypesServiceTypeIDParams creates a new PatchServicesTypesServiceTypeIDParams object
// with the default values initialized.
func NewPatchServicesTypesServiceTypeIDParams() *PatchServicesTypesServiceTypeIDParams {
	var ()
	return &PatchServicesTypesServiceTypeIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchServicesTypesServiceTypeIDParamsWithTimeout creates a new PatchServicesTypesServiceTypeIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchServicesTypesServiceTypeIDParamsWithTimeout(timeout time.Duration) *PatchServicesTypesServiceTypeIDParams {
	var ()
	return &PatchServicesTypesServiceTypeIDParams{

		timeout: timeout,
	}
}

/*PatchServicesTypesServiceTypeIDParams contains all the parameters to send to the API endpoint
for the patch services types service type ID operation typically these are written to a http.Request
*/
type PatchServicesTypesServiceTypeIDParams struct {

	/*Authorization
	  Staff Access token.

	*/
	Authorization string
	/*ServiceTypeData
	  Data to update a service-type.

	*/
	ServiceTypeData *models_core.ServiceTypeUpdateBody
	/*ServiceTypeID
	  Service-type identifier. Request to update the service-type with service_type_id.

	*/
	ServiceTypeID int64

	timeout time.Duration
}

// WithAuthorization adds the authorization to the patch services types service type ID params
func (o *PatchServicesTypesServiceTypeIDParams) WithAuthorization(Authorization string) *PatchServicesTypesServiceTypeIDParams {
	o.Authorization = Authorization
	return o
}

// WithServiceTypeData adds the serviceTypeData to the patch services types service type ID params
func (o *PatchServicesTypesServiceTypeIDParams) WithServiceTypeData(ServiceTypeData *models_core.ServiceTypeUpdateBody) *PatchServicesTypesServiceTypeIDParams {
	o.ServiceTypeData = ServiceTypeData
	return o
}

// WithServiceTypeID adds the serviceTypeId to the patch services types service type ID params
func (o *PatchServicesTypesServiceTypeIDParams) WithServiceTypeID(ServiceTypeID int64) *PatchServicesTypesServiceTypeIDParams {
	o.ServiceTypeID = ServiceTypeID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PatchServicesTypesServiceTypeIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.ServiceTypeData == nil {
		o.ServiceTypeData = new(models_core.ServiceTypeUpdateBody)
	}

	if err := r.SetBodyParam(o.ServiceTypeData); err != nil {
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
