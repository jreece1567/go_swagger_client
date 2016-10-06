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

// GetParkingSessionsReader is a Reader for the GetParkingSessions structure.
type GetParkingSessionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetParkingSessionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetParkingSessionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetParkingSessionsOK creates a GetParkingSessionsOK with default headers values
func NewGetParkingSessionsOK() *GetParkingSessionsOK {
	return &GetParkingSessionsOK{}
}

/*GetParkingSessionsOK handles this case with default header values.

Successfully retrieved the parking activity sessions.
*/
type GetParkingSessionsOK struct {
	Payload *models_secure.ParkingActivitySessionResponse
}

func (o *GetParkingSessionsOK) Error() string {
	return fmt.Sprintf("[GET /parking/sessions][%d] getParkingSessionsOK  %+v", 200, o.Payload)
}

func (o *GetParkingSessionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.ParkingActivitySessionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}