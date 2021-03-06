package p_e_o_p_l_e_c_o_n_s_u_m_e_r

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

// NewGetAccountCreditCardsParams creates a new GetAccountCreditCardsParams object
// with the default values initialized.
func NewGetAccountCreditCardsParams() *GetAccountCreditCardsParams {
	var (
		pageDefault    int64 = int64(1)
		perPageDefault int64 = int64(25)
	)
	return &GetAccountCreditCardsParams{
		Page:    &pageDefault,
		PerPage: &perPageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountCreditCardsParamsWithTimeout creates a new GetAccountCreditCardsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAccountCreditCardsParamsWithTimeout(timeout time.Duration) *GetAccountCreditCardsParams {
	var (
		pageDefault    int64 = int64(1)
		perPageDefault int64 = int64(25)
	)
	return &GetAccountCreditCardsParams{
		Page:    &pageDefault,
		PerPage: &perPageDefault,

		timeout: timeout,
	}
}

/*GetAccountCreditCardsParams contains all the parameters to send to the API endpoint
for the get account credit cards operation typically these are written to a http.Request
*/
type GetAccountCreditCardsParams struct {

	/*Authorization
	  Access token.

	*/
	Authorization string
	/*Fields
	  Fields in response. Array that lists the fields requested. If empty or not supplied, all attributes will be returned.

	*/
	Fields []string
	/*Page
	  Results page number. Specific page of results to request when paginating.

	*/
	Page *int64
	/*PerPage
	  Results per page. Number of items per result set when paginating.

	*/
	PerPage *int64

	timeout time.Duration
}

// WithAuthorization adds the authorization to the get account credit cards params
func (o *GetAccountCreditCardsParams) WithAuthorization(Authorization string) *GetAccountCreditCardsParams {
	o.Authorization = Authorization
	return o
}

// WithFields adds the fields to the get account credit cards params
func (o *GetAccountCreditCardsParams) WithFields(Fields []string) *GetAccountCreditCardsParams {
	o.Fields = Fields
	return o
}

// WithPage adds the page to the get account credit cards params
func (o *GetAccountCreditCardsParams) WithPage(Page *int64) *GetAccountCreditCardsParams {
	o.Page = Page
	return o
}

// WithPerPage adds the perPage to the get account credit cards params
func (o *GetAccountCreditCardsParams) WithPerPage(PerPage *int64) *GetAccountCreditCardsParams {
	o.PerPage = PerPage
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountCreditCardsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	valuesFields := o.Fields

	joinedFields := swag.JoinByFormat(valuesFields, "csv")
	// query array param fields
	if err := r.SetQueryParam("fields", joinedFields...); err != nil {
		return err
	}

	if o.Page != nil {

		// query param page
		var qrPage int64
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.PerPage != nil {

		// query param per_page
		var qrPerPage int64
		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt64(qrPerPage)
		if qPerPage != "" {
			if err := r.SetQueryParam("per_page", qPerPage); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
