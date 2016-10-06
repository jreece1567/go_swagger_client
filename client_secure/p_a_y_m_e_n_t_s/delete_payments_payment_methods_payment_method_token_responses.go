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

// DeletePaymentsPaymentMethodsPaymentMethodTokenReader is a Reader for the DeletePaymentsPaymentMethodsPaymentMethodToken structure.
type DeletePaymentsPaymentMethodsPaymentMethodTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeletePaymentsPaymentMethodsPaymentMethodTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeletePaymentsPaymentMethodsPaymentMethodTokenNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeletePaymentsPaymentMethodsPaymentMethodTokenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeletePaymentsPaymentMethodsPaymentMethodTokenNoContent creates a DeletePaymentsPaymentMethodsPaymentMethodTokenNoContent with default headers values
func NewDeletePaymentsPaymentMethodsPaymentMethodTokenNoContent() *DeletePaymentsPaymentMethodsPaymentMethodTokenNoContent {
	return &DeletePaymentsPaymentMethodsPaymentMethodTokenNoContent{}
}

/*DeletePaymentsPaymentMethodsPaymentMethodTokenNoContent handles this case with default header values.

Successfully redacted the payment method.
*/
type DeletePaymentsPaymentMethodsPaymentMethodTokenNoContent struct {
}

func (o *DeletePaymentsPaymentMethodsPaymentMethodTokenNoContent) Error() string {
	return fmt.Sprintf("[DELETE /payments/payment_methods/{payment_method_token}][%d] deletePaymentsPaymentMethodsPaymentMethodTokenNoContent ", 204)
}

func (o *DeletePaymentsPaymentMethodsPaymentMethodTokenNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePaymentsPaymentMethodsPaymentMethodTokenNotFound creates a DeletePaymentsPaymentMethodsPaymentMethodTokenNotFound with default headers values
func NewDeletePaymentsPaymentMethodsPaymentMethodTokenNotFound() *DeletePaymentsPaymentMethodsPaymentMethodTokenNotFound {
	return &DeletePaymentsPaymentMethodsPaymentMethodTokenNotFound{}
}

/*DeletePaymentsPaymentMethodsPaymentMethodTokenNotFound handles this case with default header values.

Payment method not found.
*/
type DeletePaymentsPaymentMethodsPaymentMethodTokenNotFound struct {
	Payload *models_secure.Http404Response
}

func (o *DeletePaymentsPaymentMethodsPaymentMethodTokenNotFound) Error() string {
	return fmt.Sprintf("[DELETE /payments/payment_methods/{payment_method_token}][%d] deletePaymentsPaymentMethodsPaymentMethodTokenNotFound  %+v", 404, o.Payload)
}

func (o *DeletePaymentsPaymentMethodsPaymentMethodTokenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.Http404Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
