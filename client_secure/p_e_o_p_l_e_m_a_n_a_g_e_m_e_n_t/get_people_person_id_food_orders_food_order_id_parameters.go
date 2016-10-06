package p_e_o_p_l_e_m_a_n_a_g_e_m_e_n_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPeoplePersonIDFoodOrdersFoodOrderIDParams creates a new GetPeoplePersonIDFoodOrdersFoodOrderIDParams object
// with the default values initialized.
func NewGetPeoplePersonIDFoodOrdersFoodOrderIDParams() *GetPeoplePersonIDFoodOrdersFoodOrderIDParams {
	var ()
	return &GetPeoplePersonIDFoodOrdersFoodOrderIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPeoplePersonIDFoodOrdersFoodOrderIDParamsWithTimeout creates a new GetPeoplePersonIDFoodOrdersFoodOrderIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPeoplePersonIDFoodOrdersFoodOrderIDParamsWithTimeout(timeout time.Duration) *GetPeoplePersonIDFoodOrdersFoodOrderIDParams {
	var ()
	return &GetPeoplePersonIDFoodOrdersFoodOrderIDParams{

		timeout: timeout,
	}
}

/*GetPeoplePersonIDFoodOrdersFoodOrderIDParams contains all the parameters to send to the API endpoint
for the get people person ID food orders food order ID operation typically these are written to a http.Request
*/
type GetPeoplePersonIDFoodOrdersFoodOrderIDParams struct {

	/*Authorization
	  Staff Access token.

	*/
	Authorization string
	/*FoodOrderID
	  Identifier of the order

	*/
	FoodOrderID string
	/*PersonID
	  Identifier of the account

	*/
	PersonID string

	timeout time.Duration
}

// WithAuthorization adds the authorization to the get people person ID food orders food order ID params
func (o *GetPeoplePersonIDFoodOrdersFoodOrderIDParams) WithAuthorization(Authorization string) *GetPeoplePersonIDFoodOrdersFoodOrderIDParams {
	o.Authorization = Authorization
	return o
}

// WithFoodOrderID adds the foodOrderId to the get people person ID food orders food order ID params
func (o *GetPeoplePersonIDFoodOrdersFoodOrderIDParams) WithFoodOrderID(FoodOrderID string) *GetPeoplePersonIDFoodOrdersFoodOrderIDParams {
	o.FoodOrderID = FoodOrderID
	return o
}

// WithPersonID adds the personId to the get people person ID food orders food order ID params
func (o *GetPeoplePersonIDFoodOrdersFoodOrderIDParams) WithPersonID(PersonID string) *GetPeoplePersonIDFoodOrdersFoodOrderIDParams {
	o.PersonID = PersonID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetPeoplePersonIDFoodOrdersFoodOrderIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param food_order_id
	if err := r.SetPathParam("food_order_id", o.FoodOrderID); err != nil {
		return err
	}

	// path param person_id
	if err := r.SetPathParam("person_id", o.PersonID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
