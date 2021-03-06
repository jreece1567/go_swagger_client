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

// NewDeleteCategoriesCategoryIDLocalesWestfieldLocaleParams creates a new DeleteCategoriesCategoryIDLocalesWestfieldLocaleParams object
// with the default values initialized.
func NewDeleteCategoriesCategoryIDLocalesWestfieldLocaleParams() *DeleteCategoriesCategoryIDLocalesWestfieldLocaleParams {
	var ()
	return &DeleteCategoriesCategoryIDLocalesWestfieldLocaleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCategoriesCategoryIDLocalesWestfieldLocaleParamsWithTimeout creates a new DeleteCategoriesCategoryIDLocalesWestfieldLocaleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCategoriesCategoryIDLocalesWestfieldLocaleParamsWithTimeout(timeout time.Duration) *DeleteCategoriesCategoryIDLocalesWestfieldLocaleParams {
	var ()
	return &DeleteCategoriesCategoryIDLocalesWestfieldLocaleParams{

		timeout: timeout,
	}
}

/*DeleteCategoriesCategoryIDLocalesWestfieldLocaleParams contains all the parameters to send to the API endpoint
for the delete categories category ID locales westfield locale operation typically these are written to a http.Request
*/
type DeleteCategoriesCategoryIDLocalesWestfieldLocaleParams struct {

	/*Authorization
	  Staff Access token.

	*/
	Authorization string
	/*CategoryID
	  Category identifier. Request to update the locale with category_id.

	*/
	CategoryID int64
	/*WestfieldLocale
	  Locale identifier. Request to update the locale with westfield_locale.

	*/
	WestfieldLocale string

	timeout time.Duration
}

// WithAuthorization adds the authorization to the delete categories category ID locales westfield locale params
func (o *DeleteCategoriesCategoryIDLocalesWestfieldLocaleParams) WithAuthorization(Authorization string) *DeleteCategoriesCategoryIDLocalesWestfieldLocaleParams {
	o.Authorization = Authorization
	return o
}

// WithCategoryID adds the categoryId to the delete categories category ID locales westfield locale params
func (o *DeleteCategoriesCategoryIDLocalesWestfieldLocaleParams) WithCategoryID(CategoryID int64) *DeleteCategoriesCategoryIDLocalesWestfieldLocaleParams {
	o.CategoryID = CategoryID
	return o
}

// WithWestfieldLocale adds the westfieldLocale to the delete categories category ID locales westfield locale params
func (o *DeleteCategoriesCategoryIDLocalesWestfieldLocaleParams) WithWestfieldLocale(WestfieldLocale string) *DeleteCategoriesCategoryIDLocalesWestfieldLocaleParams {
	o.WestfieldLocale = WestfieldLocale
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCategoriesCategoryIDLocalesWestfieldLocaleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param category_id
	if err := r.SetPathParam("category_id", swag.FormatInt64(o.CategoryID)); err != nil {
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
