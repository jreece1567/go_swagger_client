package s_t_o_r_e_t_r_a_d_i_n_g_h_o_u_r_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"restclient/models_core"
)

// PatchParticipantsParticipantIDReader is a Reader for the PatchParticipantsParticipantID structure.
type PatchParticipantsParticipantIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PatchParticipantsParticipantIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPatchParticipantsParticipantIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPatchParticipantsParticipantIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchParticipantsParticipantIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPatchParticipantsParticipantIDUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchParticipantsParticipantIDNoContent creates a PatchParticipantsParticipantIDNoContent with default headers values
func NewPatchParticipantsParticipantIDNoContent() *PatchParticipantsParticipantIDNoContent {
	return &PatchParticipantsParticipantIDNoContent{}
}

/*PatchParticipantsParticipantIDNoContent handles this case with default header values.

Successfully updated the participant.
*/
type PatchParticipantsParticipantIDNoContent struct {
}

func (o *PatchParticipantsParticipantIDNoContent) Error() string {
	return fmt.Sprintf("[PATCH /participants/{participant_id}][%d] patchParticipantsParticipantIdNoContent ", 204)
}

func (o *PatchParticipantsParticipantIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchParticipantsParticipantIDUnauthorized creates a PatchParticipantsParticipantIDUnauthorized with default headers values
func NewPatchParticipantsParticipantIDUnauthorized() *PatchParticipantsParticipantIDUnauthorized {
	return &PatchParticipantsParticipantIDUnauthorized{}
}

/*PatchParticipantsParticipantIDUnauthorized handles this case with default header values.

Unauthorized request.
*/
type PatchParticipantsParticipantIDUnauthorized struct {
	Payload *models_core.Http401Response
}

func (o *PatchParticipantsParticipantIDUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /participants/{participant_id}][%d] patchParticipantsParticipantIdUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchParticipantsParticipantIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.Http401Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchParticipantsParticipantIDNotFound creates a PatchParticipantsParticipantIDNotFound with default headers values
func NewPatchParticipantsParticipantIDNotFound() *PatchParticipantsParticipantIDNotFound {
	return &PatchParticipantsParticipantIDNotFound{}
}

/*PatchParticipantsParticipantIDNotFound handles this case with default header values.

Participant not found.
*/
type PatchParticipantsParticipantIDNotFound struct {
	Payload *models_core.Http404Response
}

func (o *PatchParticipantsParticipantIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /participants/{participant_id}][%d] patchParticipantsParticipantIdNotFound  %+v", 404, o.Payload)
}

func (o *PatchParticipantsParticipantIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.Http404Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchParticipantsParticipantIDUnprocessableEntity creates a PatchParticipantsParticipantIDUnprocessableEntity with default headers values
func NewPatchParticipantsParticipantIDUnprocessableEntity() *PatchParticipantsParticipantIDUnprocessableEntity {
	return &PatchParticipantsParticipantIDUnprocessableEntity{}
}

/*PatchParticipantsParticipantIDUnprocessableEntity handles this case with default header values.

Participant not updated. Returns an array of error messages explaining the problems with the provided attributes.
*/
type PatchParticipantsParticipantIDUnprocessableEntity struct {
	Payload *models_core.Participant422Response
}

func (o *PatchParticipantsParticipantIDUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /participants/{participant_id}][%d] patchParticipantsParticipantIdUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PatchParticipantsParticipantIDUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.Participant422Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
