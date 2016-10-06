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

// GetPaymentsMerchantsMerchantIDReader is a Reader for the GetPaymentsMerchantsMerchantID structure.
type GetPaymentsMerchantsMerchantIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetPaymentsMerchantsMerchantIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentsMerchantsMerchantIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetPaymentsMerchantsMerchantIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetPaymentsMerchantsMerchantIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentsMerchantsMerchantIDOK creates a GetPaymentsMerchantsMerchantIDOK with default headers values
func NewGetPaymentsMerchantsMerchantIDOK() *GetPaymentsMerchantsMerchantIDOK {
	return &GetPaymentsMerchantsMerchantIDOK{}
}

/*GetPaymentsMerchantsMerchantIDOK handles this case with default header values.

Successfully retrieved the individual merchant.
*/
type GetPaymentsMerchantsMerchantIDOK struct {
	Payload *models_secure.MerchantResponse
}

func (o *GetPaymentsMerchantsMerchantIDOK) Error() string {
	return fmt.Sprintf("[GET /payments/merchants/{merchant_id}][%d] getPaymentsMerchantsMerchantIdOK  %+v", 200, o.Payload)
}

func (o *GetPaymentsMerchantsMerchantIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.MerchantResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentsMerchantsMerchantIDUnauthorized creates a GetPaymentsMerchantsMerchantIDUnauthorized with default headers values
func NewGetPaymentsMerchantsMerchantIDUnauthorized() *GetPaymentsMerchantsMerchantIDUnauthorized {
	return &GetPaymentsMerchantsMerchantIDUnauthorized{}
}

/*GetPaymentsMerchantsMerchantIDUnauthorized handles this case with default header values.

Unauthorized request
*/
type GetPaymentsMerchantsMerchantIDUnauthorized struct {
	Payload *models_secure.Payment401Response
}

func (o *GetPaymentsMerchantsMerchantIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /payments/merchants/{merchant_id}][%d] getPaymentsMerchantsMerchantIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPaymentsMerchantsMerchantIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.Payment401Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentsMerchantsMerchantIDNotFound creates a GetPaymentsMerchantsMerchantIDNotFound with default headers values
func NewGetPaymentsMerchantsMerchantIDNotFound() *GetPaymentsMerchantsMerchantIDNotFound {
	return &GetPaymentsMerchantsMerchantIDNotFound{}
}

/*GetPaymentsMerchantsMerchantIDNotFound handles this case with default header values.

Merchant not found.
*/
type GetPaymentsMerchantsMerchantIDNotFound struct {
	Payload *models_secure.Http404Response
}

func (o *GetPaymentsMerchantsMerchantIDNotFound) Error() string {
	return fmt.Sprintf("[GET /payments/merchants/{merchant_id}][%d] getPaymentsMerchantsMerchantIdNotFound  %+v", 404, o.Payload)
}

func (o *GetPaymentsMerchantsMerchantIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.Http404Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
