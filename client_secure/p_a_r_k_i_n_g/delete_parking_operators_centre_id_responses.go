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

// DeleteParkingOperatorsCentreIDReader is a Reader for the DeleteParkingOperatorsCentreID structure.
type DeleteParkingOperatorsCentreIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeleteParkingOperatorsCentreIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteParkingOperatorsCentreIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteParkingOperatorsCentreIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteParkingOperatorsCentreIDNoContent creates a DeleteParkingOperatorsCentreIDNoContent with default headers values
func NewDeleteParkingOperatorsCentreIDNoContent() *DeleteParkingOperatorsCentreIDNoContent {
	return &DeleteParkingOperatorsCentreIDNoContent{}
}

/*DeleteParkingOperatorsCentreIDNoContent handles this case with default header values.

Successfully deleted the Car Park Business Operator.
*/
type DeleteParkingOperatorsCentreIDNoContent struct {
}

func (o *DeleteParkingOperatorsCentreIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /parking/operators/{centre_id}][%d] deleteParkingOperatorsCentreIdNoContent ", 204)
}

func (o *DeleteParkingOperatorsCentreIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteParkingOperatorsCentreIDNotFound creates a DeleteParkingOperatorsCentreIDNotFound with default headers values
func NewDeleteParkingOperatorsCentreIDNotFound() *DeleteParkingOperatorsCentreIDNotFound {
	return &DeleteParkingOperatorsCentreIDNotFound{}
}

/*DeleteParkingOperatorsCentreIDNotFound handles this case with default header values.

Car Park Business Operator not found.
*/
type DeleteParkingOperatorsCentreIDNotFound struct {
	Payload *models_secure.Http404Response
}

func (o *DeleteParkingOperatorsCentreIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /parking/operators/{centre_id}][%d] deleteParkingOperatorsCentreIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteParkingOperatorsCentreIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.Http404Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
