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

// PatchPaymentsMerchantsMerchantIDReader is a Reader for the PatchPaymentsMerchantsMerchantID structure.
type PatchPaymentsMerchantsMerchantIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PatchPaymentsMerchantsMerchantIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPatchPaymentsMerchantsMerchantIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPatchPaymentsMerchantsMerchantIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchPaymentsMerchantsMerchantIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPatchPaymentsMerchantsMerchantIDUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchPaymentsMerchantsMerchantIDNoContent creates a PatchPaymentsMerchantsMerchantIDNoContent with default headers values
func NewPatchPaymentsMerchantsMerchantIDNoContent() *PatchPaymentsMerchantsMerchantIDNoContent {
	return &PatchPaymentsMerchantsMerchantIDNoContent{}
}

/*PatchPaymentsMerchantsMerchantIDNoContent handles this case with default header values.

Successfully updated the merchant.
*/
type PatchPaymentsMerchantsMerchantIDNoContent struct {
}

func (o *PatchPaymentsMerchantsMerchantIDNoContent) Error() string {
	return fmt.Sprintf("[PATCH /payments/merchants/{merchant_id}][%d] patchPaymentsMerchantsMerchantIdNoContent ", 204)
}

func (o *PatchPaymentsMerchantsMerchantIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchPaymentsMerchantsMerchantIDUnauthorized creates a PatchPaymentsMerchantsMerchantIDUnauthorized with default headers values
func NewPatchPaymentsMerchantsMerchantIDUnauthorized() *PatchPaymentsMerchantsMerchantIDUnauthorized {
	return &PatchPaymentsMerchantsMerchantIDUnauthorized{}
}

/*PatchPaymentsMerchantsMerchantIDUnauthorized handles this case with default header values.

Unauthorized request
*/
type PatchPaymentsMerchantsMerchantIDUnauthorized struct {
	Payload *models_secure.Payment401Response
}

func (o *PatchPaymentsMerchantsMerchantIDUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /payments/merchants/{merchant_id}][%d] patchPaymentsMerchantsMerchantIdUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchPaymentsMerchantsMerchantIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.Payment401Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchPaymentsMerchantsMerchantIDNotFound creates a PatchPaymentsMerchantsMerchantIDNotFound with default headers values
func NewPatchPaymentsMerchantsMerchantIDNotFound() *PatchPaymentsMerchantsMerchantIDNotFound {
	return &PatchPaymentsMerchantsMerchantIDNotFound{}
}

/*PatchPaymentsMerchantsMerchantIDNotFound handles this case with default header values.

Merchant not found.
*/
type PatchPaymentsMerchantsMerchantIDNotFound struct {
	Payload *models_secure.Http404Response
}

func (o *PatchPaymentsMerchantsMerchantIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /payments/merchants/{merchant_id}][%d] patchPaymentsMerchantsMerchantIdNotFound  %+v", 404, o.Payload)
}

func (o *PatchPaymentsMerchantsMerchantIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.Http404Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchPaymentsMerchantsMerchantIDUnprocessableEntity creates a PatchPaymentsMerchantsMerchantIDUnprocessableEntity with default headers values
func NewPatchPaymentsMerchantsMerchantIDUnprocessableEntity() *PatchPaymentsMerchantsMerchantIDUnprocessableEntity {
	return &PatchPaymentsMerchantsMerchantIDUnprocessableEntity{}
}

/*PatchPaymentsMerchantsMerchantIDUnprocessableEntity handles this case with default header values.

Merchant not updated. Returns an array of error messages explaining the problems with the provided attributes.
*/
type PatchPaymentsMerchantsMerchantIDUnprocessableEntity struct {
	Payload *models_secure.Merchant422Response
}

func (o *PatchPaymentsMerchantsMerchantIDUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /payments/merchants/{merchant_id}][%d] patchPaymentsMerchantsMerchantIdUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PatchPaymentsMerchantsMerchantIDUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.Merchant422Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
