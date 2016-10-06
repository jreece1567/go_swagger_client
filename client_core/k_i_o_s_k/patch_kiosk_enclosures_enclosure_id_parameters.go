package k_i_o_s_k

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

// NewPatchKioskEnclosuresEnclosureIDParams creates a new PatchKioskEnclosuresEnclosureIDParams object
// with the default values initialized.
func NewPatchKioskEnclosuresEnclosureIDParams() *PatchKioskEnclosuresEnclosureIDParams {
	var ()
	return &PatchKioskEnclosuresEnclosureIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchKioskEnclosuresEnclosureIDParamsWithTimeout creates a new PatchKioskEnclosuresEnclosureIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchKioskEnclosuresEnclosureIDParamsWithTimeout(timeout time.Duration) *PatchKioskEnclosuresEnclosureIDParams {
	var ()
	return &PatchKioskEnclosuresEnclosureIDParams{

		timeout: timeout,
	}
}

/*PatchKioskEnclosuresEnclosureIDParams contains all the parameters to send to the API endpoint
for the patch kiosk enclosures enclosure ID operation typically these are written to a http.Request
*/
type PatchKioskEnclosuresEnclosureIDParams struct {

	/*Authorization
	  Staff Access token.

	*/
	Authorization string
	/*EnclosureData
	  Data to update an enclosure.

	*/
	EnclosureData *models_core.EnclosureUpdateBody
	/*EnclosureID
	  Enclosure identifier. Request to update the enclosure with enclosure_id.

	*/
	EnclosureID int64

	timeout time.Duration
}

// WithAuthorization adds the authorization to the patch kiosk enclosures enclosure ID params
func (o *PatchKioskEnclosuresEnclosureIDParams) WithAuthorization(Authorization string) *PatchKioskEnclosuresEnclosureIDParams {
	o.Authorization = Authorization
	return o
}

// WithEnclosureData adds the enclosureData to the patch kiosk enclosures enclosure ID params
func (o *PatchKioskEnclosuresEnclosureIDParams) WithEnclosureData(EnclosureData *models_core.EnclosureUpdateBody) *PatchKioskEnclosuresEnclosureIDParams {
	o.EnclosureData = EnclosureData
	return o
}

// WithEnclosureID adds the enclosureId to the patch kiosk enclosures enclosure ID params
func (o *PatchKioskEnclosuresEnclosureIDParams) WithEnclosureID(EnclosureID int64) *PatchKioskEnclosuresEnclosureIDParams {
	o.EnclosureID = EnclosureID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PatchKioskEnclosuresEnclosureIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.EnclosureData == nil {
		o.EnclosureData = new(models_core.EnclosureUpdateBody)
	}

	if err := r.SetBodyParam(o.EnclosureData); err != nil {
		return err
	}

	// path param enclosure_id
	if err := r.SetPathParam("enclosure_id", swag.FormatInt64(o.EnclosureID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
