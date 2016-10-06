package p_e_o_p_l_e_m_a_n_a_g_e_m_e_n_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"restclient/models_secure"
)

// GetPeoplePersonIDFoodOrdersFoodOrderIDReader is a Reader for the GetPeoplePersonIDFoodOrdersFoodOrderID structure.
type GetPeoplePersonIDFoodOrdersFoodOrderIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetPeoplePersonIDFoodOrdersFoodOrderIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPeoplePersonIDFoodOrdersFoodOrderIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetPeoplePersonIDFoodOrdersFoodOrderIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetPeoplePersonIDFoodOrdersFoodOrderIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPeoplePersonIDFoodOrdersFoodOrderIDOK creates a GetPeoplePersonIDFoodOrdersFoodOrderIDOK with default headers values
func NewGetPeoplePersonIDFoodOrdersFoodOrderIDOK() *GetPeoplePersonIDFoodOrdersFoodOrderIDOK {
	return &GetPeoplePersonIDFoodOrdersFoodOrderIDOK{}
}

/*GetPeoplePersonIDFoodOrdersFoodOrderIDOK handles this case with default header values.

Successfully retrieved the order.
*/
type GetPeoplePersonIDFoodOrdersFoodOrderIDOK struct {
	Payload *models_secure.FoodOrderResponse
}

func (o *GetPeoplePersonIDFoodOrdersFoodOrderIDOK) Error() string {
	return fmt.Sprintf("[GET /people/{person_id}/food/orders/{food_order_id}][%d] getPeoplePersonIdFoodOrdersFoodOrderIdOK  %+v", 200, o.Payload)
}

func (o *GetPeoplePersonIDFoodOrdersFoodOrderIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.FoodOrderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPeoplePersonIDFoodOrdersFoodOrderIDUnauthorized creates a GetPeoplePersonIDFoodOrdersFoodOrderIDUnauthorized with default headers values
func NewGetPeoplePersonIDFoodOrdersFoodOrderIDUnauthorized() *GetPeoplePersonIDFoodOrdersFoodOrderIDUnauthorized {
	return &GetPeoplePersonIDFoodOrdersFoodOrderIDUnauthorized{}
}

/*GetPeoplePersonIDFoodOrdersFoodOrderIDUnauthorized handles this case with default header values.

Unauthorized request.
*/
type GetPeoplePersonIDFoodOrdersFoodOrderIDUnauthorized struct {
	Payload *models_secure.UnauthorizedResponse
}

func (o *GetPeoplePersonIDFoodOrdersFoodOrderIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /people/{person_id}/food/orders/{food_order_id}][%d] getPeoplePersonIdFoodOrdersFoodOrderIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPeoplePersonIDFoodOrdersFoodOrderIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.UnauthorizedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPeoplePersonIDFoodOrdersFoodOrderIDNotFound creates a GetPeoplePersonIDFoodOrdersFoodOrderIDNotFound with default headers values
func NewGetPeoplePersonIDFoodOrdersFoodOrderIDNotFound() *GetPeoplePersonIDFoodOrdersFoodOrderIDNotFound {
	return &GetPeoplePersonIDFoodOrdersFoodOrderIDNotFound{}
}

/*GetPeoplePersonIDFoodOrdersFoodOrderIDNotFound handles this case with default header values.

Order not found with the given id
*/
type GetPeoplePersonIDFoodOrdersFoodOrderIDNotFound struct {
	Payload *models_secure.Http404Response
}

func (o *GetPeoplePersonIDFoodOrdersFoodOrderIDNotFound) Error() string {
	return fmt.Sprintf("[GET /people/{person_id}/food/orders/{food_order_id}][%d] getPeoplePersonIdFoodOrdersFoodOrderIdNotFound  %+v", 404, o.Payload)
}

func (o *GetPeoplePersonIDFoodOrdersFoodOrderIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.Http404Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}