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

// GetPaymentsMerchantsReader is a Reader for the GetPaymentsMerchants structure.
type GetPaymentsMerchantsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetPaymentsMerchantsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentsMerchantsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetPaymentsMerchantsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentsMerchantsOK creates a GetPaymentsMerchantsOK with default headers values
func NewGetPaymentsMerchantsOK() *GetPaymentsMerchantsOK {
	return &GetPaymentsMerchantsOK{}
}

/*GetPaymentsMerchantsOK handles this case with default header values.

Successfully retrieved the list of merchants.
*/
type GetPaymentsMerchantsOK struct {
	Payload *models_secure.MerchantsListResponse
}

func (o *GetPaymentsMerchantsOK) Error() string {
	return fmt.Sprintf("[GET /payments/merchants][%d] getPaymentsMerchantsOK  %+v", 200, o.Payload)
}

func (o *GetPaymentsMerchantsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.MerchantsListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentsMerchantsUnauthorized creates a GetPaymentsMerchantsUnauthorized with default headers values
func NewGetPaymentsMerchantsUnauthorized() *GetPaymentsMerchantsUnauthorized {
	return &GetPaymentsMerchantsUnauthorized{}
}

/*GetPaymentsMerchantsUnauthorized handles this case with default header values.

Unauthorized request
*/
type GetPaymentsMerchantsUnauthorized struct {
	Payload *models_secure.Payment401Response
}

func (o *GetPaymentsMerchantsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /payments/merchants][%d] getPaymentsMerchantsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPaymentsMerchantsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.Payment401Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}