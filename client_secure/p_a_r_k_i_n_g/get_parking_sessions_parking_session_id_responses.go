package p_a_r_k_i_n_g

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"restclient/models_secure"
)

// GetParkingSessionsParkingSessionIDReader is a Reader for the GetParkingSessionsParkingSessionID structure.
type GetParkingSessionsParkingSessionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetParkingSessionsParkingSessionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetParkingSessionsParkingSessionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetParkingSessionsParkingSessionIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetParkingSessionsParkingSessionIDOK creates a GetParkingSessionsParkingSessionIDOK with default headers values
func NewGetParkingSessionsParkingSessionIDOK() *GetParkingSessionsParkingSessionIDOK {
	return &GetParkingSessionsParkingSessionIDOK{}
}

/*GetParkingSessionsParkingSessionIDOK handles this case with default header values.

Successfully retrieved a single parking session.
*/
type GetParkingSessionsParkingSessionIDOK struct {
	Payload *models_secure.ParkingSessionResponse
}

func (o *GetParkingSessionsParkingSessionIDOK) Error() string {
	return fmt.Sprintf("[GET /parking/sessions/{parking_session_id}][%d] getParkingSessionsParkingSessionIdOK  %+v", 200, o.Payload)
}

func (o *GetParkingSessionsParkingSessionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.ParkingSessionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetParkingSessionsParkingSessionIDNotFound creates a GetParkingSessionsParkingSessionIDNotFound with default headers values
func NewGetParkingSessionsParkingSessionIDNotFound() *GetParkingSessionsParkingSessionIDNotFound {
	return &GetParkingSessionsParkingSessionIDNotFound{}
}

/*GetParkingSessionsParkingSessionIDNotFound handles this case with default header values.

Parking session not found
*/
type GetParkingSessionsParkingSessionIDNotFound struct {
	Payload *models_secure.Http404Response
}

func (o *GetParkingSessionsParkingSessionIDNotFound) Error() string {
	return fmt.Sprintf("[GET /parking/sessions/{parking_session_id}][%d] getParkingSessionsParkingSessionIdNotFound  %+v", 404, o.Payload)
}

func (o *GetParkingSessionsParkingSessionIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.Http404Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
