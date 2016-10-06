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
)

// NewDeleteKioskEnclosuresEnclosureIDParams creates a new DeleteKioskEnclosuresEnclosureIDParams object
// with the default values initialized.
func NewDeleteKioskEnclosuresEnclosureIDParams() *DeleteKioskEnclosuresEnclosureIDParams {
	var ()
	return &DeleteKioskEnclosuresEnclosureIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteKioskEnclosuresEnclosureIDParamsWithTimeout creates a new DeleteKioskEnclosuresEnclosureIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteKioskEnclosuresEnclosureIDParamsWithTimeout(timeout time.Duration) *DeleteKioskEnclosuresEnclosureIDParams {
	var ()
	return &DeleteKioskEnclosuresEnclosureIDParams{

		timeout: timeout,
	}
}

/*DeleteKioskEnclosuresEnclosureIDParams contains all the parameters to send to the API endpoint
for the delete kiosk enclosures enclosure ID operation typically these are written to a http.Request
*/
type DeleteKioskEnclosuresEnclosureIDParams struct {

	/*Authorization
	  Staff Access token.

	*/
	Authorization string
	/*EnclosureID
	  Enclosure identifier. Request the enclosure with enclosure_id.

	*/
	EnclosureID int64

	timeout time.Duration
}

// WithAuthorization adds the authorization to the delete kiosk enclosures enclosure ID params
func (o *DeleteKioskEnclosuresEnclosureIDParams) WithAuthorization(Authorization string) *DeleteKioskEnclosuresEnclosureIDParams {
	o.Authorization = Authorization
	return o
}

// WithEnclosureID adds the enclosureId to the delete kiosk enclosures enclosure ID params
func (o *DeleteKioskEnclosuresEnclosureIDParams) WithEnclosureID(EnclosureID int64) *DeleteKioskEnclosuresEnclosureIDParams {
	o.EnclosureID = EnclosureID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteKioskEnclosuresEnclosureIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
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
