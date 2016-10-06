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

// PostParticipantsReader is a Reader for the PostParticipants structure.
type PostParticipantsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PostParticipantsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostParticipantsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostParticipantsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPostParticipantsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostParticipantsCreated creates a PostParticipantsCreated with default headers values
func NewPostParticipantsCreated() *PostParticipantsCreated {
	return &PostParticipantsCreated{}
}

/*PostParticipantsCreated handles this case with default header values.

Successfully created a participant.
*/
type PostParticipantsCreated struct {
	/*URI of created participant.
	 */
	Location string

	Payload *models_core.ParticipantInstanceResponse
}

func (o *PostParticipantsCreated) Error() string {
	return fmt.Sprintf("[POST /participants][%d] postParticipantsCreated  %+v", 201, o.Payload)
}

func (o *PostParticipantsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header location
	o.Location = response.GetHeader("location")

	o.Payload = new(models_core.ParticipantInstanceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostParticipantsUnauthorized creates a PostParticipantsUnauthorized with default headers values
func NewPostParticipantsUnauthorized() *PostParticipantsUnauthorized {
	return &PostParticipantsUnauthorized{}
}

/*PostParticipantsUnauthorized handles this case with default header values.

Unauthorized request.
*/
type PostParticipantsUnauthorized struct {
	Payload *models_core.Http401Response
}

func (o *PostParticipantsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /participants][%d] postParticipantsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostParticipantsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.Http401Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostParticipantsUnprocessableEntity creates a PostParticipantsUnprocessableEntity with default headers values
func NewPostParticipantsUnprocessableEntity() *PostParticipantsUnprocessableEntity {
	return &PostParticipantsUnprocessableEntity{}
}

/*PostParticipantsUnprocessableEntity handles this case with default header values.

Participant not created. Returns an array of error messages explaining the problems with the provided attributes.
*/
type PostParticipantsUnprocessableEntity struct {
	Payload *models_core.Participant422Response
}

func (o *PostParticipantsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /participants][%d] postParticipantsUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostParticipantsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.Participant422Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
