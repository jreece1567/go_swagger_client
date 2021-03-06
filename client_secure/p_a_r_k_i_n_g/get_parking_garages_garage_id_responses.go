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

// GetParkingGaragesGarageIDReader is a Reader for the GetParkingGaragesGarageID structure.
type GetParkingGaragesGarageIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetParkingGaragesGarageIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetParkingGaragesGarageIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetParkingGaragesGarageIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetParkingGaragesGarageIDOK creates a GetParkingGaragesGarageIDOK with default headers values
func NewGetParkingGaragesGarageIDOK() *GetParkingGaragesGarageIDOK {
	return &GetParkingGaragesGarageIDOK{}
}

/*GetParkingGaragesGarageIDOK handles this case with default header values.

Successfully retrieved garage data
*/
type GetParkingGaragesGarageIDOK struct {
	Payload *models_secure.GarageResponse
}

func (o *GetParkingGaragesGarageIDOK) Error() string {
	return fmt.Sprintf("[GET /parking/garages/{garage_id}][%d] getParkingGaragesGarageIdOK  %+v", 200, o.Payload)
}

func (o *GetParkingGaragesGarageIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.GarageResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetParkingGaragesGarageIDNotFound creates a GetParkingGaragesGarageIDNotFound with default headers values
func NewGetParkingGaragesGarageIDNotFound() *GetParkingGaragesGarageIDNotFound {
	return &GetParkingGaragesGarageIDNotFound{}
}

/*GetParkingGaragesGarageIDNotFound handles this case with default header values.

Parking Garage not found.
*/
type GetParkingGaragesGarageIDNotFound struct {
	Payload *models_secure.Http404Response
}

func (o *GetParkingGaragesGarageIDNotFound) Error() string {
	return fmt.Sprintf("[GET /parking/garages/{garage_id}][%d] getParkingGaragesGarageIdNotFound  %+v", 404, o.Payload)
}

func (o *GetParkingGaragesGarageIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.Http404Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
