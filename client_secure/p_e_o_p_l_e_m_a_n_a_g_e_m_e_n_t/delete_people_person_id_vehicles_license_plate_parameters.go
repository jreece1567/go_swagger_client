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

// NewDeletePeoplePersonIDVehiclesLicensePlateParams creates a new DeletePeoplePersonIDVehiclesLicensePlateParams object
// with the default values initialized.
func NewDeletePeoplePersonIDVehiclesLicensePlateParams() *DeletePeoplePersonIDVehiclesLicensePlateParams {
	var ()
	return &DeletePeoplePersonIDVehiclesLicensePlateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePeoplePersonIDVehiclesLicensePlateParamsWithTimeout creates a new DeletePeoplePersonIDVehiclesLicensePlateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeletePeoplePersonIDVehiclesLicensePlateParamsWithTimeout(timeout time.Duration) *DeletePeoplePersonIDVehiclesLicensePlateParams {
	var ()
	return &DeletePeoplePersonIDVehiclesLicensePlateParams{

		timeout: timeout,
	}
}

/*DeletePeoplePersonIDVehiclesLicensePlateParams contains all the parameters to send to the API endpoint
for the delete people person ID vehicles license plate operation typically these are written to a http.Request
*/
type DeletePeoplePersonIDVehiclesLicensePlateParams struct {

	/*Authorization
	  Staff access token.

	*/
	Authorization string
	/*LicensePlate
	  Vehicle identifier. Request to delete the vehicle with license_plate.

	*/
	LicensePlate string
	/*PersonID
	  Identifier of the account

	*/
	PersonID string

	timeout time.Duration
}

// WithAuthorization adds the authorization to the delete people person ID vehicles license plate params
func (o *DeletePeoplePersonIDVehiclesLicensePlateParams) WithAuthorization(Authorization string) *DeletePeoplePersonIDVehiclesLicensePlateParams {
	o.Authorization = Authorization
	return o
}

// WithLicensePlate adds the licensePlate to the delete people person ID vehicles license plate params
func (o *DeletePeoplePersonIDVehiclesLicensePlateParams) WithLicensePlate(LicensePlate string) *DeletePeoplePersonIDVehiclesLicensePlateParams {
	o.LicensePlate = LicensePlate
	return o
}

// WithPersonID adds the personId to the delete people person ID vehicles license plate params
func (o *DeletePeoplePersonIDVehiclesLicensePlateParams) WithPersonID(PersonID string) *DeletePeoplePersonIDVehiclesLicensePlateParams {
	o.PersonID = PersonID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePeoplePersonIDVehiclesLicensePlateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// path param license_plate
	if err := r.SetPathParam("license_plate", o.LicensePlate); err != nil {
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