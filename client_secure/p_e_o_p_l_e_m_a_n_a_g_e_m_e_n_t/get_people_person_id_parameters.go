package p_e_o_p_l_e_m_a_n_a_g_e_m_e_n_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPeoplePersonIDParams creates a new GetPeoplePersonIDParams object
// with the default values initialized.
func NewGetPeoplePersonIDParams() *GetPeoplePersonIDParams {
	var ()
	return &GetPeoplePersonIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPeoplePersonIDParamsWithTimeout creates a new GetPeoplePersonIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPeoplePersonIDParamsWithTimeout(timeout time.Duration) *GetPeoplePersonIDParams {
	var ()
	return &GetPeoplePersonIDParams{

		timeout: timeout,
	}
}

/*GetPeoplePersonIDParams contains all the parameters to send to the API endpoint
for the get people person ID operation typically these are written to a http.Request
*/
type GetPeoplePersonIDParams struct {

	/*Authorization
	  Staff access token

	*/
	Authorization string
	/*PersonID
	  Identifier of the account

	*/
	PersonID string

	timeout time.Duration
}

// WithAuthorization adds the authorization to the get people person ID params
func (o *GetPeoplePersonIDParams) WithAuthorization(Authorization string) *GetPeoplePersonIDParams {
	o.Authorization = Authorization
	return o
}

// WithPersonID adds the personId to the get people person ID params
func (o *GetPeoplePersonIDParams) WithPersonID(PersonID string) *GetPeoplePersonIDParams {
	o.PersonID = PersonID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetPeoplePersonIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
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
