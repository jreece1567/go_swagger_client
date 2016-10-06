package c_a_t_e_g_o_r_i_e_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"restclient/models_core"
)

// NewPostCategoriesGroupsParams creates a new PostCategoriesGroupsParams object
// with the default values initialized.
func NewPostCategoriesGroupsParams() *PostCategoriesGroupsParams {
	var ()
	return &PostCategoriesGroupsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCategoriesGroupsParamsWithTimeout creates a new PostCategoriesGroupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCategoriesGroupsParamsWithTimeout(timeout time.Duration) *PostCategoriesGroupsParams {
	var ()
	return &PostCategoriesGroupsParams{

		timeout: timeout,
	}
}

/*PostCategoriesGroupsParams contains all the parameters to send to the API endpoint
for the post categories groups operation typically these are written to a http.Request
*/
type PostCategoriesGroupsParams struct {

	/*Authorization
	  Staff Access token.

	*/
	Authorization string
	/*GroupData
	  Data to create a group.

	*/
	GroupData *models_core.GroupCreateBody

	timeout time.Duration
}

// WithAuthorization adds the authorization to the post categories groups params
func (o *PostCategoriesGroupsParams) WithAuthorization(Authorization string) *PostCategoriesGroupsParams {
	o.Authorization = Authorization
	return o
}

// WithGroupData adds the groupData to the post categories groups params
func (o *PostCategoriesGroupsParams) WithGroupData(GroupData *models_core.GroupCreateBody) *PostCategoriesGroupsParams {
	o.GroupData = GroupData
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostCategoriesGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.GroupData == nil {
		o.GroupData = new(models_core.GroupCreateBody)
	}

	if err := r.SetBodyParam(o.GroupData); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}