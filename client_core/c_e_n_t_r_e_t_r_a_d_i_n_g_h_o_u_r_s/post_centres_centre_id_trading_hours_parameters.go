package c_e_n_t_r_e_t_r_a_d_i_n_g_h_o_u_r_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"restclient/models_core"
)

// NewPostCentresCentreIDTradingHoursParams creates a new PostCentresCentreIDTradingHoursParams object
// with the default values initialized.
func NewPostCentresCentreIDTradingHoursParams() *PostCentresCentreIDTradingHoursParams {
	var ()
	return &PostCentresCentreIDTradingHoursParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCentresCentreIDTradingHoursParamsWithTimeout creates a new PostCentresCentreIDTradingHoursParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCentresCentreIDTradingHoursParamsWithTimeout(timeout time.Duration) *PostCentresCentreIDTradingHoursParams {
	var ()
	return &PostCentresCentreIDTradingHoursParams{

		timeout: timeout,
	}
}

/*PostCentresCentreIDTradingHoursParams contains all the parameters to send to the API endpoint
for the post centres centre ID trading hours operation typically these are written to a http.Request
*/
type PostCentresCentreIDTradingHoursParams struct {

	/*Authorization
	  Staff Access token.

	*/
	Authorization string
	/*CentreID
	  Centre identifier. Request the centre trading-hours with centre_id. Lowercase code name for a specific centre.

	*/
	CentreID string
	/*CentreTradingHourData
	  Data to create centre trading-hours.

	*/
	CentreTradingHourData *models_core.CreateCentreHoursBody

	timeout time.Duration
}

// WithAuthorization adds the authorization to the post centres centre ID trading hours params
func (o *PostCentresCentreIDTradingHoursParams) WithAuthorization(Authorization string) *PostCentresCentreIDTradingHoursParams {
	o.Authorization = Authorization
	return o
}

// WithCentreID adds the centreId to the post centres centre ID trading hours params
func (o *PostCentresCentreIDTradingHoursParams) WithCentreID(CentreID string) *PostCentresCentreIDTradingHoursParams {
	o.CentreID = CentreID
	return o
}

// WithCentreTradingHourData adds the centreTradingHourData to the post centres centre ID trading hours params
func (o *PostCentresCentreIDTradingHoursParams) WithCentreTradingHourData(CentreTradingHourData *models_core.CreateCentreHoursBody) *PostCentresCentreIDTradingHoursParams {
	o.CentreTradingHourData = CentreTradingHourData
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostCentresCentreIDTradingHoursParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param centre_id
	if err := r.SetPathParam("centre_id", o.CentreID); err != nil {
		return err
	}

	if o.CentreTradingHourData == nil {
		o.CentreTradingHourData = new(models_core.CreateCentreHoursBody)
	}

	if err := r.SetBodyParam(o.CentreTradingHourData); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
