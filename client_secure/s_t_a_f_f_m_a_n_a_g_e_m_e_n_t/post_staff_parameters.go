package s_t_a_f_f_m_a_n_a_g_e_m_e_n_t

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

// NewPostStaffParams creates a new PostStaffParams object
// with the default values initialized.
func NewPostStaffParams() *PostStaffParams {
	var ()
	return &PostStaffParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostStaffParamsWithTimeout creates a new PostStaffParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostStaffParamsWithTimeout(timeout time.Duration) *PostStaffParams {
	var ()
	return &PostStaffParams{

		timeout: timeout,
	}
}

/*PostStaffParams contains all the parameters to send to the API endpoint
for the post staff operation typically these are written to a http.Request
*/
type PostStaffParams struct {

	/*Authorization
	  Staff Access token.

	*/
	Authorization string
	/*Staff*/
	Staff *models_secure.Staff

	timeout time.Duration
}

// WithAuthorization adds the authorization to the post staff params
func (o *PostStaffParams) WithAuthorization(Authorization string) *PostStaffParams {
	o.Authorization = Authorization
	return o
}

// WithStaff adds the staff to the post staff params
func (o *PostStaffParams) WithStaff(Staff *models_secure.Staff) *PostStaffParams {
	o.Staff = Staff
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostStaffParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.Staff == nil {
		o.Staff = new(models_secure.Staff)
	}

	if err := r.SetBodyParam(o.Staff); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
