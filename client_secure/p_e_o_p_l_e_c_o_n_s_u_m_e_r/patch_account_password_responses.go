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

// PatchAccountPasswordReader is a Reader for the PatchAccountPassword structure.
type PatchAccountPasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PatchAccountPasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPatchAccountPasswordNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPatchAccountPasswordBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPatchAccountPasswordUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPatchAccountPasswordUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchAccountPasswordNoContent creates a PatchAccountPasswordNoContent with default headers values
func NewPatchAccountPasswordNoContent() *PatchAccountPasswordNoContent {
	return &PatchAccountPasswordNoContent{}
}

/*PatchAccountPasswordNoContent handles this case with default header values.

Successfully updated account password.
*/
type PatchAccountPasswordNoContent struct {
}

func (o *PatchAccountPasswordNoContent) Error() string {
	return fmt.Sprintf("[PATCH /account/password][%d] patchAccountPasswordNoContent ", 204)
}

func (o *PatchAccountPasswordNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchAccountPasswordBadRequest creates a PatchAccountPasswordBadRequest with default headers values
func NewPatchAccountPasswordBadRequest() *PatchAccountPasswordBadRequest {
	return &PatchAccountPasswordBadRequest{}
}

/*PatchAccountPasswordBadRequest handles this case with default header values.

Bad request
*/
type PatchAccountPasswordBadRequest struct {
	Payload *models_secure.Person400Response
}

func (o *PatchAccountPasswordBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /account/password][%d] patchAccountPasswordBadRequest  %+v", 400, o.Payload)
}

func (o *PatchAccountPasswordBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.Person400Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAccountPasswordUnauthorized creates a PatchAccountPasswordUnauthorized with default headers values
func NewPatchAccountPasswordUnauthorized() *PatchAccountPasswordUnauthorized {
	return &PatchAccountPasswordUnauthorized{}
}

/*PatchAccountPasswordUnauthorized handles this case with default header values.

Unauthorized
*/
type PatchAccountPasswordUnauthorized struct {
	Payload *models_secure.UnauthorizedResponse
}

func (o *PatchAccountPasswordUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /account/password][%d] patchAccountPasswordUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchAccountPasswordUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.UnauthorizedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAccountPasswordUnprocessableEntity creates a PatchAccountPasswordUnprocessableEntity with default headers values
func NewPatchAccountPasswordUnprocessableEntity() *PatchAccountPasswordUnprocessableEntity {
	return &PatchAccountPasswordUnprocessableEntity{}
}

/*PatchAccountPasswordUnprocessableEntity handles this case with default header values.

Current or new password is invalid.
*/
type PatchAccountPasswordUnprocessableEntity struct {
	Payload *models_secure.UpdatePassword422Response
}

func (o *PatchAccountPasswordUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /account/password][%d] patchAccountPasswordUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PatchAccountPasswordUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.UpdatePassword422Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}