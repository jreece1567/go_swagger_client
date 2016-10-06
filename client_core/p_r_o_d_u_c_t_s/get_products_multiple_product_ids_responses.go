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

// GetProductsMultipleProductIdsReader is a Reader for the GetProductsMultipleProductIds structure.
type GetProductsMultipleProductIdsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetProductsMultipleProductIdsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetProductsMultipleProductIdsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetProductsMultipleProductIdsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetProductsMultipleProductIdsOK creates a GetProductsMultipleProductIdsOK with default headers values
func NewGetProductsMultipleProductIdsOK() *GetProductsMultipleProductIdsOK {
	return &GetProductsMultipleProductIdsOK{}
}

/*GetProductsMultipleProductIdsOK handles this case with default header values.

Information about the requested products
*/
type GetProductsMultipleProductIdsOK struct {
	Payload *models_core.ProductListResponse
}

func (o *GetProductsMultipleProductIdsOK) Error() string {
	return fmt.Sprintf("[GET /products/multiple/{product_ids}][%d] getProductsMultipleProductIdsOK  %+v", 200, o.Payload)
}

func (o *GetProductsMultipleProductIdsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.ProductListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProductsMultipleProductIdsNotFound creates a GetProductsMultipleProductIdsNotFound with default headers values
func NewGetProductsMultipleProductIdsNotFound() *GetProductsMultipleProductIdsNotFound {
	return &GetProductsMultipleProductIdsNotFound{}
}

/*GetProductsMultipleProductIdsNotFound handles this case with default header values.

Products not found
*/
type GetProductsMultipleProductIdsNotFound struct {
	Payload *models_core.Producthttp404Response
}

func (o *GetProductsMultipleProductIdsNotFound) Error() string {
	return fmt.Sprintf("[GET /products/multiple/{product_ids}][%d] getProductsMultipleProductIdsNotFound  %+v", 404, o.Payload)
}

func (o *GetProductsMultipleProductIdsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.Producthttp404Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
