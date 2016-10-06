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

// GetCentresCentreIDVehiclesVehicleIDReader is a Reader for the GetCentresCentreIDVehiclesVehicleID structure.
type GetCentresCentreIDVehiclesVehicleIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetCentresCentreIDVehiclesVehicleIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCentresCentreIDVehiclesVehicleIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetCentresCentreIDVehiclesVehicleIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 502:
		result := NewGetCentresCentreIDVehiclesVehicleIDBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCentresCentreIDVehiclesVehicleIDOK creates a GetCentresCentreIDVehiclesVehicleIDOK with default headers values
func NewGetCentresCentreIDVehiclesVehicleIDOK() *GetCentresCentreIDVehiclesVehicleIDOK {
	return &GetCentresCentreIDVehiclesVehicleIDOK{}
}

/*GetCentresCentreIDVehiclesVehicleIDOK handles this case with default header values.

Successfully retrieved the vehicle information.
*/
type GetCentresCentreIDVehiclesVehicleIDOK struct {
	Payload *models_core.VehicleResponse
}

func (o *GetCentresCentreIDVehiclesVehicleIDOK) Error() string {
	return fmt.Sprintf("[GET /centres/{centre_id}/vehicles/{vehicle_id}][%d] getCentresCentreIdVehiclesVehicleIdOK  %+v", 200, o.Payload)
}

func (o *GetCentresCentreIDVehiclesVehicleIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.VehicleResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCentresCentreIDVehiclesVehicleIDNotFound creates a GetCentresCentreIDVehiclesVehicleIDNotFound with default headers values
func NewGetCentresCentreIDVehiclesVehicleIDNotFound() *GetCentresCentreIDVehiclesVehicleIDNotFound {
	return &GetCentresCentreIDVehiclesVehicleIDNotFound{}
}

/*GetCentresCentreIDVehiclesVehicleIDNotFound handles this case with default header values.

Vehicle not found.
*/
type GetCentresCentreIDVehiclesVehicleIDNotFound struct {
	Payload *models_core.Http404Response
}

func (o *GetCentresCentreIDVehiclesVehicleIDNotFound) Error() string {
	return fmt.Sprintf("[GET /centres/{centre_id}/vehicles/{vehicle_id}][%d] getCentresCentreIdVehiclesVehicleIdNotFound  %+v", 404, o.Payload)
}

func (o *GetCentresCentreIDVehiclesVehicleIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.Http404Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCentresCentreIDVehiclesVehicleIDBadGateway creates a GetCentresCentreIDVehiclesVehicleIDBadGateway with default headers values
func NewGetCentresCentreIDVehiclesVehicleIDBadGateway() *GetCentresCentreIDVehiclesVehicleIDBadGateway {
	return &GetCentresCentreIDVehiclesVehicleIDBadGateway{}
}

/*GetCentresCentreIDVehiclesVehicleIDBadGateway handles this case with default header values.

Bad Gateway, error or unexpected response received from upstream server.
*/
type GetCentresCentreIDVehiclesVehicleIDBadGateway struct {
	Payload *models_core.Http502Response
}

func (o *GetCentresCentreIDVehiclesVehicleIDBadGateway) Error() string {
	return fmt.Sprintf("[GET /centres/{centre_id}/vehicles/{vehicle_id}][%d] getCentresCentreIdVehiclesVehicleIdBadGateway  %+v", 502, o.Payload)
}

func (o *GetCentresCentreIDVehiclesVehicleIDBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.Http502Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}