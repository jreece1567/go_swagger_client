package c_e_n_t_r_e_s_e_r_v_i_c_e_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"restclient/models_core"
)

// NewPostServicesTypesParams creates a new PostServicesTypesParams object
// with the default values initialized.
func NewPostServicesTypesParams() *PostServicesTypesParams {
	var ()
	return &PostServicesTypesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostServicesTypesParamsWithTimeout creates a new PostServicesTypesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostServicesTypesParamsWithTimeout(timeout time.Duration) *PostServicesTypesParams {
	var ()
	return &PostServicesTypesParams{

		timeout: timeout,
	}
}

/*PostServicesTypesParams contains all the parameters to send to the API endpoint
for the post services types operation typically these are written to a http.Request
*/
type PostServicesTypesParams struct {

	/*Authorization
	  Staff Access token.

	*/
	Authorization string
	/*ServiceTypeData
	  Data to create a service-type.

	*/
	ServiceTypeData *models_core.ServiceTypeCreateBody

	timeout time.Duration
}

// WithAuthorization adds the authorization to the post services types params
func (o *PostServicesTypesParams) WithAuthorization(Authorization string) *PostServicesTypesParams {
	o.Authorization = Authorization
	return o
}

// WithServiceTypeData adds the serviceTypeData to the post services types params
func (o *PostServicesTypesParams) WithServiceTypeData(ServiceTypeData *models_core.ServiceTypeCreateBody) *PostServicesTypesParams {
	o.ServiceTypeData = ServiceTypeData
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostServicesTypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.ServiceTypeData == nil {
		o.ServiceTypeData = new(models_core.ServiceTypeCreateBody)
	}

	if err := r.SetBodyParam(o.ServiceTypeData); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
