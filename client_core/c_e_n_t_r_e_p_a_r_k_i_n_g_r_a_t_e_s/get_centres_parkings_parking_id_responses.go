package c_e_n_t_r_e_p_a_r_k_i_n_g_r_a_t_e_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"restclient/models_core"
)

// GetCentresParkingsParkingIDReader is a Reader for the GetCentresParkingsParkingID structure.
type GetCentresParkingsParkingIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetCentresParkingsParkingIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCentresParkingsParkingIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetCentresParkingsParkingIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCentresParkingsParkingIDOK creates a GetCentresParkingsParkingIDOK with default headers values
func NewGetCentresParkingsParkingIDOK() *GetCentresParkingsParkingIDOK {
	return &GetCentresParkingsParkingIDOK{}
}

/*GetCentresParkingsParkingIDOK handles this case with default header values.

Successfully retrieved the individual centre parking rates.
*/
type GetCentresParkingsParkingIDOK struct {
	Payload *models_core.ParkingResponse
}

func (o *GetCentresParkingsParkingIDOK) Error() string {
	return fmt.Sprintf("[GET /centres/parkings/{parking_id}][%d] getCentresParkingsParkingIdOK  %+v", 200, o.Payload)
}

func (o *GetCentresParkingsParkingIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.ParkingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCentresParkingsParkingIDNotFound creates a GetCentresParkingsParkingIDNotFound with default headers values
func NewGetCentresParkingsParkingIDNotFound() *GetCentresParkingsParkingIDNotFound {
	return &GetCentresParkingsParkingIDNotFound{}
}

/*GetCentresParkingsParkingIDNotFound handles this case with default header values.

Centre parking not found.
*/
type GetCentresParkingsParkingIDNotFound struct {
	Payload *models_core.Http404Response
}

func (o *GetCentresParkingsParkingIDNotFound) Error() string {
	return fmt.Sprintf("[GET /centres/parkings/{parking_id}][%d] getCentresParkingsParkingIdNotFound  %+v", 404, o.Payload)
}

func (o *GetCentresParkingsParkingIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.Http404Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
