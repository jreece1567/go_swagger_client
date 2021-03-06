package p_e_o_p_l_e_m_a_n_a_g_e_m_e_n_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"restclient/models_secure"
)

// PatchPeoplePersonIDReader is a Reader for the PatchPeoplePersonID structure.
type PatchPeoplePersonIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PatchPeoplePersonIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPatchPeoplePersonIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPatchPeoplePersonIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPatchPeoplePersonIDUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchPeoplePersonIDNoContent creates a PatchPeoplePersonIDNoContent with default headers values
func NewPatchPeoplePersonIDNoContent() *PatchPeoplePersonIDNoContent {
	return &PatchPeoplePersonIDNoContent{}
}

/*PatchPeoplePersonIDNoContent handles this case with default header values.

Updated with success
*/
type PatchPeoplePersonIDNoContent struct {
}

func (o *PatchPeoplePersonIDNoContent) Error() string {
	return fmt.Sprintf("[PATCH /people/{person_id}][%d] patchPeoplePersonIdNoContent ", 204)
}

func (o *PatchPeoplePersonIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchPeoplePersonIDUnauthorized creates a PatchPeoplePersonIDUnauthorized with default headers values
func NewPatchPeoplePersonIDUnauthorized() *PatchPeoplePersonIDUnauthorized {
	return &PatchPeoplePersonIDUnauthorized{}
}

/*PatchPeoplePersonIDUnauthorized handles this case with default header values.

Unauthorized
*/
type PatchPeoplePersonIDUnauthorized struct {
	Payload *models_secure.UnauthorizedResponse
}

func (o *PatchPeoplePersonIDUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /people/{person_id}][%d] patchPeoplePersonIdUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchPeoplePersonIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.UnauthorizedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchPeoplePersonIDUnprocessableEntity creates a PatchPeoplePersonIDUnprocessableEntity with default headers values
func NewPatchPeoplePersonIDUnprocessableEntity() *PatchPeoplePersonIDUnprocessableEntity {
	return &PatchPeoplePersonIDUnprocessableEntity{}
}

/*PatchPeoplePersonIDUnprocessableEntity handles this case with default header values.

Unprocessable entity
*/
type PatchPeoplePersonIDUnprocessableEntity struct {
	Payload *models_secure.Person422Response
}

func (o *PatchPeoplePersonIDUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /people/{person_id}][%d] patchPeoplePersonIdUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PatchPeoplePersonIDUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.Person422Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
