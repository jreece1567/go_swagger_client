package p_a_y_m_e_n_t_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"restclient/models_secure"
)

// NewPatchPaymentsMerchantsMerchantIDParams creates a new PatchPaymentsMerchantsMerchantIDParams object
// with the default values initialized.
func NewPatchPaymentsMerchantsMerchantIDParams() *PatchPaymentsMerchantsMerchantIDParams {
	var ()
	return &PatchPaymentsMerchantsMerchantIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchPaymentsMerchantsMerchantIDParamsWithTimeout creates a new PatchPaymentsMerchantsMerchantIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchPaymentsMerchantsMerchantIDParamsWithTimeout(timeout time.Duration) *PatchPaymentsMerchantsMerchantIDParams {
	var ()
	return &PatchPaymentsMerchantsMerchantIDParams{

		timeout: timeout,
	}
}

/*PatchPaymentsMerchantsMerchantIDParams contains all the parameters to send to the API endpoint
for the patch payments merchants merchant ID operation typically these are written to a http.Request
*/
type PatchPaymentsMerchantsMerchantIDParams struct {

	/*Authorization
	  Access token

	*/
	Authorization string
	/*MerchantData
	  Data to update a merchant.

	*/
	MerchantData *models_secure.MerchantUpdateBody
	/*MerchantID
	  Merchant identifier.

	*/
	MerchantID int64

	timeout time.Duration
}

// WithAuthorization adds the authorization to the patch payments merchants merchant ID params
func (o *PatchPaymentsMerchantsMerchantIDParams) WithAuthorization(Authorization string) *PatchPaymentsMerchantsMerchantIDParams {
	o.Authorization = Authorization
	return o
}

// WithMerchantData adds the merchantData to the patch payments merchants merchant ID params
func (o *PatchPaymentsMerchantsMerchantIDParams) WithMerchantData(MerchantData *models_secure.MerchantUpdateBody) *PatchPaymentsMerchantsMerchantIDParams {
	o.MerchantData = MerchantData
	return o
}

// WithMerchantID adds the merchantId to the patch payments merchants merchant ID params
func (o *PatchPaymentsMerchantsMerchantIDParams) WithMerchantID(MerchantID int64) *PatchPaymentsMerchantsMerchantIDParams {
	o.MerchantID = MerchantID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PatchPaymentsMerchantsMerchantIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.MerchantData == nil {
		o.MerchantData = new(models_secure.MerchantUpdateBody)
	}

	if err := r.SetBodyParam(o.MerchantData); err != nil {
		return err
	}

	// path param merchant_id
	if err := r.SetPathParam("merchant_id", swag.FormatInt64(o.MerchantID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}