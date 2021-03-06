package p_e_o_p_l_e_c_o_n_s_u_m_e_r

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

// NewGetAccountVehiclesParams creates a new GetAccountVehiclesParams object
// with the default values initialized.
func NewGetAccountVehiclesParams() *GetAccountVehiclesParams {
	var (
		includeDeletedDefault bool  = bool(false)
		pageDefault           int64 = int64(1)
		perPageDefault        int64 = int64(25)
	)
	return &GetAccountVehiclesParams{
		IncludeDeleted: &includeDeletedDefault,
		Page:           &pageDefault,
		PerPage:        &perPageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountVehiclesParamsWithTimeout creates a new GetAccountVehiclesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAccountVehiclesParamsWithTimeout(timeout time.Duration) *GetAccountVehiclesParams {
	var (
		includeDeletedDefault bool  = bool(false)
		pageDefault           int64 = int64(1)
		perPageDefault        int64 = int64(25)
	)
	return &GetAccountVehiclesParams{
		IncludeDeleted: &includeDeletedDefault,
		Page:           &pageDefault,
		PerPage:        &perPageDefault,

		timeout: timeout,
	}
}

/*GetAccountVehiclesParams contains all the parameters to send to the API endpoint
for the get account vehicles operation typically these are written to a http.Request
*/
type GetAccountVehiclesParams struct {

	/*Authorization
	  Access token.

	*/
	Authorization string
	/*Fields
	  Fields in response. Array that lists the fields requested. If empty or not supplied, all attributes will be returned.

	*/
	Fields []string
	/*IncludeDeleted
	  Value is true if deleted vehicles will be included in the response.

	*/
	IncludeDeleted *bool
	/*Page
	  Results page number. Specific page of results to request when paginating.

	*/
	Page *int64
	/*PerPage
	  Results per page. Number of items per result set when paginating.

	*/
	PerPage *int64
	/*UsedForParking
	  if Value is true return vehicles already used for parking, if false return vehicles not used for parking, if the param is not passed return both. Apply the filter and respect any other passed filter

	*/
	UsedForParking *bool

	timeout time.Duration
}

// WithAuthorization adds the authorization to the get account vehicles params
func (o *GetAccountVehiclesParams) WithAuthorization(Authorization string) *GetAccountVehiclesParams {
	o.Authorization = Authorization
	return o
}

// WithFields adds the fields to the get account vehicles params
func (o *GetAccountVehiclesParams) WithFields(Fields []string) *GetAccountVehiclesParams {
	o.Fields = Fields
	return o
}

// WithIncludeDeleted adds the includeDeleted to the get account vehicles params
func (o *GetAccountVehiclesParams) WithIncludeDeleted(IncludeDeleted *bool) *GetAccountVehiclesParams {
	o.IncludeDeleted = IncludeDeleted
	return o
}

// WithPage adds the page to the get account vehicles params
func (o *GetAccountVehiclesParams) WithPage(Page *int64) *GetAccountVehiclesParams {
	o.Page = Page
	return o
}

// WithPerPage adds the perPage to the get account vehicles params
func (o *GetAccountVehiclesParams) WithPerPage(PerPage *int64) *GetAccountVehiclesParams {
	o.PerPage = PerPage
	return o
}

// WithUsedForParking adds the usedForParking to the get account vehicles params
func (o *GetAccountVehiclesParams) WithUsedForParking(UsedForParking *bool) *GetAccountVehiclesParams {
	o.UsedForParking = UsedForParking
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountVehiclesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	valuesFields := o.Fields

	joinedFields := swag.JoinByFormat(valuesFields, "csv")
	// query array param fields
	if err := r.SetQueryParam("fields", joinedFields...); err != nil {
		return err
	}

	if o.IncludeDeleted != nil {

		// query param include_deleted
		var qrIncludeDeleted bool
		if o.IncludeDeleted != nil {
			qrIncludeDeleted = *o.IncludeDeleted
		}
		qIncludeDeleted := swag.FormatBool(qrIncludeDeleted)
		if qIncludeDeleted != "" {
			if err := r.SetQueryParam("include_deleted", qIncludeDeleted); err != nil {
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

	if o.UsedForParking != nil {

		// query param used_for_parking
		var qrUsedForParking bool
		if o.UsedForParking != nil {
			qrUsedForParking = *o.UsedForParking
		}
		qUsedForParking := swag.FormatBool(qrUsedForParking)
		if qUsedForParking != "" {
			if err := r.SetQueryParam("used_for_parking", qUsedForParking); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
