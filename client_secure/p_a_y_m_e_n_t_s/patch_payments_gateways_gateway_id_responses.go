package p_a_y_m_e_n_t_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"restclient/models_secure"
)

// PatchPaymentsGatewaysGatewayIDReader is a Reader for the PatchPaymentsGatewaysGatewayID structure.
type PatchPaymentsGatewaysGatewayIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PatchPaymentsGatewaysGatewayIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPatchPaymentsGatewaysGatewayIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPatchPaymentsGatewaysGatewayIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchPaymentsGatewaysGatewayIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPatchPaymentsGatewaysGatewayIDUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchPaymentsGatewaysGatewayIDNoContent creates a PatchPaymentsGatewaysGatewayIDNoContent with default headers values
func NewPatchPaymentsGatewaysGatewayIDNoContent() *PatchPaymentsGatewaysGatewayIDNoContent {
	return &PatchPaymentsGatewaysGatewayIDNoContent{}
}

/*PatchPaymentsGatewaysGatewayIDNoContent handles this case with default header values.

Successfully updated the gateway.
*/
type PatchPaymentsGatewaysGatewayIDNoContent struct {
}

func (o *PatchPaymentsGatewaysGatewayIDNoContent) Error() string {
	return fmt.Sprintf("[PATCH /payments/gateways/{gateway_id}][%d] patchPaymentsGatewaysGatewayIdNoContent ", 204)
}

func (o *PatchPaymentsGatewaysGatewayIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchPaymentsGatewaysGatewayIDUnauthorized creates a PatchPaymentsGatewaysGatewayIDUnauthorized with default headers values
func NewPatchPaymentsGatewaysGatewayIDUnauthorized() *PatchPaymentsGatewaysGatewayIDUnauthorized {
	return &PatchPaymentsGatewaysGatewayIDUnauthorized{}
}

/*PatchPaymentsGatewaysGatewayIDUnauthorized handles this case with default header values.

Unauthorized request
*/
type PatchPaymentsGatewaysGatewayIDUnauthorized struct {
	Payload *models_secure.Payment401Response
}

func (o *PatchPaymentsGatewaysGatewayIDUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /payments/gateways/{gateway_id}][%d] patchPaymentsGatewaysGatewayIdUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchPaymentsGatewaysGatewayIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.Payment401Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchPaymentsGatewaysGatewayIDNotFound creates a PatchPaymentsGatewaysGatewayIDNotFound with default headers values
func NewPatchPaymentsGatewaysGatewayIDNotFound() *PatchPaymentsGatewaysGatewayIDNotFound {
	return &PatchPaymentsGatewaysGatewayIDNotFound{}
}

/*PatchPaymentsGatewaysGatewayIDNotFound handles this case with default header values.

Gateway not found.
*/
type PatchPaymentsGatewaysGatewayIDNotFound struct {
	Payload *models_secure.Http404Response
}

func (o *PatchPaymentsGatewaysGatewayIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /payments/gateways/{gateway_id}][%d] patchPaymentsGatewaysGatewayIdNotFound  %+v", 404, o.Payload)
}

func (o *PatchPaymentsGatewaysGatewayIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.Http404Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchPaymentsGatewaysGatewayIDUnprocessableEntity creates a PatchPaymentsGatewaysGatewayIDUnprocessableEntity with default headers values
func NewPatchPaymentsGatewaysGatewayIDUnprocessableEntity() *PatchPaymentsGatewaysGatewayIDUnprocessableEntity {
	return &PatchPaymentsGatewaysGatewayIDUnprocessableEntity{}
}

/*PatchPaymentsGatewaysGatewayIDUnprocessableEntity handles this case with default header values.

Gateway not updated. Returns an array of error messages explaining the problems with the provided attributes.
*/
type PatchPaymentsGatewaysGatewayIDUnprocessableEntity struct {
	Payload *models_secure.Gateway422Response
}

func (o *PatchPaymentsGatewaysGatewayIDUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /payments/gateways/{gateway_id}][%d] patchPaymentsGatewaysGatewayIdUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PatchPaymentsGatewaysGatewayIDUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.Gateway422Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
