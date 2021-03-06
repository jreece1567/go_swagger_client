package p_a_y_m_e_n_t_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"restclient/models_secure"
)

// GetPaymentsGatewaysReader is a Reader for the GetPaymentsGateways structure.
type GetPaymentsGatewaysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetPaymentsGatewaysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentsGatewaysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetPaymentsGatewaysUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentsGatewaysOK creates a GetPaymentsGatewaysOK with default headers values
func NewGetPaymentsGatewaysOK() *GetPaymentsGatewaysOK {
	return &GetPaymentsGatewaysOK{}
}

/*GetPaymentsGatewaysOK handles this case with default header values.

Successfully retrieved the list of gateways.
*/
type GetPaymentsGatewaysOK struct {
	Payload *models_secure.GatewaysListResponse
}

func (o *GetPaymentsGatewaysOK) Error() string {
	return fmt.Sprintf("[GET /payments/gateways][%d] getPaymentsGatewaysOK  %+v", 200, o.Payload)
}

func (o *GetPaymentsGatewaysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.GatewaysListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentsGatewaysUnauthorized creates a GetPaymentsGatewaysUnauthorized with default headers values
func NewGetPaymentsGatewaysUnauthorized() *GetPaymentsGatewaysUnauthorized {
	return &GetPaymentsGatewaysUnauthorized{}
}

/*GetPaymentsGatewaysUnauthorized handles this case with default header values.

Unauthorized request
*/
type GetPaymentsGatewaysUnauthorized struct {
	Payload *models_secure.Payment401Response
}

func (o *GetPaymentsGatewaysUnauthorized) Error() string {
	return fmt.Sprintf("[GET /payments/gateways][%d] getPaymentsGatewaysUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPaymentsGatewaysUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.Payment401Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
