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

// GetParticipantsKindKindIDReader is a Reader for the GetParticipantsKindKindID structure.
type GetParticipantsKindKindIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetParticipantsKindKindIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetParticipantsKindKindIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetParticipantsKindKindIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetParticipantsKindKindIDOK creates a GetParticipantsKindKindIDOK with default headers values
func NewGetParticipantsKindKindIDOK() *GetParticipantsKindKindIDOK {
	return &GetParticipantsKindKindIDOK{}
}

/*GetParticipantsKindKindIDOK handles this case with default header values.

Successfully retrieved the individual participant.
*/
type GetParticipantsKindKindIDOK struct {
	Payload *models_core.ParticipantInstanceResponse
}

func (o *GetParticipantsKindKindIDOK) Error() string {
	return fmt.Sprintf("[GET /participants/{kind}/{kind_id}][%d] getParticipantsKindKindIdOK  %+v", 200, o.Payload)
}

func (o *GetParticipantsKindKindIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.ParticipantInstanceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetParticipantsKindKindIDNotFound creates a GetParticipantsKindKindIDNotFound with default headers values
func NewGetParticipantsKindKindIDNotFound() *GetParticipantsKindKindIDNotFound {
	return &GetParticipantsKindKindIDNotFound{}
}

/*GetParticipantsKindKindIDNotFound handles this case with default header values.

Participant not found.
*/
type GetParticipantsKindKindIDNotFound struct {
	Payload *models_core.Http404Response
}

func (o *GetParticipantsKindKindIDNotFound) Error() string {
	return fmt.Sprintf("[GET /participants/{kind}/{kind_id}][%d] getParticipantsKindKindIdNotFound  %+v", 404, o.Payload)
}

func (o *GetParticipantsKindKindIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.Http404Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}