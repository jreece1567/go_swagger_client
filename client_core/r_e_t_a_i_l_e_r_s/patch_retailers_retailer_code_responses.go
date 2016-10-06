package r_e_t_a_i_l_e_r_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"restclient/models_core"
)

// PatchRetailersRetailerCodeReader is a Reader for the PatchRetailersRetailerCode structure.
type PatchRetailersRetailerCodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PatchRetailersRetailerCodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPatchRetailersRetailerCodeNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPatchRetailersRetailerCodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchRetailersRetailerCodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPatchRetailersRetailerCodeUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchRetailersRetailerCodeNoContent creates a PatchRetailersRetailerCodeNoContent with default headers values
func NewPatchRetailersRetailerCodeNoContent() *PatchRetailersRetailerCodeNoContent {
	return &PatchRetailersRetailerCodeNoContent{}
}

/*PatchRetailersRetailerCodeNoContent handles this case with default header values.

Successfully updated the retailer.
*/
type PatchRetailersRetailerCodeNoContent struct {
}

func (o *PatchRetailersRetailerCodeNoContent) Error() string {
	return fmt.Sprintf("[PATCH /retailers/{retailer_code}][%d] patchRetailersRetailerCodeNoContent ", 204)
}

func (o *PatchRetailersRetailerCodeNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchRetailersRetailerCodeUnauthorized creates a PatchRetailersRetailerCodeUnauthorized with default headers values
func NewPatchRetailersRetailerCodeUnauthorized() *PatchRetailersRetailerCodeUnauthorized {
	return &PatchRetailersRetailerCodeUnauthorized{}
}

/*PatchRetailersRetailerCodeUnauthorized handles this case with default header values.

Unauthorized request.
*/
type PatchRetailersRetailerCodeUnauthorized struct {
	Payload *models_core.Http401Response
}

func (o *PatchRetailersRetailerCodeUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /retailers/{retailer_code}][%d] patchRetailersRetailerCodeUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchRetailersRetailerCodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.Http401Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRetailersRetailerCodeNotFound creates a PatchRetailersRetailerCodeNotFound with default headers values
func NewPatchRetailersRetailerCodeNotFound() *PatchRetailersRetailerCodeNotFound {
	return &PatchRetailersRetailerCodeNotFound{}
}

/*PatchRetailersRetailerCodeNotFound handles this case with default header values.

Retailer not found.
*/
type PatchRetailersRetailerCodeNotFound struct {
	Payload *models_core.Http404Response
}

func (o *PatchRetailersRetailerCodeNotFound) Error() string {
	return fmt.Sprintf("[PATCH /retailers/{retailer_code}][%d] patchRetailersRetailerCodeNotFound  %+v", 404, o.Payload)
}

func (o *PatchRetailersRetailerCodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.Http404Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRetailersRetailerCodeUnprocessableEntity creates a PatchRetailersRetailerCodeUnprocessableEntity with default headers values
func NewPatchRetailersRetailerCodeUnprocessableEntity() *PatchRetailersRetailerCodeUnprocessableEntity {
	return &PatchRetailersRetailerCodeUnprocessableEntity{}
}

/*PatchRetailersRetailerCodeUnprocessableEntity handles this case with default header values.

Retailer not updated. Returns an array of error messages explaining the problems with the provided attributes.
*/
type PatchRetailersRetailerCodeUnprocessableEntity struct {
	Payload *models_core.Retailer422Response
}

func (o *PatchRetailersRetailerCodeUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /retailers/{retailer_code}][%d] patchRetailersRetailerCodeUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PatchRetailersRetailerCodeUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.Retailer422Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
