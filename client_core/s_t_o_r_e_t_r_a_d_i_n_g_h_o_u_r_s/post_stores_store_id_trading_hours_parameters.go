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

// NewPostStoresStoreIDTradingHoursParams creates a new PostStoresStoreIDTradingHoursParams object
// with the default values initialized.
func NewPostStoresStoreIDTradingHoursParams() *PostStoresStoreIDTradingHoursParams {
	var ()
	return &PostStoresStoreIDTradingHoursParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostStoresStoreIDTradingHoursParamsWithTimeout creates a new PostStoresStoreIDTradingHoursParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostStoresStoreIDTradingHoursParamsWithTimeout(timeout time.Duration) *PostStoresStoreIDTradingHoursParams {
	var ()
	return &PostStoresStoreIDTradingHoursParams{

		timeout: timeout,
	}
}

/*PostStoresStoreIDTradingHoursParams contains all the parameters to send to the API endpoint
for the post stores store ID trading hours operation typically these are written to a http.Request
*/
type PostStoresStoreIDTradingHoursParams struct {

	/*Authorization
	  Staff Access token.

	*/
	Authorization string
	/*CentreID
	  Centre identifier. Request the store trading-hours with centre_id. Lowercase code name for a specific centre.

	*/
	CentreID string
	/*CentreTradingHourData
	  Data to create store trading-hours.

	*/
	CentreTradingHourData *models_core.CreateStoreHoursBody
	/*StoreID
	  Store identifier. Request the store trading-hours with store_id.

	*/
	StoreID int64

	timeout time.Duration
}

// WithAuthorization adds the authorization to the post stores store ID trading hours params
func (o *PostStoresStoreIDTradingHoursParams) WithAuthorization(Authorization string) *PostStoresStoreIDTradingHoursParams {
	o.Authorization = Authorization
	return o
}

// WithCentreID adds the centreId to the post stores store ID trading hours params
func (o *PostStoresStoreIDTradingHoursParams) WithCentreID(CentreID string) *PostStoresStoreIDTradingHoursParams {
	o.CentreID = CentreID
	return o
}

// WithCentreTradingHourData adds the centreTradingHourData to the post stores store ID trading hours params
func (o *PostStoresStoreIDTradingHoursParams) WithCentreTradingHourData(CentreTradingHourData *models_core.CreateStoreHoursBody) *PostStoresStoreIDTradingHoursParams {
	o.CentreTradingHourData = CentreTradingHourData
	return o
}

// WithStoreID adds the storeId to the post stores store ID trading hours params
func (o *PostStoresStoreIDTradingHoursParams) WithStoreID(StoreID int64) *PostStoresStoreIDTradingHoursParams {
	o.StoreID = StoreID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostStoresStoreIDTradingHoursParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// query param centre_id
	qrCentreID := o.CentreID
	qCentreID := qrCentreID
	if qCentreID != "" {
		if err := r.SetQueryParam("centre_id", qCentreID); err != nil {
			return err
		}
	}

	if o.CentreTradingHourData == nil {
		o.CentreTradingHourData = new(models_core.CreateStoreHoursBody)
	}

	if err := r.SetBodyParam(o.CentreTradingHourData); err != nil {
		return err
	}

	// path param store_id
	if err := r.SetPathParam("store_id", swag.FormatInt64(o.StoreID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
