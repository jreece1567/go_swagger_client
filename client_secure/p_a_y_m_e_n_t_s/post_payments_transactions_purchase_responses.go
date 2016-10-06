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

// PostPaymentsTransactionsPurchaseReader is a Reader for the PostPaymentsTransactionsPurchase structure.
type PostPaymentsTransactionsPurchaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PostPaymentsTransactionsPurchaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostPaymentsTransactionsPurchaseCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 422:
		result := NewPostPaymentsTransactionsPurchaseUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPaymentsTransactionsPurchaseCreated creates a PostPaymentsTransactionsPurchaseCreated with default headers values
func NewPostPaymentsTransactionsPurchaseCreated() *PostPaymentsTransactionsPurchaseCreated {
	return &PostPaymentsTransactionsPurchaseCreated{}
}

/*PostPaymentsTransactionsPurchaseCreated handles this case with default header values.

Successfully created the purchase transaction.
*/
type PostPaymentsTransactionsPurchaseCreated struct {
	/*URI of created transaction.
	 */
	Location string

	Payload *models_secure.TransactionResponse
}

func (o *PostPaymentsTransactionsPurchaseCreated) Error() string {
	return fmt.Sprintf("[POST /payments/transactions/purchase][%d] postPaymentsTransactionsPurchaseCreated  %+v", 201, o.Payload)
}

func (o *PostPaymentsTransactionsPurchaseCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header location
	o.Location = response.GetHeader("location")

	o.Payload = new(models_secure.TransactionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsTransactionsPurchaseUnprocessableEntity creates a PostPaymentsTransactionsPurchaseUnprocessableEntity with default headers values
func NewPostPaymentsTransactionsPurchaseUnprocessableEntity() *PostPaymentsTransactionsPurchaseUnprocessableEntity {
	return &PostPaymentsTransactionsPurchaseUnprocessableEntity{}
}

/*PostPaymentsTransactionsPurchaseUnprocessableEntity handles this case with default header values.

Transaction not created. Returns an array of error messages explaining the problems with the provided attributes.
*/
type PostPaymentsTransactionsPurchaseUnprocessableEntity struct {
	Payload *models_secure.Transaction422Response
}

func (o *PostPaymentsTransactionsPurchaseUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /payments/transactions/purchase][%d] postPaymentsTransactionsPurchaseUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostPaymentsTransactionsPurchaseUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.Transaction422Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
