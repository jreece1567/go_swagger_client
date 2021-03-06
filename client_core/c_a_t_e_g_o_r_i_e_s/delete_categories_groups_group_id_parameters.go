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

// NewDeleteCategoriesGroupsGroupIDParams creates a new DeleteCategoriesGroupsGroupIDParams object
// with the default values initialized.
func NewDeleteCategoriesGroupsGroupIDParams() *DeleteCategoriesGroupsGroupIDParams {
	var ()
	return &DeleteCategoriesGroupsGroupIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCategoriesGroupsGroupIDParamsWithTimeout creates a new DeleteCategoriesGroupsGroupIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCategoriesGroupsGroupIDParamsWithTimeout(timeout time.Duration) *DeleteCategoriesGroupsGroupIDParams {
	var ()
	return &DeleteCategoriesGroupsGroupIDParams{

		timeout: timeout,
	}
}

/*DeleteCategoriesGroupsGroupIDParams contains all the parameters to send to the API endpoint
for the delete categories groups group ID operation typically these are written to a http.Request
*/
type DeleteCategoriesGroupsGroupIDParams struct {

	/*Authorization
	  Staff Access token.

	*/
	Authorization string
	/*GroupID
	  Group identifier. Request to delete the group with group_id.

	*/
	GroupID int64

	timeout time.Duration
}

// WithAuthorization adds the authorization to the delete categories groups group ID params
func (o *DeleteCategoriesGroupsGroupIDParams) WithAuthorization(Authorization string) *DeleteCategoriesGroupsGroupIDParams {
	o.Authorization = Authorization
	return o
}

// WithGroupID adds the groupId to the delete categories groups group ID params
func (o *DeleteCategoriesGroupsGroupIDParams) WithGroupID(GroupID int64) *DeleteCategoriesGroupsGroupIDParams {
	o.GroupID = GroupID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCategoriesGroupsGroupIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param group_id
	if err := r.SetPathParam("group_id", swag.FormatInt64(o.GroupID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
