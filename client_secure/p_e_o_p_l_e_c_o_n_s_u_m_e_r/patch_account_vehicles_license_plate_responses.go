package p_e_o_p_l_e_c_o_n_s_u_m_e_r

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"restclient/models_secure"
)

// PatchAccountVehiclesLicensePlateReader is a Reader for the PatchAccountVehiclesLicensePlate structure.
type PatchAccountVehiclesLicensePlateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PatchAccountVehiclesLicensePlateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPatchAccountVehiclesLicensePlateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPatchAccountVehiclesLicensePlateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPatchAccountVehiclesLicensePlateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchAccountVehiclesLicensePlateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPatchAccountVehiclesLicensePlateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 423:
		result := NewPatchAccountVehiclesLicensePlateLocked()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchAccountVehiclesLicensePlateNoContent creates a PatchAccountVehiclesLicensePlateNoContent with default headers values
func NewPatchAccountVehiclesLicensePlateNoContent() *PatchAccountVehiclesLicensePlateNoContent {
	return &PatchAccountVehiclesLicensePlateNoContent{}
}

/*PatchAccountVehiclesLicensePlateNoContent handles this case with default header values.

Successfully updated the vehicle.
*/
type PatchAccountVehiclesLicensePlateNoContent struct {
}

func (o *PatchAccountVehiclesLicensePlateNoContent) Error() string {
	return fmt.Sprintf("[PATCH /account/vehicles/{license_plate}][%d] patchAccountVehiclesLicensePlateNoContent ", 204)
}

func (o *PatchAccountVehiclesLicensePlateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchAccountVehiclesLicensePlateBadRequest creates a PatchAccountVehiclesLicensePlateBadRequest with default headers values
func NewPatchAccountVehiclesLicensePlateBadRequest() *PatchAccountVehiclesLicensePlateBadRequest {
	return &PatchAccountVehiclesLicensePlateBadRequest{}
}

/*PatchAccountVehiclesLicensePlateBadRequest handles this case with default header values.

Bad request
*/
type PatchAccountVehiclesLicensePlateBadRequest struct {
	Payload *models_secure.Person400Response
}

func (o *PatchAccountVehiclesLicensePlateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /account/vehicles/{license_plate}][%d] patchAccountVehiclesLicensePlateBadRequest  %+v", 400, o.Payload)
}

func (o *PatchAccountVehiclesLicensePlateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.Person400Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAccountVehiclesLicensePlateUnauthorized creates a PatchAccountVehiclesLicensePlateUnauthorized with default headers values
func NewPatchAccountVehiclesLicensePlateUnauthorized() *PatchAccountVehiclesLicensePlateUnauthorized {
	return &PatchAccountVehiclesLicensePlateUnauthorized{}
}

/*PatchAccountVehiclesLicensePlateUnauthorized handles this case with default header values.

Unauthorized request.
*/
type PatchAccountVehiclesLicensePlateUnauthorized struct {
	Payload *models_secure.UnauthorizedResponse
}

func (o *PatchAccountVehiclesLicensePlateUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /account/vehicles/{license_plate}][%d] patchAccountVehiclesLicensePlateUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchAccountVehiclesLicensePlateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.UnauthorizedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAccountVehiclesLicensePlateNotFound creates a PatchAccountVehiclesLicensePlateNotFound with default headers values
func NewPatchAccountVehiclesLicensePlateNotFound() *PatchAccountVehiclesLicensePlateNotFound {
	return &PatchAccountVehiclesLicensePlateNotFound{}
}

/*PatchAccountVehiclesLicensePlateNotFound handles this case with default header values.

Vehicle not found for a user.
*/
type PatchAccountVehiclesLicensePlateNotFound struct {
	Payload *models_secure.Http404Response
}

func (o *PatchAccountVehiclesLicensePlateNotFound) Error() string {
	return fmt.Sprintf("[PATCH /account/vehicles/{license_plate}][%d] patchAccountVehiclesLicensePlateNotFound  %+v", 404, o.Payload)
}

func (o *PatchAccountVehiclesLicensePlateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.Http404Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAccountVehiclesLicensePlateUnprocessableEntity creates a PatchAccountVehiclesLicensePlateUnprocessableEntity with default headers values
func NewPatchAccountVehiclesLicensePlateUnprocessableEntity() *PatchAccountVehiclesLicensePlateUnprocessableEntity {
	return &PatchAccountVehiclesLicensePlateUnprocessableEntity{}
}

/*PatchAccountVehiclesLicensePlateUnprocessableEntity handles this case with default header values.

Vehicle not updated. Returns an array of error messages explaining the problems with the provided attributes.
*/
type PatchAccountVehiclesLicensePlateUnprocessableEntity struct {
	Payload *models_secure.AccountVehicle422Response
}

func (o *PatchAccountVehiclesLicensePlateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /account/vehicles/{license_plate}][%d] patchAccountVehiclesLicensePlateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PatchAccountVehiclesLicensePlateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.AccountVehicle422Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAccountVehiclesLicensePlateLocked creates a PatchAccountVehiclesLicensePlateLocked with default headers values
func NewPatchAccountVehiclesLicensePlateLocked() *PatchAccountVehiclesLicensePlateLocked {
	return &PatchAccountVehiclesLicensePlateLocked{}
}

/*PatchAccountVehiclesLicensePlateLocked handles this case with default header values.

Resource not updated because it's locked.
*/
type PatchAccountVehiclesLicensePlateLocked struct {
}

func (o *PatchAccountVehiclesLicensePlateLocked) Error() string {
	return fmt.Sprintf("[PATCH /account/vehicles/{license_plate}][%d] patchAccountVehiclesLicensePlateLocked ", 423)
}

func (o *PatchAccountVehiclesLicensePlateLocked) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}