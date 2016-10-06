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

// NewGetStoresStoreIDTradingHoursParams creates a new GetStoresStoreIDTradingHoursParams object
// with the default values initialized.
func NewGetStoresStoreIDTradingHoursParams() *GetStoresStoreIDTradingHoursParams {
	var ()
	return &GetStoresStoreIDTradingHoursParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStoresStoreIDTradingHoursParamsWithTimeout creates a new GetStoresStoreIDTradingHoursParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStoresStoreIDTradingHoursParamsWithTimeout(timeout time.Duration) *GetStoresStoreIDTradingHoursParams {
	var ()
	return &GetStoresStoreIDTradingHoursParams{

		timeout: timeout,
	}
}

/*GetStoresStoreIDTradingHoursParams contains all the parameters to send to the API endpoint
for the get stores store ID trading hours operation typically these are written to a http.Request
*/
type GetStoresStoreIDTradingHoursParams struct {

	/*CentreID
	  Centre identifier. Request the store trading-hours with centre_id. Lowercase code name for a specific centre. Must relate to store_id.

	*/
	CentreID string
	/*StoreID
	  Store identifier. Request the store trading-hours with store_id.

	*/
	StoreID int64

	timeout time.Duration
}

// WithCentreID adds the centreId to the get stores store ID trading hours params
func (o *GetStoresStoreIDTradingHoursParams) WithCentreID(CentreID string) *GetStoresStoreIDTradingHoursParams {
	o.CentreID = CentreID
	return o
}

// WithStoreID adds the storeId to the get stores store ID trading hours params
func (o *GetStoresStoreIDTradingHoursParams) WithStoreID(StoreID int64) *GetStoresStoreIDTradingHoursParams {
	o.StoreID = StoreID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetStoresStoreIDTradingHoursParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// query param centre_id
	qrCentreID := o.CentreID
	qCentreID := qrCentreID
	if qCentreID != "" {
		if err := r.SetQueryParam("centre_id", qCentreID); err != nil {
			return err
		}
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
