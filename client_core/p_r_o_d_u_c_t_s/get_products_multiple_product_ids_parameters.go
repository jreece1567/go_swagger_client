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

// NewGetProductsMultipleProductIdsParams creates a new GetProductsMultipleProductIdsParams object
// with the default values initialized.
func NewGetProductsMultipleProductIdsParams() *GetProductsMultipleProductIdsParams {
	var ()
	return &GetProductsMultipleProductIdsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProductsMultipleProductIdsParamsWithTimeout creates a new GetProductsMultipleProductIdsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProductsMultipleProductIdsParamsWithTimeout(timeout time.Duration) *GetProductsMultipleProductIdsParams {
	var ()
	return &GetProductsMultipleProductIdsParams{

		timeout: timeout,
	}
}

/*GetProductsMultipleProductIdsParams contains all the parameters to send to the API endpoint
for the get products multiple product ids operation typically these are written to a http.Request
*/
type GetProductsMultipleProductIdsParams struct {

	/*Fields
	  Fields in response. Array that lists the fields requested.

	*/
	Fields []string
	/*ProductIds
	  Comma delimited list of product IDs to fetch

	*/
	ProductIds []string

	timeout time.Duration
}

// WithFields adds the fields to the get products multiple product ids params
func (o *GetProductsMultipleProductIdsParams) WithFields(Fields []string) *GetProductsMultipleProductIdsParams {
	o.Fields = Fields
	return o
}

// WithProductIds adds the productIds to the get products multiple product ids params
func (o *GetProductsMultipleProductIdsParams) WithProductIds(ProductIds []string) *GetProductsMultipleProductIdsParams {
	o.ProductIds = ProductIds
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetProductsMultipleProductIdsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	valuesFields := o.Fields

	joinedFields := swag.JoinByFormat(valuesFields, "csv")
	// query array param fields
	if err := r.SetQueryParam("fields", joinedFields...); err != nil {
		return err
	}

	//valuesProductIds := o.ProductIds

	//joinedProductIds := swag.JoinByFormat(valuesProductIds, "csv")

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
