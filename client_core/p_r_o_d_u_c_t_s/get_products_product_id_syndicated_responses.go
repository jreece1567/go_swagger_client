package p_r_o_d_u_c_t_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"restclient/models_core"
)

// GetProductsProductIDSyndicatedReader is a Reader for the GetProductsProductIDSyndicated structure.
type GetProductsProductIDSyndicatedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetProductsProductIDSyndicatedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetProductsProductIDSyndicatedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetProductsProductIDSyndicatedNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetProductsProductIDSyndicatedOK creates a GetProductsProductIDSyndicatedOK with default headers values
func NewGetProductsProductIDSyndicatedOK() *GetProductsProductIDSyndicatedOK {
	return &GetProductsProductIDSyndicatedOK{}
}

/*GetProductsProductIDSyndicatedOK handles this case with default header values.

Successfully retrieved the individual syndicated product.
*/
type GetProductsProductIDSyndicatedOK struct {
	Payload *models_core.SyndicatedProductResponse
}

func (o *GetProductsProductIDSyndicatedOK) Error() string {
	return fmt.Sprintf("[GET /products/{product_id}/syndicated][%d] getProductsProductIdSyndicatedOK  %+v", 200, o.Payload)
}

func (o *GetProductsProductIDSyndicatedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.SyndicatedProductResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProductsProductIDSyndicatedNotFound creates a GetProductsProductIDSyndicatedNotFound with default headers values
func NewGetProductsProductIDSyndicatedNotFound() *GetProductsProductIDSyndicatedNotFound {
	return &GetProductsProductIDSyndicatedNotFound{}
}

/*GetProductsProductIDSyndicatedNotFound handles this case with default header values.

Syndicated product not found.
*/
type GetProductsProductIDSyndicatedNotFound struct {
	Payload *models_core.SyndicatedProducthttp404Response
}

func (o *GetProductsProductIDSyndicatedNotFound) Error() string {
	return fmt.Sprintf("[GET /products/{product_id}/syndicated][%d] getProductsProductIdSyndicatedNotFound  %+v", 404, o.Payload)
}

func (o *GetProductsProductIDSyndicatedNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.SyndicatedProducthttp404Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
