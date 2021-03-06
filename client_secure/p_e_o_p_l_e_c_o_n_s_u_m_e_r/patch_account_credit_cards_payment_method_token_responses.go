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

// PatchAccountCreditCardsPaymentMethodTokenReader is a Reader for the PatchAccountCreditCardsPaymentMethodToken structure.
type PatchAccountCreditCardsPaymentMethodTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PatchAccountCreditCardsPaymentMethodTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPatchAccountCreditCardsPaymentMethodTokenNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPatchAccountCreditCardsPaymentMethodTokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPatchAccountCreditCardsPaymentMethodTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchAccountCreditCardsPaymentMethodTokenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPatchAccountCreditCardsPaymentMethodTokenUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 423:
		result := NewPatchAccountCreditCardsPaymentMethodTokenLocked()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchAccountCreditCardsPaymentMethodTokenNoContent creates a PatchAccountCreditCardsPaymentMethodTokenNoContent with default headers values
func NewPatchAccountCreditCardsPaymentMethodTokenNoContent() *PatchAccountCreditCardsPaymentMethodTokenNoContent {
	return &PatchAccountCreditCardsPaymentMethodTokenNoContent{}
}

/*PatchAccountCreditCardsPaymentMethodTokenNoContent handles this case with default header values.

Successfully updated the credit card.
*/
type PatchAccountCreditCardsPaymentMethodTokenNoContent struct {
}

func (o *PatchAccountCreditCardsPaymentMethodTokenNoContent) Error() string {
	return fmt.Sprintf("[PATCH /account/credit_cards/{payment_method_token}][%d] patchAccountCreditCardsPaymentMethodTokenNoContent ", 204)
}

func (o *PatchAccountCreditCardsPaymentMethodTokenNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchAccountCreditCardsPaymentMethodTokenBadRequest creates a PatchAccountCreditCardsPaymentMethodTokenBadRequest with default headers values
func NewPatchAccountCreditCardsPaymentMethodTokenBadRequest() *PatchAccountCreditCardsPaymentMethodTokenBadRequest {
	return &PatchAccountCreditCardsPaymentMethodTokenBadRequest{}
}

/*PatchAccountCreditCardsPaymentMethodTokenBadRequest handles this case with default header values.

Bad request
*/
type PatchAccountCreditCardsPaymentMethodTokenBadRequest struct {
	Payload *models_secure.Person400Response
}

func (o *PatchAccountCreditCardsPaymentMethodTokenBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /account/credit_cards/{payment_method_token}][%d] patchAccountCreditCardsPaymentMethodTokenBadRequest  %+v", 400, o.Payload)
}

func (o *PatchAccountCreditCardsPaymentMethodTokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.Person400Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAccountCreditCardsPaymentMethodTokenUnauthorized creates a PatchAccountCreditCardsPaymentMethodTokenUnauthorized with default headers values
func NewPatchAccountCreditCardsPaymentMethodTokenUnauthorized() *PatchAccountCreditCardsPaymentMethodTokenUnauthorized {
	return &PatchAccountCreditCardsPaymentMethodTokenUnauthorized{}
}

/*PatchAccountCreditCardsPaymentMethodTokenUnauthorized handles this case with default header values.

Unauthorized request.
*/
type PatchAccountCreditCardsPaymentMethodTokenUnauthorized struct {
	Payload *models_secure.UnauthorizedResponse
}

func (o *PatchAccountCreditCardsPaymentMethodTokenUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /account/credit_cards/{payment_method_token}][%d] patchAccountCreditCardsPaymentMethodTokenUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchAccountCreditCardsPaymentMethodTokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.UnauthorizedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAccountCreditCardsPaymentMethodTokenNotFound creates a PatchAccountCreditCardsPaymentMethodTokenNotFound with default headers values
func NewPatchAccountCreditCardsPaymentMethodTokenNotFound() *PatchAccountCreditCardsPaymentMethodTokenNotFound {
	return &PatchAccountCreditCardsPaymentMethodTokenNotFound{}
}

/*PatchAccountCreditCardsPaymentMethodTokenNotFound handles this case with default header values.

Credit card not found.
*/
type PatchAccountCreditCardsPaymentMethodTokenNotFound struct {
	Payload *models_secure.Http404Response
}

func (o *PatchAccountCreditCardsPaymentMethodTokenNotFound) Error() string {
	return fmt.Sprintf("[PATCH /account/credit_cards/{payment_method_token}][%d] patchAccountCreditCardsPaymentMethodTokenNotFound  %+v", 404, o.Payload)
}

func (o *PatchAccountCreditCardsPaymentMethodTokenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.Http404Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAccountCreditCardsPaymentMethodTokenUnprocessableEntity creates a PatchAccountCreditCardsPaymentMethodTokenUnprocessableEntity with default headers values
func NewPatchAccountCreditCardsPaymentMethodTokenUnprocessableEntity() *PatchAccountCreditCardsPaymentMethodTokenUnprocessableEntity {
	return &PatchAccountCreditCardsPaymentMethodTokenUnprocessableEntity{}
}

/*PatchAccountCreditCardsPaymentMethodTokenUnprocessableEntity handles this case with default header values.

Credit card not updated. Returns an array of error messages explaining the problems with the provided attributes.
*/
type PatchAccountCreditCardsPaymentMethodTokenUnprocessableEntity struct {
	Payload *models_secure.CreditCard422Response
}

func (o *PatchAccountCreditCardsPaymentMethodTokenUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /account/credit_cards/{payment_method_token}][%d] patchAccountCreditCardsPaymentMethodTokenUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PatchAccountCreditCardsPaymentMethodTokenUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.CreditCard422Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAccountCreditCardsPaymentMethodTokenLocked creates a PatchAccountCreditCardsPaymentMethodTokenLocked with default headers values
func NewPatchAccountCreditCardsPaymentMethodTokenLocked() *PatchAccountCreditCardsPaymentMethodTokenLocked {
	return &PatchAccountCreditCardsPaymentMethodTokenLocked{}
}

/*PatchAccountCreditCardsPaymentMethodTokenLocked handles this case with default header values.

Resource not updated because it's locked.
*/
type PatchAccountCreditCardsPaymentMethodTokenLocked struct {
}

func (o *PatchAccountCreditCardsPaymentMethodTokenLocked) Error() string {
	return fmt.Sprintf("[PATCH /account/credit_cards/{payment_method_token}][%d] patchAccountCreditCardsPaymentMethodTokenLocked ", 423)
}

func (o *PatchAccountCreditCardsPaymentMethodTokenLocked) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
