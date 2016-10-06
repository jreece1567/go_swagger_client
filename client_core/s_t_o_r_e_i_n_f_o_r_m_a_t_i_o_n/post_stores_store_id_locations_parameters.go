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

	"restclient/models_core"
)

// NewPostStoresStoreIDLocationsParams creates a new PostStoresStoreIDLocationsParams object
// with the default values initialized.
func NewPostStoresStoreIDLocationsParams() *PostStoresStoreIDLocationsParams {
	var ()
	return &PostStoresStoreIDLocationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostStoresStoreIDLocationsParamsWithTimeout creates a new PostStoresStoreIDLocationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostStoresStoreIDLocationsParamsWithTimeout(timeout time.Duration) *PostStoresStoreIDLocationsParams {
	var ()
	return &PostStoresStoreIDLocationsParams{

		timeout: timeout,
	}
}

/*PostStoresStoreIDLocationsParams contains all the parameters to send to the API endpoint
for the post stores store ID locations operation typically these are written to a http.Request
*/
type PostStoresStoreIDLocationsParams struct {

	/*Authorization
	  Staff Access token.

	*/
	Authorization string
	/*LocationData
	  Data to create a store location.

	*/
	LocationData *models_core.StoreLocation
	/*StoreID
	  Store identifier. Request to create the store location with store_id.

	*/
	StoreID int64

	timeout time.Duration
}

// WithAuthorization adds the authorization to the post stores store ID locations params
func (o *PostStoresStoreIDLocationsParams) WithAuthorization(Authorization string) *PostStoresStoreIDLocationsParams {
	o.Authorization = Authorization
	return o
}

// WithLocationData adds the locationData to the post stores store ID locations params
func (o *PostStoresStoreIDLocationsParams) WithLocationData(LocationData *models_core.StoreLocation) *PostStoresStoreIDLocationsParams {
	o.LocationData = LocationData
	return o
}

// WithStoreID adds the storeId to the post stores store ID locations params
func (o *PostStoresStoreIDLocationsParams) WithStoreID(StoreID int64) *PostStoresStoreIDLocationsParams {
	o.StoreID = StoreID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostStoresStoreIDLocationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.LocationData == nil {
		o.LocationData = new(models_core.StoreLocation)
	}

	if err := r.SetBodyParam(o.LocationData); err != nil {
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
