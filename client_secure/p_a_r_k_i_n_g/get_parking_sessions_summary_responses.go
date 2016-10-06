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

// GetParkingSessionsSummaryReader is a Reader for the GetParkingSessionsSummary structure.
type GetParkingSessionsSummaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetParkingSessionsSummaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetParkingSessionsSummaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 422:
		result := NewGetParkingSessionsSummaryUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetParkingSessionsSummaryOK creates a GetParkingSessionsSummaryOK with default headers values
func NewGetParkingSessionsSummaryOK() *GetParkingSessionsSummaryOK {
	return &GetParkingSessionsSummaryOK{}
}

/*GetParkingSessionsSummaryOK handles this case with default header values.

Successfully retrieved the summary.
*/
type GetParkingSessionsSummaryOK struct {
	Payload *models_secure.ParkingSessionSummaryResponse
}

func (o *GetParkingSessionsSummaryOK) Error() string {
	return fmt.Sprintf("[GET /parking/sessions/summary][%d] getParkingSessionsSummaryOK  %+v", 200, o.Payload)
}

func (o *GetParkingSessionsSummaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.ParkingSessionSummaryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetParkingSessionsSummaryUnprocessableEntity creates a GetParkingSessionsSummaryUnprocessableEntity with default headers values
func NewGetParkingSessionsSummaryUnprocessableEntity() *GetParkingSessionsSummaryUnprocessableEntity {
	return &GetParkingSessionsSummaryUnprocessableEntity{}
}

/*GetParkingSessionsSummaryUnprocessableEntity handles this case with default header values.

Missing required parameters.
*/
type GetParkingSessionsSummaryUnprocessableEntity struct {
	Payload *models_secure.ParkingSessionSummary422Response
}

func (o *GetParkingSessionsSummaryUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /parking/sessions/summary][%d] getParkingSessionsSummaryUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetParkingSessionsSummaryUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.ParkingSessionSummary422Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}