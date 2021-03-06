package c_u_r_a_t_i_o_n_s

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

// NewDeleteCurationsCurationIDParams creates a new DeleteCurationsCurationIDParams object
// with the default values initialized.
func NewDeleteCurationsCurationIDParams() *DeleteCurationsCurationIDParams {
	var ()
	return &DeleteCurationsCurationIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCurationsCurationIDParamsWithTimeout creates a new DeleteCurationsCurationIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCurationsCurationIDParamsWithTimeout(timeout time.Duration) *DeleteCurationsCurationIDParams {
	var ()
	return &DeleteCurationsCurationIDParams{

		timeout: timeout,
	}
}

/*DeleteCurationsCurationIDParams contains all the parameters to send to the API endpoint
for the delete curations curation ID operation typically these are written to a http.Request
*/
type DeleteCurationsCurationIDParams struct {

	/*Authorization
	  Staff Access token.

	*/
	Authorization string
	/*CurationID
	  Curation identitier. Request the curation with curation_id.

	*/
	CurationID int64

	timeout time.Duration
}

// WithAuthorization adds the authorization to the delete curations curation ID params
func (o *DeleteCurationsCurationIDParams) WithAuthorization(Authorization string) *DeleteCurationsCurationIDParams {
	o.Authorization = Authorization
	return o
}

// WithCurationID adds the curationId to the delete curations curation ID params
func (o *DeleteCurationsCurationIDParams) WithCurationID(CurationID int64) *DeleteCurationsCurationIDParams {
	o.CurationID = CurationID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCurationsCurationIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param curation_id
	if err := r.SetPathParam("curation_id", swag.FormatInt64(o.CurationID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
