package c_e_n_t_r_e_p_a_r_k_i_n_g_r_a_t_e_s

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

// NewGetCentresParkingsParams creates a new GetCentresParkingsParams object
// with the default values initialized.
func NewGetCentresParkingsParams() *GetCentresParkingsParams {
	var (
		deletedDefault bool = bool(false)
	)
	return &GetCentresParkingsParams{
		Deleted: &deletedDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCentresParkingsParamsWithTimeout creates a new GetCentresParkingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCentresParkingsParamsWithTimeout(timeout time.Duration) *GetCentresParkingsParams {
	var (
		deletedDefault bool = bool(false)
	)
	return &GetCentresParkingsParams{
		Deleted: &deletedDefault,

		timeout: timeout,
	}
}

/*GetCentresParkingsParams contains all the parameters to send to the API endpoint
for the get centres parkings operation typically these are written to a http.Request
*/
type GetCentresParkingsParams struct {

	/*CentreID
	  Centre identifier. Request centre parking rates with centre_id. Lowercase code name for a specific centre. Must relate to store_id. Examples: ['sanfrancisco','sydney','london']

	*/
	CentreID *string
	/*Deleted
	  Deleted centre parking rates. Request to include (true) or not include (false) centre parking rates that are deleted.

	*/
	Deleted *bool
	/*UpdatedSince
	  Updated since. Request the centre parking rates updated since a specific date and time. ISO-8601 format.

	*/
	UpdatedSince *strfmt.DateTime

	timeout time.Duration
}

// WithCentreID adds the centreId to the get centres parkings params
func (o *GetCentresParkingsParams) WithCentreID(CentreID *string) *GetCentresParkingsParams {
	o.CentreID = CentreID
	return o
}

// WithDeleted adds the deleted to the get centres parkings params
func (o *GetCentresParkingsParams) WithDeleted(Deleted *bool) *GetCentresParkingsParams {
	o.Deleted = Deleted
	return o
}

// WithUpdatedSince adds the updatedSince to the get centres parkings params
func (o *GetCentresParkingsParams) WithUpdatedSince(UpdatedSince *strfmt.DateTime) *GetCentresParkingsParams {
	o.UpdatedSince = UpdatedSince
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetCentresParkingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.CentreID != nil {

		// query param centre_id
		var qrCentreID string
		if o.CentreID != nil {
			qrCentreID = *o.CentreID
		}
		qCentreID := qrCentreID
		if qCentreID != "" {
			if err := r.SetQueryParam("centre_id", qCentreID); err != nil {
				return err
			}
		}

	}

	if o.Deleted != nil {

		// query param deleted
		var qrDeleted bool
		if o.Deleted != nil {
			qrDeleted = *o.Deleted
		}
		qDeleted := swag.FormatBool(qrDeleted)
		if qDeleted != "" {
			if err := r.SetQueryParam("deleted", qDeleted); err != nil {
				return err
			}
		}

	}

	if o.UpdatedSince != nil {

		// query param updated_since
		var qrUpdatedSince strfmt.DateTime
		if o.UpdatedSince != nil {
			qrUpdatedSince = *o.UpdatedSince
		}
		qUpdatedSince := qrUpdatedSince.String()
		if qUpdatedSince != "" {
			if err := r.SetQueryParam("updated_since", qUpdatedSince); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}