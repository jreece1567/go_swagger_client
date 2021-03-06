package s_t_o_r_e_i_n_f_o_r_m_a_t_i_o_n

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

// NewGetStoresStoreIDParams creates a new GetStoresStoreIDParams object
// with the default values initialized.
func NewGetStoresStoreIDParams() *GetStoresStoreIDParams {
	var ()
	return &GetStoresStoreIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStoresStoreIDParamsWithTimeout creates a new GetStoresStoreIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStoresStoreIDParamsWithTimeout(timeout time.Duration) *GetStoresStoreIDParams {
	var ()
	return &GetStoresStoreIDParams{

		timeout: timeout,
	}
}

/*GetStoresStoreIDParams contains all the parameters to send to the API endpoint
for the get stores store ID operation typically these are written to a http.Request
*/
type GetStoresStoreIDParams struct {

	/*StoreID
	  Store identifier. Request the store with store_id.

	*/
	StoreID int64

	timeout time.Duration
}

// WithStoreID adds the storeId to the get stores store ID params
func (o *GetStoresStoreIDParams) WithStoreID(StoreID int64) *GetStoresStoreIDParams {
	o.StoreID = StoreID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetStoresStoreIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param store_id
	if err := r.SetPathParam("store_id", swag.FormatInt64(o.StoreID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
