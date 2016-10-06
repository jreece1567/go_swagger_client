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

// NewPostServicesParams creates a new PostServicesParams object
// with the default values initialized.
func NewPostServicesParams() *PostServicesParams {
	var ()
	return &PostServicesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostServicesParamsWithTimeout creates a new PostServicesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostServicesParamsWithTimeout(timeout time.Duration) *PostServicesParams {
	var ()
	return &PostServicesParams{

		timeout: timeout,
	}
}

/*PostServicesParams contains all the parameters to send to the API endpoint
for the post services operation typically these are written to a http.Request
*/
type PostServicesParams struct {

	/*Authorization
	  Staff Access token.

	*/
	Authorization string
	/*ServiceData
	  Data to create a service.

	*/
	ServiceData *models_core.ServiceCreateBody

	timeout time.Duration
}

// WithAuthorization adds the authorization to the post services params
func (o *PostServicesParams) WithAuthorization(Authorization string) *PostServicesParams {
	o.Authorization = Authorization
	return o
}

// WithServiceData adds the serviceData to the post services params
func (o *PostServicesParams) WithServiceData(ServiceData *models_core.ServiceCreateBody) *PostServicesParams {
	o.ServiceData = ServiceData
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostServicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.ServiceData == nil {
		o.ServiceData = new(models_core.ServiceCreateBody)
	}

	if err := r.SetBodyParam(o.ServiceData); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}