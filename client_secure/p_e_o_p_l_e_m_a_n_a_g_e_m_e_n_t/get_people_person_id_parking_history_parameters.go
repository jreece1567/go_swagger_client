package p_e_o_p_l_e_m_a_n_a_g_e_m_e_n_t

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

// NewGetPeoplePersonIDParkingHistoryParams creates a new GetPeoplePersonIDParkingHistoryParams object
// with the default values initialized.
func NewGetPeoplePersonIDParkingHistoryParams() *GetPeoplePersonIDParkingHistoryParams {
	var (
		centreIdDefault string = string("london")
		pageDefault     int64  = int64(1)
		perPageDefault  int64  = int64(25)
	)
	return &GetPeoplePersonIDParkingHistoryParams{
		CentreID: &centreIdDefault,
		Page:     &pageDefault,
		PerPage:  &perPageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPeoplePersonIDParkingHistoryParamsWithTimeout creates a new GetPeoplePersonIDParkingHistoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPeoplePersonIDParkingHistoryParamsWithTimeout(timeout time.Duration) *GetPeoplePersonIDParkingHistoryParams {
	var (
		centreIdDefault string = string("london")
		pageDefault     int64  = int64(1)
		perPageDefault  int64  = int64(25)
	)
	return &GetPeoplePersonIDParkingHistoryParams{
		CentreID: &centreIdDefault,
		Page:     &pageDefault,
		PerPage:  &perPageDefault,

		timeout: timeout,
	}
}

/*GetPeoplePersonIDParkingHistoryParams contains all the parameters to send to the API endpoint
for the get people person ID parking history operation typically these are written to a http.Request
*/
type GetPeoplePersonIDParkingHistoryParams struct {

	/*Authorization
	  Staff Access token.

	*/
	Authorization string
	/*CentreID
	  Centre identifier. Return parking sessions of the given center.

	*/
	CentreID *string
	/*EndDate
	  Return parking history on or before a specified date.

	*/
	EndDate *strfmt.Date
	/*LicensePlate
	  Vehicle identifier. Return parking sessions with the given license_plate.

	*/
	LicensePlate *string
	/*Page
	  Results page number. Specific page of results to request when paginating.

	*/
	Page *int64
	/*PerPage
	  Results per page. Number of items per result set when paginating.

	*/
	PerPage *int64
	/*PersonID
	  Identifier of the account

	*/
	PersonID string
	/*StartDate
	  Return parking history on or after a specified date.

	*/
	StartDate *strfmt.Date

	timeout time.Duration
}

// WithAuthorization adds the authorization to the get people person ID parking history params
func (o *GetPeoplePersonIDParkingHistoryParams) WithAuthorization(Authorization string) *GetPeoplePersonIDParkingHistoryParams {
	o.Authorization = Authorization
	return o
}

// WithCentreID adds the centreId to the get people person ID parking history params
func (o *GetPeoplePersonIDParkingHistoryParams) WithCentreID(CentreID *string) *GetPeoplePersonIDParkingHistoryParams {
	o.CentreID = CentreID
	return o
}

// WithEndDate adds the endDate to the get people person ID parking history params
func (o *GetPeoplePersonIDParkingHistoryParams) WithEndDate(EndDate *strfmt.Date) *GetPeoplePersonIDParkingHistoryParams {
	o.EndDate = EndDate
	return o
}

// WithLicensePlate adds the licensePlate to the get people person ID parking history params
func (o *GetPeoplePersonIDParkingHistoryParams) WithLicensePlate(LicensePlate *string) *GetPeoplePersonIDParkingHistoryParams {
	o.LicensePlate = LicensePlate
	return o
}

// WithPage adds the page to the get people person ID parking history params
func (o *GetPeoplePersonIDParkingHistoryParams) WithPage(Page *int64) *GetPeoplePersonIDParkingHistoryParams {
	o.Page = Page
	return o
}

// WithPerPage adds the perPage to the get people person ID parking history params
func (o *GetPeoplePersonIDParkingHistoryParams) WithPerPage(PerPage *int64) *GetPeoplePersonIDParkingHistoryParams {
	o.PerPage = PerPage
	return o
}

// WithPersonID adds the personId to the get people person ID parking history params
func (o *GetPeoplePersonIDParkingHistoryParams) WithPersonID(PersonID string) *GetPeoplePersonIDParkingHistoryParams {
	o.PersonID = PersonID
	return o
}

// WithStartDate adds the startDate to the get people person ID parking history params
func (o *GetPeoplePersonIDParkingHistoryParams) WithStartDate(StartDate *strfmt.Date) *GetPeoplePersonIDParkingHistoryParams {
	o.StartDate = StartDate
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetPeoplePersonIDParkingHistoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

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

	if o.EndDate != nil {

		// query param end_date
		var qrEndDate strfmt.Date
		if o.EndDate != nil {
			qrEndDate = *o.EndDate
		}
		qEndDate := qrEndDate.String()
		if qEndDate != "" {
			if err := r.SetQueryParam("end_date", qEndDate); err != nil {
				return err
			}
		}

	}

	if o.LicensePlate != nil {

		// query param license_plate
		var qrLicensePlate string
		if o.LicensePlate != nil {
			qrLicensePlate = *o.LicensePlate
		}
		qLicensePlate := qrLicensePlate
		if qLicensePlate != "" {
			if err := r.SetQueryParam("license_plate", qLicensePlate); err != nil {
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

	// path param person_id
	if err := r.SetPathParam("person_id", o.PersonID); err != nil {
		return err
	}

	if o.StartDate != nil {

		// query param start_date
		var qrStartDate strfmt.Date
		if o.StartDate != nil {
			qrStartDate = *o.StartDate
		}
		qStartDate := qrStartDate.String()
		if qStartDate != "" {
			if err := r.SetQueryParam("start_date", qStartDate); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}