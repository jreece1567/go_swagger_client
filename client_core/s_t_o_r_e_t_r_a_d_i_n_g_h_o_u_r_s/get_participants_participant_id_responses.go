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

// GetParticipantsParticipantIDReader is a Reader for the GetParticipantsParticipantID structure.
type GetParticipantsParticipantIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetParticipantsParticipantIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetParticipantsParticipantIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetParticipantsParticipantIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetParticipantsParticipantIDOK creates a GetParticipantsParticipantIDOK with default headers values
func NewGetParticipantsParticipantIDOK() *GetParticipantsParticipantIDOK {
	return &GetParticipantsParticipantIDOK{}
}

/*GetParticipantsParticipantIDOK handles this case with default header values.

Successfully retrieved the individual participant.
*/
type GetParticipantsParticipantIDOK struct {
	Payload *models_core.ParticipantInstanceResponse
}

func (o *GetParticipantsParticipantIDOK) Error() string {
	return fmt.Sprintf("[GET /participants/{participant_id}][%d] getParticipantsParticipantIdOK  %+v", 200, o.Payload)
}

func (o *GetParticipantsParticipantIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.ParticipantInstanceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetParticipantsParticipantIDNotFound creates a GetParticipantsParticipantIDNotFound with default headers values
func NewGetParticipantsParticipantIDNotFound() *GetParticipantsParticipantIDNotFound {
	return &GetParticipantsParticipantIDNotFound{}
}

/*GetParticipantsParticipantIDNotFound handles this case with default header values.

Participant not found.
*/
type GetParticipantsParticipantIDNotFound struct {
	Payload *models_core.Http404Response
}

func (o *GetParticipantsParticipantIDNotFound) Error() string {
	return fmt.Sprintf("[GET /participants/{participant_id}][%d] getParticipantsParticipantIdNotFound  %+v", 404, o.Payload)
}

func (o *GetParticipantsParticipantIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.Http404Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}