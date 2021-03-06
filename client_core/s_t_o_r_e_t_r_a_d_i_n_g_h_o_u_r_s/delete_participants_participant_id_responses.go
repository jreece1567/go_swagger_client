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

// DeleteParticipantsParticipantIDReader is a Reader for the DeleteParticipantsParticipantID structure.
type DeleteParticipantsParticipantIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeleteParticipantsParticipantIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteParticipantsParticipantIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteParticipantsParticipantIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteParticipantsParticipantIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteParticipantsParticipantIDNoContent creates a DeleteParticipantsParticipantIDNoContent with default headers values
func NewDeleteParticipantsParticipantIDNoContent() *DeleteParticipantsParticipantIDNoContent {
	return &DeleteParticipantsParticipantIDNoContent{}
}

/*DeleteParticipantsParticipantIDNoContent handles this case with default header values.

Successfully deleted the participant.
*/
type DeleteParticipantsParticipantIDNoContent struct {
}

func (o *DeleteParticipantsParticipantIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /participants/{participant_id}][%d] deleteParticipantsParticipantIdNoContent ", 204)
}

func (o *DeleteParticipantsParticipantIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteParticipantsParticipantIDUnauthorized creates a DeleteParticipantsParticipantIDUnauthorized with default headers values
func NewDeleteParticipantsParticipantIDUnauthorized() *DeleteParticipantsParticipantIDUnauthorized {
	return &DeleteParticipantsParticipantIDUnauthorized{}
}

/*DeleteParticipantsParticipantIDUnauthorized handles this case with default header values.

Unauthorized request.
*/
type DeleteParticipantsParticipantIDUnauthorized struct {
	Payload *models_core.Http401Response
}

func (o *DeleteParticipantsParticipantIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /participants/{participant_id}][%d] deleteParticipantsParticipantIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteParticipantsParticipantIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.Http401Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteParticipantsParticipantIDNotFound creates a DeleteParticipantsParticipantIDNotFound with default headers values
func NewDeleteParticipantsParticipantIDNotFound() *DeleteParticipantsParticipantIDNotFound {
	return &DeleteParticipantsParticipantIDNotFound{}
}

/*DeleteParticipantsParticipantIDNotFound handles this case with default header values.

Participant not found.
*/
type DeleteParticipantsParticipantIDNotFound struct {
	Payload *models_core.Http404Response
}

func (o *DeleteParticipantsParticipantIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /participants/{participant_id}][%d] deleteParticipantsParticipantIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteParticipantsParticipantIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.Http404Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
