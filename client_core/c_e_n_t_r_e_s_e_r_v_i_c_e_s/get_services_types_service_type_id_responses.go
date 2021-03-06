package c_e_n_t_r_e_s_e_r_v_i_c_e_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"restclient/models_core"
)

// GetServicesTypesServiceTypeIDReader is a Reader for the GetServicesTypesServiceTypeID structure.
type GetServicesTypesServiceTypeIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetServicesTypesServiceTypeIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetServicesTypesServiceTypeIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetServicesTypesServiceTypeIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetServicesTypesServiceTypeIDOK creates a GetServicesTypesServiceTypeIDOK with default headers values
func NewGetServicesTypesServiceTypeIDOK() *GetServicesTypesServiceTypeIDOK {
	return &GetServicesTypesServiceTypeIDOK{}
}

/*GetServicesTypesServiceTypeIDOK handles this case with default header values.

Successfully retrieved the individual service-type.
*/
type GetServicesTypesServiceTypeIDOK struct {
	Payload *models_core.ServiceTypeResponse
}

func (o *GetServicesTypesServiceTypeIDOK) Error() string {
	return fmt.Sprintf("[GET /services/types/{service_type_id}][%d] getServicesTypesServiceTypeIdOK  %+v", 200, o.Payload)
}

func (o *GetServicesTypesServiceTypeIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.ServiceTypeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetServicesTypesServiceTypeIDNotFound creates a GetServicesTypesServiceTypeIDNotFound with default headers values
func NewGetServicesTypesServiceTypeIDNotFound() *GetServicesTypesServiceTypeIDNotFound {
	return &GetServicesTypesServiceTypeIDNotFound{}
}

/*GetServicesTypesServiceTypeIDNotFound handles this case with default header values.

Service-type not found.
*/
type GetServicesTypesServiceTypeIDNotFound struct {
	Payload *models_core.Http404Response
}

func (o *GetServicesTypesServiceTypeIDNotFound) Error() string {
	return fmt.Sprintf("[GET /services/types/{service_type_id}][%d] getServicesTypesServiceTypeIdNotFound  %+v", 404, o.Payload)
}

func (o *GetServicesTypesServiceTypeIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.Http404Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
