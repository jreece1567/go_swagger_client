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
)

// NewGetParticipantsParams creates a new GetParticipantsParams object
// with the default values initialized.
func NewGetParticipantsParams() *GetParticipantsParams {
	var (
		pageDefault    int64 = int64(1)
		perPageDefault int64 = int64(10)
	)
	return &GetParticipantsParams{
		Page:    &pageDefault,
		PerPage: &perPageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetParticipantsParamsWithTimeout creates a new GetParticipantsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetParticipantsParamsWithTimeout(timeout time.Duration) *GetParticipantsParams {
	var (
		pageDefault    int64 = int64(1)
		perPageDefault int64 = int64(10)
	)
	return &GetParticipantsParams{
		Page:    &pageDefault,
		PerPage: &perPageDefault,

		timeout: timeout,
	}
}

/*GetParticipantsParams contains all the parameters to send to the API endpoint
for the get participants operation typically these are written to a http.Request
*/
type GetParticipantsParams struct {

	/*Fields
	  Fields in response. Array that lists the fields requested. If empty or not supplied, all attributes will be returned.

	*/
	Fields []string
	/*Kind
	  Type of participant.

	*/
	Kind *string
	/*KindID
	  Key of participant.

	*/
	KindID *string
	/*Page
	  Results page number. Specific page of results to request when paginating.

	*/
	Page *int64
	/*ParticipantIds
	  List of Participant identifiers. Request participants for each schedule_id.

	*/
	ParticipantIds []int64
	/*PerPage
	  Results per page. Number of items per result set when paginating.

	*/
	PerPage *int64

	timeout time.Duration
}

// WithFields adds the fields to the get participants params
func (o *GetParticipantsParams) WithFields(Fields []string) *GetParticipantsParams {
	o.Fields = Fields
	return o
}

// WithKind adds the kind to the get participants params
func (o *GetParticipantsParams) WithKind(Kind *string) *GetParticipantsParams {
	o.Kind = Kind
	return o
}

// WithKindID adds the kindId to the get participants params
func (o *GetParticipantsParams) WithKindID(KindID *string) *GetParticipantsParams {
	o.KindID = KindID
	return o
}

// WithPage adds the page to the get participants params
func (o *GetParticipantsParams) WithPage(Page *int64) *GetParticipantsParams {
	o.Page = Page
	return o
}

// WithParticipantIds adds the participantIds to the get participants params
func (o *GetParticipantsParams) WithParticipantIds(ParticipantIds []int64) *GetParticipantsParams {
	o.ParticipantIds = ParticipantIds
	return o
}

// WithPerPage adds the perPage to the get participants params
func (o *GetParticipantsParams) WithPerPage(PerPage *int64) *GetParticipantsParams {
	o.PerPage = PerPage
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetParticipantsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	valuesFields := o.Fields

	joinedFields := swag.JoinByFormat(valuesFields, "csv")
	// query array param fields
	if err := r.SetQueryParam("fields", joinedFields...); err != nil {
		return err
	}

	if o.Kind != nil {

		// query param kind
		var qrKind string
		if o.Kind != nil {
			qrKind = *o.Kind
		}
		qKind := qrKind
		if qKind != "" {
			if err := r.SetQueryParam("kind", qKind); err != nil {
				return err
			}
		}

	}

	if o.KindID != nil {

		// query param kind_id
		var qrKindID string
		if o.KindID != nil {
			qrKindID = *o.KindID
		}
		qKindID := qrKindID
		if qKindID != "" {
			if err := r.SetQueryParam("kind_id", qKindID); err != nil {
				return err
			}
		}

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

	var valuesParticipantIds []string
	for _, v := range o.ParticipantIds {
		valuesParticipantIds = append(valuesParticipantIds, swag.FormatInt64(v))
	}

	joinedParticipantIds := swag.JoinByFormat(valuesParticipantIds, "csv")
	// query array param participant_ids
	if err := r.SetQueryParam("participant_ids", joinedParticipantIds...); err != nil {
		return err
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
