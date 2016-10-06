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

// NewGetCategoriesGroupsParams creates a new GetCategoriesGroupsParams object
// with the default values initialized.
func NewGetCategoriesGroupsParams() *GetCategoriesGroupsParams {
	var (
		pageDefault    int64 = int64(1)
		perPageDefault int64 = int64(25)
	)
	return &GetCategoriesGroupsParams{
		Page:    &pageDefault,
		PerPage: &perPageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCategoriesGroupsParamsWithTimeout creates a new GetCategoriesGroupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCategoriesGroupsParamsWithTimeout(timeout time.Duration) *GetCategoriesGroupsParams {
	var (
		pageDefault    int64 = int64(1)
		perPageDefault int64 = int64(25)
	)
	return &GetCategoriesGroupsParams{
		Page:    &pageDefault,
		PerPage: &perPageDefault,

		timeout: timeout,
	}
}

/*GetCategoriesGroupsParams contains all the parameters to send to the API endpoint
for the get categories groups operation typically these are written to a http.Request
*/
type GetCategoriesGroupsParams struct {

	/*CategoryIds
	  Category identifiers to use as a filter.

	*/
	CategoryIds []int64
	/*Country
	  Country filter.

	*/
	Country *string
	/*GroupIds
	  Group identifiers to use as a filter.

	*/
	GroupIds []int64
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

// WithCategoryIds adds the categoryIds to the get categories groups params
func (o *GetCategoriesGroupsParams) WithCategoryIds(CategoryIds []int64) *GetCategoriesGroupsParams {
	o.CategoryIds = CategoryIds
	return o
}

// WithCountry adds the country to the get categories groups params
func (o *GetCategoriesGroupsParams) WithCountry(Country *string) *GetCategoriesGroupsParams {
	o.Country = Country
	return o
}

// WithGroupIds adds the groupIds to the get categories groups params
func (o *GetCategoriesGroupsParams) WithGroupIds(GroupIds []int64) *GetCategoriesGroupsParams {
	o.GroupIds = GroupIds
	return o
}

// WithPage adds the page to the get categories groups params
func (o *GetCategoriesGroupsParams) WithPage(Page *int64) *GetCategoriesGroupsParams {
	o.Page = Page
	return o
}

// WithPerPage adds the perPage to the get categories groups params
func (o *GetCategoriesGroupsParams) WithPerPage(PerPage *int64) *GetCategoriesGroupsParams {
	o.PerPage = PerPage
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetCategoriesGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	var valuesGroupIds []string
	for _, v := range o.GroupIds {
		valuesGroupIds = append(valuesGroupIds, swag.FormatInt64(v))
	}

	joinedGroupIds := swag.JoinByFormat(valuesGroupIds, "csv")
	// query array param group_ids
	if err := r.SetQueryParam("group_ids", joinedGroupIds...); err != nil {
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