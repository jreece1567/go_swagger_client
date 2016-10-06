package p_e_o_p_l_e_c_o_n_s_u_m_e_r

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAccountNewslettersParams creates a new GetAccountNewslettersParams object
// with the default values initialized.
func NewGetAccountNewslettersParams() *GetAccountNewslettersParams {
	var ()
	return &GetAccountNewslettersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountNewslettersParamsWithTimeout creates a new GetAccountNewslettersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAccountNewslettersParamsWithTimeout(timeout time.Duration) *GetAccountNewslettersParams {
	var ()
	return &GetAccountNewslettersParams{

		timeout: timeout,
	}
}

/*GetAccountNewslettersParams contains all the parameters to send to the API endpoint
for the get account newsletters operation typically these are written to a http.Request
*/
type GetAccountNewslettersParams struct {

	/*Country
	  Country code of the account (required if person_id is not present).

	*/
	Country *string
	/*Email
	  Email of the account (required if person_id is not present).

	*/
	Email *string
	/*NewsletterAccessToken
	  Newsletter access token. Token to authorize access to newsletter information

	*/
	NewsletterAccessToken *string
	/*PersonID
	  Person Identifier

	*/
	PersonID *string

	timeout time.Duration
}

// WithCountry adds the country to the get account newsletters params
func (o *GetAccountNewslettersParams) WithCountry(Country *string) *GetAccountNewslettersParams {
	o.Country = Country
	return o
}

// WithEmail adds the email to the get account newsletters params
func (o *GetAccountNewslettersParams) WithEmail(Email *string) *GetAccountNewslettersParams {
	o.Email = Email
	return o
}

// WithNewsletterAccessToken adds the newsletterAccessToken to the get account newsletters params
func (o *GetAccountNewslettersParams) WithNewsletterAccessToken(NewsletterAccessToken *string) *GetAccountNewslettersParams {
	o.NewsletterAccessToken = NewsletterAccessToken
	return o
}

// WithPersonID adds the personId to the get account newsletters params
func (o *GetAccountNewslettersParams) WithPersonID(PersonID *string) *GetAccountNewslettersParams {
	o.PersonID = PersonID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountNewslettersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

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

	if o.NewsletterAccessToken != nil {

		// query param newsletter_access_token
		var qrNewsletterAccessToken string
		if o.NewsletterAccessToken != nil {
			qrNewsletterAccessToken = *o.NewsletterAccessToken
		}
		qNewsletterAccessToken := qrNewsletterAccessToken
		if qNewsletterAccessToken != "" {
			if err := r.SetQueryParam("newsletter_access_token", qNewsletterAccessToken); err != nil {
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
