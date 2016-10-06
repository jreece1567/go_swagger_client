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

// GetPaymentsTransactionsTransactionIDReader is a Reader for the GetPaymentsTransactionsTransactionID structure.
type GetPaymentsTransactionsTransactionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetPaymentsTransactionsTransactionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentsTransactionsTransactionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetPaymentsTransactionsTransactionIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentsTransactionsTransactionIDOK creates a GetPaymentsTransactionsTransactionIDOK with default headers values
func NewGetPaymentsTransactionsTransactionIDOK() *GetPaymentsTransactionsTransactionIDOK {
	return &GetPaymentsTransactionsTransactionIDOK{}
}

/*GetPaymentsTransactionsTransactionIDOK handles this case with default header values.

Successfully retrieved the individual transaction.
*/
type GetPaymentsTransactionsTransactionIDOK struct {
	Payload *models_secure.TransactionResponse
}

func (o *GetPaymentsTransactionsTransactionIDOK) Error() string {
	return fmt.Sprintf("[GET /payments/transactions/{transaction_id}][%d] getPaymentsTransactionsTransactionIdOK  %+v", 200, o.Payload)
}

func (o *GetPaymentsTransactionsTransactionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.TransactionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentsTransactionsTransactionIDNotFound creates a GetPaymentsTransactionsTransactionIDNotFound with default headers values
func NewGetPaymentsTransactionsTransactionIDNotFound() *GetPaymentsTransactionsTransactionIDNotFound {
	return &GetPaymentsTransactionsTransactionIDNotFound{}
}

/*GetPaymentsTransactionsTransactionIDNotFound handles this case with default header values.

Transaction not found.
*/
type GetPaymentsTransactionsTransactionIDNotFound struct {
	Payload *models_secure.Http404Response
}

func (o *GetPaymentsTransactionsTransactionIDNotFound) Error() string {
	return fmt.Sprintf("[GET /payments/transactions/{transaction_id}][%d] getPaymentsTransactionsTransactionIdNotFound  %+v", 404, o.Payload)
}

func (o *GetPaymentsTransactionsTransactionIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.Http404Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}