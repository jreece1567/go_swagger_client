package p_r_o_d_u_c_t_s

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

// NewGetProductsProductIDSyndicatedParams creates a new GetProductsProductIDSyndicatedParams object
// with the default values initialized.
func NewGetProductsProductIDSyndicatedParams() *GetProductsProductIDSyndicatedParams {
	var ()
	return &GetProductsProductIDSyndicatedParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProductsProductIDSyndicatedParamsWithTimeout creates a new GetProductsProductIDSyndicatedParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProductsProductIDSyndicatedParamsWithTimeout(timeout time.Duration) *GetProductsProductIDSyndicatedParams {
	var ()
	return &GetProductsProductIDSyndicatedParams{

		timeout: timeout,
	}
}

/*GetProductsProductIDSyndicatedParams contains all the parameters to send to the API endpoint
for the get products product ID syndicated operation typically these are written to a http.Request
*/
type GetProductsProductIDSyndicatedParams struct {

	/*Fields
	  Fields in response. Array that lists the fields requested.

	*/
	Fields []string
	/*ProductID
	  Product identifier. Request the syndicated product with product_id.

	*/
	ProductID string

	timeout time.Duration
}

// WithFields adds the fields to the get products product ID syndicated params
func (o *GetProductsProductIDSyndicatedParams) WithFields(Fields []string) *GetProductsProductIDSyndicatedParams {
	o.Fields = Fields
	return o
}

// WithProductID adds the productId to the get products product ID syndicated params
func (o *GetProductsProductIDSyndicatedParams) WithProductID(ProductID string) *GetProductsProductIDSyndicatedParams {
	o.ProductID = ProductID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetProductsProductIDSyndicatedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	valuesFields := o.Fields

	joinedFields := swag.JoinByFormat(valuesFields, "csv")
	// query array param fields
	if err := r.SetQueryParam("fields", joinedFields...); err != nil {
		return err
	}

	// path param product_id
	if err := r.SetPathParam("product_id", o.ProductID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
