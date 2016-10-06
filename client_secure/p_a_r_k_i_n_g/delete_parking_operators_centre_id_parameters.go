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

// NewDeleteParkingOperatorsCentreIDParams creates a new DeleteParkingOperatorsCentreIDParams object
// with the default values initialized.
func NewDeleteParkingOperatorsCentreIDParams() *DeleteParkingOperatorsCentreIDParams {
	var ()
	return &DeleteParkingOperatorsCentreIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteParkingOperatorsCentreIDParamsWithTimeout creates a new DeleteParkingOperatorsCentreIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteParkingOperatorsCentreIDParamsWithTimeout(timeout time.Duration) *DeleteParkingOperatorsCentreIDParams {
	var ()
	return &DeleteParkingOperatorsCentreIDParams{

		timeout: timeout,
	}
}

/*DeleteParkingOperatorsCentreIDParams contains all the parameters to send to the API endpoint
for the delete parking operators centre ID operation typically these are written to a http.Request
*/
type DeleteParkingOperatorsCentreIDParams struct {

	/*CentreID
	  Centre identifier. Request to delete the Car Park Business Operator with centre_id.

	*/
	CentreID int64

	timeout time.Duration
}

// WithCentreID adds the centreId to the delete parking operators centre ID params
func (o *DeleteParkingOperatorsCentreIDParams) WithCentreID(CentreID int64) *DeleteParkingOperatorsCentreIDParams {
	o.CentreID = CentreID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteParkingOperatorsCentreIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param centre_id
	if err := r.SetPathParam("centre_id", swag.FormatInt64(o.CentreID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}