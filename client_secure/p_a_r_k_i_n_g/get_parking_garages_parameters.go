package p_a_r_k_i_n_g

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

// NewGetParkingGaragesParams creates a new GetParkingGaragesParams object
// with the default values initialized.
func NewGetParkingGaragesParams() *GetParkingGaragesParams {
	var (
		centreIdDefault string = string("sanfrancisco")
		pageDefault     int64  = int64(1)
		perPageDefault  int64  = int64(25)
	)
	return &GetParkingGaragesParams{
		CentreID: &centreIdDefault,
		Page:     &pageDefault,
		PerPage:  &perPageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetParkingGaragesParamsWithTimeout creates a new GetParkingGaragesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetParkingGaragesParamsWithTimeout(timeout time.Duration) *GetParkingGaragesParams {
	var (
		centreIdDefault string = string("sanfrancisco")
		pageDefault     int64  = int64(1)
		perPageDefault  int64  = int64(25)
	)
	return &GetParkingGaragesParams{
		CentreID: &centreIdDefault,
		Page:     &pageDefault,
		PerPage:  &perPageDefault,

		timeout: timeout,
	}
}

/*GetParkingGaragesParams contains all the parameters to send to the API endpoint
for the get parking garages operation typically these are written to a http.Request
*/
type GetParkingGaragesParams struct {

	/*CentreID
	  Centre identifier. Request the parking garage details with centre_id. Lowercase code name for a specific centre.

	*/
	CentreID *string
	/*Fields
	  Fields in response. Array that lists the fields requested. If empty or not supplied, all attributes will be returned.

	*/
	Fields []string
	/*Page
	  Results page number. Specific page of results to request when paginating.

	*/
	Page *int64
	/*PerPage
	  Results per page. Number of items per result set when paginating.

	*/
	PerPage *int64

	timeout time.Duration
}

// WithCentreID adds the centreId to the get parking garages params
func (o *GetParkingGaragesParams) WithCentreID(CentreID *string) *GetParkingGaragesParams {
	o.CentreID = CentreID
	return o
}

// WithFields adds the fields to the get parking garages params
func (o *GetParkingGaragesParams) WithFields(Fields []string) *GetParkingGaragesParams {
	o.Fields = Fields
	return o
}

// WithPage adds the page to the get parking garages params
func (o *GetParkingGaragesParams) WithPage(Page *int64) *GetParkingGaragesParams {
	o.Page = Page
	return o
}

// WithPerPage adds the perPage to the get parking garages params
func (o *GetParkingGaragesParams) WithPerPage(PerPage *int64) *GetParkingGaragesParams {
	o.PerPage = PerPage
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetParkingGaragesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	valuesFields := o.Fields

	joinedFields := swag.JoinByFormat(valuesFields, "csv")
	// query array param fields
	if err := r.SetQueryParam("fields", joinedFields...); err != nil {
		return err
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
