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

// NewGetPeoplePersonIDCreditCardsParams creates a new GetPeoplePersonIDCreditCardsParams object
// with the default values initialized.
func NewGetPeoplePersonIDCreditCardsParams() *GetPeoplePersonIDCreditCardsParams {
	var ()
	return &GetPeoplePersonIDCreditCardsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPeoplePersonIDCreditCardsParamsWithTimeout creates a new GetPeoplePersonIDCreditCardsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPeoplePersonIDCreditCardsParamsWithTimeout(timeout time.Duration) *GetPeoplePersonIDCreditCardsParams {
	var ()
	return &GetPeoplePersonIDCreditCardsParams{

		timeout: timeout,
	}
}

/*GetPeoplePersonIDCreditCardsParams contains all the parameters to send to the API endpoint
for the get people person ID credit cards operation typically these are written to a http.Request
*/
type GetPeoplePersonIDCreditCardsParams struct {

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

// WithAuthorization adds the authorization to the get people person ID credit cards params
func (o *GetPeoplePersonIDCreditCardsParams) WithAuthorization(Authorization string) *GetPeoplePersonIDCreditCardsParams {
	o.Authorization = Authorization
	return o
}

// WithPersonID adds the personId to the get people person ID credit cards params
func (o *GetPeoplePersonIDCreditCardsParams) WithPersonID(PersonID string) *GetPeoplePersonIDCreditCardsParams {
	o.PersonID = PersonID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetPeoplePersonIDCreditCardsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
