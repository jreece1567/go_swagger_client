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

	"restclient/models_core"
)

// NewPatchCategoriesCategoryIDParams creates a new PatchCategoriesCategoryIDParams object
// with the default values initialized.
func NewPatchCategoriesCategoryIDParams() *PatchCategoriesCategoryIDParams {
	var ()
	return &PatchCategoriesCategoryIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchCategoriesCategoryIDParamsWithTimeout creates a new PatchCategoriesCategoryIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchCategoriesCategoryIDParamsWithTimeout(timeout time.Duration) *PatchCategoriesCategoryIDParams {
	var ()
	return &PatchCategoriesCategoryIDParams{

		timeout: timeout,
	}
}

/*PatchCategoriesCategoryIDParams contains all the parameters to send to the API endpoint
for the patch categories category ID operation typically these are written to a http.Request
*/
type PatchCategoriesCategoryIDParams struct {

	/*Authorization
	  Staff Access token.

	*/
	Authorization string
	/*CategoryData
	  Data to update a category.

	*/
	CategoryData *models_core.CategoryUpdateBody
	/*CategoryID
	  Category identifier. Request to update the category with category_id.

	*/
	CategoryID int64

	timeout time.Duration
}

// WithAuthorization adds the authorization to the patch categories category ID params
func (o *PatchCategoriesCategoryIDParams) WithAuthorization(Authorization string) *PatchCategoriesCategoryIDParams {
	o.Authorization = Authorization
	return o
}

// WithCategoryData adds the categoryData to the patch categories category ID params
func (o *PatchCategoriesCategoryIDParams) WithCategoryData(CategoryData *models_core.CategoryUpdateBody) *PatchCategoriesCategoryIDParams {
	o.CategoryData = CategoryData
	return o
}

// WithCategoryID adds the categoryId to the patch categories category ID params
func (o *PatchCategoriesCategoryIDParams) WithCategoryID(CategoryID int64) *PatchCategoriesCategoryIDParams {
	o.CategoryID = CategoryID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PatchCategoriesCategoryIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.CategoryData == nil {
		o.CategoryData = new(models_core.CategoryUpdateBody)
	}

	if err := r.SetBodyParam(o.CategoryData); err != nil {
		return err
	}

	// path param category_id
	if err := r.SetPathParam("category_id", swag.FormatInt64(o.CategoryID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}