package p_e_o_p_l_e_m_a_n_a_g_e_m_e_n_t

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

// NewPatchPeoplePersonIDParams creates a new PatchPeoplePersonIDParams object
// with the default values initialized.
func NewPatchPeoplePersonIDParams() *PatchPeoplePersonIDParams {
	var ()
	return &PatchPeoplePersonIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchPeoplePersonIDParamsWithTimeout creates a new PatchPeoplePersonIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchPeoplePersonIDParamsWithTimeout(timeout time.Duration) *PatchPeoplePersonIDParams {
	var ()
	return &PatchPeoplePersonIDParams{

		timeout: timeout,
	}
}

/*PatchPeoplePersonIDParams contains all the parameters to send to the API endpoint
for the patch people person ID operation typically these are written to a http.Request
*/
type PatchPeoplePersonIDParams struct {

	/*Authorization
	  Staff Access token

	*/
	Authorization string
	/*Body*/
	Body *models_secure.ConciergeUpdateAccount
	/*PersonID
	  Identifier of the account

	*/
	PersonID string

	timeout time.Duration
}

// WithAuthorization adds the authorization to the patch people person ID params
func (o *PatchPeoplePersonIDParams) WithAuthorization(Authorization string) *PatchPeoplePersonIDParams {
	o.Authorization = Authorization
	return o
}

// WithBody adds the body to the patch people person ID params
func (o *PatchPeoplePersonIDParams) WithBody(Body *models_secure.ConciergeUpdateAccount) *PatchPeoplePersonIDParams {
	o.Body = Body
	return o
}

// WithPersonID adds the personId to the patch people person ID params
func (o *PatchPeoplePersonIDParams) WithPersonID(PersonID string) *PatchPeoplePersonIDParams {
	o.PersonID = PersonID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PatchPeoplePersonIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.Body == nil {
		o.Body = new(models_secure.ConciergeUpdateAccount)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param person_id
	if err := r.SetPathParam("person_id", o.PersonID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}