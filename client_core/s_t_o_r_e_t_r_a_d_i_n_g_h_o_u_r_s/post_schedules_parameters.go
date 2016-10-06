package s_t_o_r_e_t_r_a_d_i_n_g_h_o_u_r_s

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

// NewPostSchedulesParams creates a new PostSchedulesParams object
// with the default values initialized.
func NewPostSchedulesParams() *PostSchedulesParams {
	var ()
	return &PostSchedulesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSchedulesParamsWithTimeout creates a new PostSchedulesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSchedulesParamsWithTimeout(timeout time.Duration) *PostSchedulesParams {
	var ()
	return &PostSchedulesParams{

		timeout: timeout,
	}
}

/*PostSchedulesParams contains all the parameters to send to the API endpoint
for the post schedules operation typically these are written to a http.Request
*/
type PostSchedulesParams struct {

	/*Authorization
	  Staff Access token.

	*/
	Authorization string
	/*ScheduleData
	  Data to create a schedule.

	*/
	ScheduleData *models_core.ScheduleCreateBody

	timeout time.Duration
}

// WithAuthorization adds the authorization to the post schedules params
func (o *PostSchedulesParams) WithAuthorization(Authorization string) *PostSchedulesParams {
	o.Authorization = Authorization
	return o
}

// WithScheduleData adds the scheduleData to the post schedules params
func (o *PostSchedulesParams) WithScheduleData(ScheduleData *models_core.ScheduleCreateBody) *PostSchedulesParams {
	o.ScheduleData = ScheduleData
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostSchedulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.ScheduleData == nil {
		o.ScheduleData = new(models_core.ScheduleCreateBody)
	}

	if err := r.SetBodyParam(o.ScheduleData); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}