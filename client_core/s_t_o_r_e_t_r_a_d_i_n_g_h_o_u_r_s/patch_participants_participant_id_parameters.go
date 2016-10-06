package s_t_o_r_e_t_r_a_d_i_n_g_h_o_u_r_s

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

// NewPatchParticipantsParticipantIDParams creates a new PatchParticipantsParticipantIDParams object
// with the default values initialized.
func NewPatchParticipantsParticipantIDParams() *PatchParticipantsParticipantIDParams {
	var ()
	return &PatchParticipantsParticipantIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchParticipantsParticipantIDParamsWithTimeout creates a new PatchParticipantsParticipantIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchParticipantsParticipantIDParamsWithTimeout(timeout time.Duration) *PatchParticipantsParticipantIDParams {
	var ()
	return &PatchParticipantsParticipantIDParams{

		timeout: timeout,
	}
}

/*PatchParticipantsParticipantIDParams contains all the parameters to send to the API endpoint
for the patch participants participant ID operation typically these are written to a http.Request
*/
type PatchParticipantsParticipantIDParams struct {

	/*Authorization
	  Staff Access token.

	*/
	Authorization string
	/*ParticipantData
	  Data to update a participant.

	*/
	ParticipantData *models_core.ParticipantUpdateBody
	/*ParticipantID
	  Participant identifier. Request to update the participant with participant_id.

	*/
	ParticipantID int64

	timeout time.Duration
}

// WithAuthorization adds the authorization to the patch participants participant ID params
func (o *PatchParticipantsParticipantIDParams) WithAuthorization(Authorization string) *PatchParticipantsParticipantIDParams {
	o.Authorization = Authorization
	return o
}

// WithParticipantData adds the participantData to the patch participants participant ID params
func (o *PatchParticipantsParticipantIDParams) WithParticipantData(ParticipantData *models_core.ParticipantUpdateBody) *PatchParticipantsParticipantIDParams {
	o.ParticipantData = ParticipantData
	return o
}

// WithParticipantID adds the participantId to the patch participants participant ID params
func (o *PatchParticipantsParticipantIDParams) WithParticipantID(ParticipantID int64) *PatchParticipantsParticipantIDParams {
	o.ParticipantID = ParticipantID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PatchParticipantsParticipantIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.ParticipantData == nil {
		o.ParticipantData = new(models_core.ParticipantUpdateBody)
	}

	if err := r.SetBodyParam(o.ParticipantData); err != nil {
		return err
	}

	// path param participant_id
	if err := r.SetPathParam("participant_id", swag.FormatInt64(o.ParticipantID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
