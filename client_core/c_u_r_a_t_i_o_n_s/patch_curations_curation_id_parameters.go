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

	"restclient/models_core"
)

// NewPatchCurationsCurationIDParams creates a new PatchCurationsCurationIDParams object
// with the default values initialized.
func NewPatchCurationsCurationIDParams() *PatchCurationsCurationIDParams {
	var ()
	return &PatchCurationsCurationIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchCurationsCurationIDParamsWithTimeout creates a new PatchCurationsCurationIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchCurationsCurationIDParamsWithTimeout(timeout time.Duration) *PatchCurationsCurationIDParams {
	var ()
	return &PatchCurationsCurationIDParams{

		timeout: timeout,
	}
}

/*PatchCurationsCurationIDParams contains all the parameters to send to the API endpoint
for the patch curations curation ID operation typically these are written to a http.Request
*/
type PatchCurationsCurationIDParams struct {

	/*Authorization
	  Staff Access token.

	*/
	Authorization string
	/*CurationData
	  Data to update a curation.

	*/
	CurationData *models_core.CurationUpdateBody
	/*CurationID
	  Curation identitier. Request the curation with curation_id.

	*/
	CurationID int64

	timeout time.Duration
}

// WithAuthorization adds the authorization to the patch curations curation ID params
func (o *PatchCurationsCurationIDParams) WithAuthorization(Authorization string) *PatchCurationsCurationIDParams {
	o.Authorization = Authorization
	return o
}

// WithCurationData adds the curationData to the patch curations curation ID params
func (o *PatchCurationsCurationIDParams) WithCurationData(CurationData *models_core.CurationUpdateBody) *PatchCurationsCurationIDParams {
	o.CurationData = CurationData
	return o
}

// WithCurationID adds the curationId to the patch curations curation ID params
func (o *PatchCurationsCurationIDParams) WithCurationID(CurationID int64) *PatchCurationsCurationIDParams {
	o.CurationID = CurationID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PatchCurationsCurationIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.CurationData == nil {
		o.CurationData = new(models_core.CurationUpdateBody)
	}

	if err := r.SetBodyParam(o.CurationData); err != nil {
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