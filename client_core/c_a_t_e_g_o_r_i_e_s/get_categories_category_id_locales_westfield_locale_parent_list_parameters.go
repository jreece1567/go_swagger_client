package c_a_t_e_g_o_r_i_e_s

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

// NewGetCategoriesCategoryIDLocalesWestfieldLocaleParentListParams creates a new GetCategoriesCategoryIDLocalesWestfieldLocaleParentListParams object
// with the default values initialized.
func NewGetCategoriesCategoryIDLocalesWestfieldLocaleParentListParams() *GetCategoriesCategoryIDLocalesWestfieldLocaleParentListParams {
	var ()
	return &GetCategoriesCategoryIDLocalesWestfieldLocaleParentListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCategoriesCategoryIDLocalesWestfieldLocaleParentListParamsWithTimeout creates a new GetCategoriesCategoryIDLocalesWestfieldLocaleParentListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCategoriesCategoryIDLocalesWestfieldLocaleParentListParamsWithTimeout(timeout time.Duration) *GetCategoriesCategoryIDLocalesWestfieldLocaleParentListParams {
	var ()
	return &GetCategoriesCategoryIDLocalesWestfieldLocaleParentListParams{

		timeout: timeout,
	}
}

/*GetCategoriesCategoryIDLocalesWestfieldLocaleParentListParams contains all the parameters to send to the API endpoint
for the get categories category ID locales westfield locale parent list operation typically these are written to a http.Request
*/
type GetCategoriesCategoryIDLocalesWestfieldLocaleParentListParams struct {

	/*CategoryID
	  Category identifier. Filter child categories of the parent with category_id.

	*/
	CategoryID int64
	/*Fields
	  Fields in response. Array that lists the fields requested. If empty or not supplied, all attributes will be returned.

	*/
	Fields []string
	/*UpdatedSince
	  Updated since. Request the categories updated since a specific date and time. ISO-8601 format.

	*/
	UpdatedSince *strfmt.DateTime
	/*WestfieldLocale
	  Locale identifier. Request the categories with the locale westfield_locale.

	*/
	WestfieldLocale string

	timeout time.Duration
}

// WithCategoryID adds the categoryId to the get categories category ID locales westfield locale parent list params
func (o *GetCategoriesCategoryIDLocalesWestfieldLocaleParentListParams) WithCategoryID(CategoryID int64) *GetCategoriesCategoryIDLocalesWestfieldLocaleParentListParams {
	o.CategoryID = CategoryID
	return o
}

// WithFields adds the fields to the get categories category ID locales westfield locale parent list params
func (o *GetCategoriesCategoryIDLocalesWestfieldLocaleParentListParams) WithFields(Fields []string) *GetCategoriesCategoryIDLocalesWestfieldLocaleParentListParams {
	o.Fields = Fields
	return o
}

// WithUpdatedSince adds the updatedSince to the get categories category ID locales westfield locale parent list params
func (o *GetCategoriesCategoryIDLocalesWestfieldLocaleParentListParams) WithUpdatedSince(UpdatedSince *strfmt.DateTime) *GetCategoriesCategoryIDLocalesWestfieldLocaleParentListParams {
	o.UpdatedSince = UpdatedSince
	return o
}

// WithWestfieldLocale adds the westfieldLocale to the get categories category ID locales westfield locale parent list params
func (o *GetCategoriesCategoryIDLocalesWestfieldLocaleParentListParams) WithWestfieldLocale(WestfieldLocale string) *GetCategoriesCategoryIDLocalesWestfieldLocaleParentListParams {
	o.WestfieldLocale = WestfieldLocale
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetCategoriesCategoryIDLocalesWestfieldLocaleParentListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param category_id
	if err := r.SetPathParam("category_id", swag.FormatInt64(o.CategoryID)); err != nil {
		return err
	}

	valuesFields := o.Fields

	joinedFields := swag.JoinByFormat(valuesFields, "csv")
	// query array param fields
	if err := r.SetQueryParam("fields", joinedFields...); err != nil {
		return err
	}

	if o.UpdatedSince != nil {

		// query param updated_since
		var qrUpdatedSince strfmt.DateTime
		if o.UpdatedSince != nil {
			qrUpdatedSince = *o.UpdatedSince
		}
		qUpdatedSince := qrUpdatedSince.String()
		if qUpdatedSince != "" {
			if err := r.SetQueryParam("updated_since", qUpdatedSince); err != nil {
				return err
			}
		}

	}

	// path param westfield_locale
	if err := r.SetPathParam("westfield_locale", o.WestfieldLocale); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
