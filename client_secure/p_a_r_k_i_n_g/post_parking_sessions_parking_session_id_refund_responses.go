package p_a_r_k_i_n_g

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"restclient/models_secure"
)

// PostParkingSessionsParkingSessionIDRefundReader is a Reader for the PostParkingSessionsParkingSessionIDRefund structure.
type PostParkingSessionsParkingSessionIDRefundReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PostParkingSessionsParkingSessionIDRefundReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostParkingSessionsParkingSessionIDRefundOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPostParkingSessionsParkingSessionIDRefundNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPostParkingSessionsParkingSessionIDRefundUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostParkingSessionsParkingSessionIDRefundOK creates a PostParkingSessionsParkingSessionIDRefundOK with default headers values
func NewPostParkingSessionsParkingSessionIDRefundOK() *PostParkingSessionsParkingSessionIDRefundOK {
	return &PostParkingSessionsParkingSessionIDRefundOK{}
}

/*PostParkingSessionsParkingSessionIDRefundOK handles this case with default header values.

Successfully refunded the charge.
*/
type PostParkingSessionsParkingSessionIDRefundOK struct {
	Payload *models_secure.ParkingSessionTransactionResponse
}

func (o *PostParkingSessionsParkingSessionIDRefundOK) Error() string {
	return fmt.Sprintf("[POST /parking/sessions/{parking_session_id}/refund][%d] postParkingSessionsParkingSessionIdRefundOK  %+v", 200, o.Payload)
}

func (o *PostParkingSessionsParkingSessionIDRefundOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.ParkingSessionTransactionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostParkingSessionsParkingSessionIDRefundNotFound creates a PostParkingSessionsParkingSessionIDRefundNotFound with default headers values
func NewPostParkingSessionsParkingSessionIDRefundNotFound() *PostParkingSessionsParkingSessionIDRefundNotFound {
	return &PostParkingSessionsParkingSessionIDRefundNotFound{}
}

/*PostParkingSessionsParkingSessionIDRefundNotFound handles this case with default header values.

Parking session not found
*/
type PostParkingSessionsParkingSessionIDRefundNotFound struct {
	Payload *models_secure.Http404Response
}

func (o *PostParkingSessionsParkingSessionIDRefundNotFound) Error() string {
	return fmt.Sprintf("[POST /parking/sessions/{parking_session_id}/refund][%d] postParkingSessionsParkingSessionIdRefundNotFound  %+v", 404, o.Payload)
}

func (o *PostParkingSessionsParkingSessionIDRefundNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.Http404Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostParkingSessionsParkingSessionIDRefundUnprocessableEntity creates a PostParkingSessionsParkingSessionIDRefundUnprocessableEntity with default headers values
func NewPostParkingSessionsParkingSessionIDRefundUnprocessableEntity() *PostParkingSessionsParkingSessionIDRefundUnprocessableEntity {
	return &PostParkingSessionsParkingSessionIDRefundUnprocessableEntity{}
}

/*PostParkingSessionsParkingSessionIDRefundUnprocessableEntity handles this case with default header values.

Parking session not eligible for a refund.
*/
type PostParkingSessionsParkingSessionIDRefundUnprocessableEntity struct {
	Payload *models_secure.ParkingSessionTransaction422Response
}

func (o *PostParkingSessionsParkingSessionIDRefundUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /parking/sessions/{parking_session_id}/refund][%d] postParkingSessionsParkingSessionIdRefundUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostParkingSessionsParkingSessionIDRefundUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.ParkingSessionTransaction422Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}