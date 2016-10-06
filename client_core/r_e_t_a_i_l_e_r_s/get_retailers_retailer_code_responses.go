package r_e_t_a_i_l_e_r_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"restclient/models_core"
)

// GetRetailersRetailerCodeReader is a Reader for the GetRetailersRetailerCode structure.
type GetRetailersRetailerCodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetRetailersRetailerCodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRetailersRetailerCodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetRetailersRetailerCodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRetailersRetailerCodeOK creates a GetRetailersRetailerCodeOK with default headers values
func NewGetRetailersRetailerCodeOK() *GetRetailersRetailerCodeOK {
	return &GetRetailersRetailerCodeOK{}
}

/*GetRetailersRetailerCodeOK handles this case with default header values.

Successfully retrieved the individual retailer.
*/
type GetRetailersRetailerCodeOK struct {
	Payload *models_core.RetailerResponse
}

func (o *GetRetailersRetailerCodeOK) Error() string {
	return fmt.Sprintf("[GET /retailers/{retailer_code}][%d] getRetailersRetailerCodeOK  %+v", 200, o.Payload)
}

func (o *GetRetailersRetailerCodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.RetailerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRetailersRetailerCodeNotFound creates a GetRetailersRetailerCodeNotFound with default headers values
func NewGetRetailersRetailerCodeNotFound() *GetRetailersRetailerCodeNotFound {
	return &GetRetailersRetailerCodeNotFound{}
}

/*GetRetailersRetailerCodeNotFound handles this case with default header values.

Retailer not found.
*/
type GetRetailersRetailerCodeNotFound struct {
	Payload *models_core.Http404Response
}

func (o *GetRetailersRetailerCodeNotFound) Error() string {
	return fmt.Sprintf("[GET /retailers/{retailer_code}][%d] getRetailersRetailerCodeNotFound  %+v", 404, o.Payload)
}

func (o *GetRetailersRetailerCodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.Http404Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}