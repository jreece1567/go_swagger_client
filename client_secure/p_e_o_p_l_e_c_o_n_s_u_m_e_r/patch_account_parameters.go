package p_e_o_p_l_e_c_o_n_s_u_m_e_r

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

// NewPatchAccountParams creates a new PatchAccountParams object
// with the default values initialized.
func NewPatchAccountParams() *PatchAccountParams {
	var ()
	return &PatchAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAccountParamsWithTimeout creates a new PatchAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchAccountParamsWithTimeout(timeout time.Duration) *PatchAccountParams {
	var ()
	return &PatchAccountParams{

		timeout: timeout,
	}
}

/*PatchAccountParams contains all the parameters to send to the API endpoint
for the patch account operation typically these are written to a http.Request
*/
type PatchAccountParams struct {

	/*Authorization
	  Access token

	*/
	Authorization *string
	/*Body*/
	Body *models_secure.UpdateAccount
	/*Country
	  Country code of the account.

	*/
	Country *string
	/*Email
	  Email address of the account

	*/
	Email *string
	/*PersonID
	  Identifier of the account

	*/
	PersonID *string

	timeout time.Duration
}

// WithAuthorization adds the authorization to the patch account params
func (o *PatchAccountParams) WithAuthorization(Authorization *string) *PatchAccountParams {
	o.Authorization = Authorization
	return o
}

// WithBody adds the body to the patch account params
func (o *PatchAccountParams) WithBody(Body *models_secure.UpdateAccount) *PatchAccountParams {
	o.Body = Body
	return o
}

// WithCountry adds the country to the patch account params
func (o *PatchAccountParams) WithCountry(Country *string) *PatchAccountParams {
	o.Country = Country
	return o
}

// WithEmail adds the email to the patch account params
func (o *PatchAccountParams) WithEmail(Email *string) *PatchAccountParams {
	o.Email = Email
	return o
}

// WithPersonID adds the personId to the patch account params
func (o *PatchAccountParams) WithPersonID(PersonID *string) *PatchAccountParams {
	o.PersonID = PersonID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Authorization != nil {

		// header param Authorization
		if err := r.SetHeaderParam("Authorization", *o.Authorization); err != nil {
			return err
		}

	}

	if o.Body == nil {
		o.Body = new(models_secure.UpdateAccount)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if o.Country != nil {

		// query param country
		var qrCountry string
		if o.Country != nil {
			qrCountry = *o.Country
		}
		qCountry := qrCountry
		if qCountry != "" {
			if err := r.SetQueryParam("country", qCountry); err != nil {
				return err
			}
		}

	}

	if o.Email != nil {

		// query param email
		var qrEmail string
		if o.Email != nil {
			qrEmail = *o.Email
		}
		qEmail := qrEmail
		if qEmail != "" {
			if err := r.SetQueryParam("email", qEmail); err != nil {
				return err
			}
		}

	}

	if o.PersonID != nil {

		// query param person_id
		var qrPersonID string
		if o.PersonID != nil {
			qrPersonID = *o.PersonID
		}
		qPersonID := qrPersonID
		if qPersonID != "" {
			if err := r.SetQueryParam("person_id", qPersonID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}