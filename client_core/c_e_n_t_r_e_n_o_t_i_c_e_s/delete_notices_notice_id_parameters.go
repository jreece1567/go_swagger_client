package c_e_n_t_r_e_n_o_t_i_c_e_s

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

// NewDeleteNoticesNoticeIDParams creates a new DeleteNoticesNoticeIDParams object
// with the default values initialized.
func NewDeleteNoticesNoticeIDParams() *DeleteNoticesNoticeIDParams {
	var ()
	return &DeleteNoticesNoticeIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNoticesNoticeIDParamsWithTimeout creates a new DeleteNoticesNoticeIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteNoticesNoticeIDParamsWithTimeout(timeout time.Duration) *DeleteNoticesNoticeIDParams {
	var ()
	return &DeleteNoticesNoticeIDParams{

		timeout: timeout,
	}
}

/*DeleteNoticesNoticeIDParams contains all the parameters to send to the API endpoint
for the delete notices notice ID operation typically these are written to a http.Request
*/
type DeleteNoticesNoticeIDParams struct {

	/*Authorization
	  Staff Access token.

	*/
	Authorization string
	/*NoticeID
	  Notice identifier. Request to delete the notice with notice_id.

	*/
	NoticeID int64

	timeout time.Duration
}

// WithAuthorization adds the authorization to the delete notices notice ID params
func (o *DeleteNoticesNoticeIDParams) WithAuthorization(Authorization string) *DeleteNoticesNoticeIDParams {
	o.Authorization = Authorization
	return o
}

// WithNoticeID adds the noticeId to the delete notices notice ID params
func (o *DeleteNoticesNoticeIDParams) WithNoticeID(NoticeID int64) *DeleteNoticesNoticeIDParams {
	o.NoticeID = NoticeID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNoticesNoticeIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param notice_id
	if err := r.SetPathParam("notice_id", swag.FormatInt64(o.NoticeID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
