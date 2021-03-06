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

// NewGetParkingSessionsParams creates a new GetParkingSessionsParams object
// with the default values initialized.
func NewGetParkingSessionsParams() *GetParkingSessionsParams {
	var (
		centreIdDefault string = string("sanfrancisco")
		pageDefault     int64  = int64(1)
		perPageDefault  int64  = int64(25)
	)
	return &GetParkingSessionsParams{
		CentreID: &centreIdDefault,
		Page:     &pageDefault,
		PerPage:  &perPageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetParkingSessionsParamsWithTimeout creates a new GetParkingSessionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetParkingSessionsParamsWithTimeout(timeout time.Duration) *GetParkingSessionsParams {
	var (
		centreIdDefault string = string("sanfrancisco")
		pageDefault     int64  = int64(1)
		perPageDefault  int64  = int64(25)
	)
	return &GetParkingSessionsParams{
		CentreID: &centreIdDefault,
		Page:     &pageDefault,
		PerPage:  &perPageDefault,

		timeout: timeout,
	}
}

/*GetParkingSessionsParams contains all the parameters to send to the API endpoint
for the get parking sessions operation typically these are written to a http.Request
*/
type GetParkingSessionsParams struct {

	/*AccessDevice
	  Access device identifier. Return parking sessions with the given access_device.

	*/
	AccessDevice *string
	/*CentreID
	  Centre identifier. Return parking sessions with the given centre_id. Lowercase code name for a specific centre.

	*/
	CentreID *string
	/*EndDate
	  Return parking sessions on or before a specified date.

	*/
	EndDate *strfmt.Date
	/*Fields
	  Fields in response. Array that lists the fields requested. If empty or not supplied, all attributes will be returned.

	*/
	Fields []string
	/*Page
	  Results page number. Specific page of results to request when paginating.

	*/
	Page *int64
	/*PaymentStatus
	  status of payment

	*/
	PaymentStatus *string
	/*PerPage
	  Results per page. Number of items per result set when paginating.

	*/
	PerPage *int64
	/*PersonID
	  Person identifier. Request all parking sessions for a user with the given person_id.

	*/
	PersonID string
	/*StartDate
	  Return parking sessions on or after a specified date.

	*/
	StartDate *strfmt.Date

	timeout time.Duration
}

// WithAccessDevice adds the accessDevice to the get parking sessions params
func (o *GetParkingSessionsParams) WithAccessDevice(AccessDevice *string) *GetParkingSessionsParams {
	o.AccessDevice = AccessDevice
	return o
}

// WithCentreID adds the centreId to the get parking sessions params
func (o *GetParkingSessionsParams) WithCentreID(CentreID *string) *GetParkingSessionsParams {
	o.CentreID = CentreID
	return o
}

// WithEndDate adds the endDate to the get parking sessions params
func (o *GetParkingSessionsParams) WithEndDate(EndDate *strfmt.Date) *GetParkingSessionsParams {
	o.EndDate = EndDate
	return o
}

// WithFields adds the fields to the get parking sessions params
func (o *GetParkingSessionsParams) WithFields(Fields []string) *GetParkingSessionsParams {
	o.Fields = Fields
	return o
}

// WithPage adds the page to the get parking sessions params
func (o *GetParkingSessionsParams) WithPage(Page *int64) *GetParkingSessionsParams {
	o.Page = Page
	return o
}

// WithPaymentStatus adds the paymentStatus to the get parking sessions params
func (o *GetParkingSessionsParams) WithPaymentStatus(PaymentStatus *string) *GetParkingSessionsParams {
	o.PaymentStatus = PaymentStatus
	return o
}

// WithPerPage adds the perPage to the get parking sessions params
func (o *GetParkingSessionsParams) WithPerPage(PerPage *int64) *GetParkingSessionsParams {
	o.PerPage = PerPage
	return o
}

// WithPersonID adds the personId to the get parking sessions params
func (o *GetParkingSessionsParams) WithPersonID(PersonID string) *GetParkingSessionsParams {
	o.PersonID = PersonID
	return o
}

// WithStartDate adds the startDate to the get parking sessions params
func (o *GetParkingSessionsParams) WithStartDate(StartDate *strfmt.Date) *GetParkingSessionsParams {
	o.StartDate = StartDate
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetParkingSessionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.AccessDevice != nil {

		// query param access_device
		var qrAccessDevice string
		if o.AccessDevice != nil {
			qrAccessDevice = *o.AccessDevice
		}
		qAccessDevice := qrAccessDevice
		if qAccessDevice != "" {
			if err := r.SetQueryParam("access_device", qAccessDevice); err != nil {
				return err
			}
		}

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

	if o.PaymentStatus != nil {

		// query param payment_status
		var qrPaymentStatus string
		if o.PaymentStatus != nil {
			qrPaymentStatus = *o.PaymentStatus
		}
		qPaymentStatus := qrPaymentStatus
		if qPaymentStatus != "" {
			if err := r.SetQueryParam("payment_status", qPaymentStatus); err != nil {
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

	// query param person_id
	qrPersonID := o.PersonID
	qPersonID := qrPersonID
	if qPersonID != "" {
		if err := r.SetQueryParam("person_id", qPersonID); err != nil {
			return err
		}
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
