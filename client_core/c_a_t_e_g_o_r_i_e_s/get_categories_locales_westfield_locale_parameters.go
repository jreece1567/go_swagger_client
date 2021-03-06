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

// NewGetCategoriesLocalesWestfieldLocaleParams creates a new GetCategoriesLocalesWestfieldLocaleParams object
// with the default values initialized.
func NewGetCategoriesLocalesWestfieldLocaleParams() *GetCategoriesLocalesWestfieldLocaleParams {
	var ()
	return &GetCategoriesLocalesWestfieldLocaleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCategoriesLocalesWestfieldLocaleParamsWithTimeout creates a new GetCategoriesLocalesWestfieldLocaleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCategoriesLocalesWestfieldLocaleParamsWithTimeout(timeout time.Duration) *GetCategoriesLocalesWestfieldLocaleParams {
	var ()
	return &GetCategoriesLocalesWestfieldLocaleParams{

		timeout: timeout,
	}
}

/*GetCategoriesLocalesWestfieldLocaleParams contains all the parameters to send to the API endpoint
for the get categories locales westfield locale operation typically these are written to a http.Request
*/
type GetCategoriesLocalesWestfieldLocaleParams struct {

	/*CategoryIds
	  Category identifiers. Request categories with these category_ids.

	*/
	CategoryIds []int64
	/*Fields
	  Fields in response. Array that lists the fields requested. If empty or not supplied, all attributes will be returned.

	*/
	Fields []string
	/*Statuses
	  Statuses. Requests locales that include one of the listed statuses.

	*/
	Statuses []string
	/*WestfieldLocale
	  Locale identifier. Request the locale with westfield_locale.

	*/
	WestfieldLocale string

	timeout time.Duration
}

// WithCategoryIds adds the categoryIds to the get categories locales westfield locale params
func (o *GetCategoriesLocalesWestfieldLocaleParams) WithCategoryIds(CategoryIds []int64) *GetCategoriesLocalesWestfieldLocaleParams {
	o.CategoryIds = CategoryIds
	return o
}

// WithFields adds the fields to the get categories locales westfield locale params
func (o *GetCategoriesLocalesWestfieldLocaleParams) WithFields(Fields []string) *GetCategoriesLocalesWestfieldLocaleParams {
	o.Fields = Fields
	return o
}

// WithStatuses adds the statuses to the get categories locales westfield locale params
func (o *GetCategoriesLocalesWestfieldLocaleParams) WithStatuses(Statuses []string) *GetCategoriesLocalesWestfieldLocaleParams {
	o.Statuses = Statuses
	return o
}

// WithWestfieldLocale adds the westfieldLocale to the get categories locales westfield locale params
func (o *GetCategoriesLocalesWestfieldLocaleParams) WithWestfieldLocale(WestfieldLocale string) *GetCategoriesLocalesWestfieldLocaleParams {
	o.WestfieldLocale = WestfieldLocale
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetCategoriesLocalesWestfieldLocaleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	var valuesCategoryIds []string
	for _, v := range o.CategoryIds {
		valuesCategoryIds = append(valuesCategoryIds, swag.FormatInt64(v))
	}

	joinedCategoryIds := swag.JoinByFormat(valuesCategoryIds, "csv")
	// query array param category_ids
	if err := r.SetQueryParam("category_ids", joinedCategoryIds...); err != nil {
		return err
	}

	valuesFields := o.Fields

	joinedFields := swag.JoinByFormat(valuesFields, "csv")
	// query array param fields
	if err := r.SetQueryParam("fields", joinedFields...); err != nil {
		return err
	}

	valuesStatuses := o.Statuses

	joinedStatuses := swag.JoinByFormat(valuesStatuses, "csv")
	// query array param statuses
	if err := r.SetQueryParam("statuses", joinedStatuses...); err != nil {
		return err
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
