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

// PostParkingGaragesReader is a Reader for the PostParkingGarages structure.
type PostParkingGaragesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PostParkingGaragesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostParkingGaragesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 422:
		result := NewPostParkingGaragesUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostParkingGaragesCreated creates a PostParkingGaragesCreated with default headers values
func NewPostParkingGaragesCreated() *PostParkingGaragesCreated {
	return &PostParkingGaragesCreated{}
}

/*PostParkingGaragesCreated handles this case with default header values.

Successfully created the parking garage.
*/
type PostParkingGaragesCreated struct {
	Payload *models_secure.GarageResponse
}

func (o *PostParkingGaragesCreated) Error() string {
	return fmt.Sprintf("[POST /parking/garages][%d] postParkingGaragesCreated  %+v", 201, o.Payload)
}

func (o *PostParkingGaragesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.GarageResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostParkingGaragesUnprocessableEntity creates a PostParkingGaragesUnprocessableEntity with default headers values
func NewPostParkingGaragesUnprocessableEntity() *PostParkingGaragesUnprocessableEntity {
	return &PostParkingGaragesUnprocessableEntity{}
}

/*PostParkingGaragesUnprocessableEntity handles this case with default header values.

Parking garage not created. Returns an array of error messages explaining the problems with the provided attributes.
*/
type PostParkingGaragesUnprocessableEntity struct {
	Payload *models_secure.Garages422Response
}

func (o *PostParkingGaragesUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /parking/garages][%d] postParkingGaragesUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostParkingGaragesUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.Garages422Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
