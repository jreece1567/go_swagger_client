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

// PostPaymentsTransactionsAuthorizeReader is a Reader for the PostPaymentsTransactionsAuthorize structure.
type PostPaymentsTransactionsAuthorizeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PostPaymentsTransactionsAuthorizeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostPaymentsTransactionsAuthorizeCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 422:
		result := NewPostPaymentsTransactionsAuthorizeUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPaymentsTransactionsAuthorizeCreated creates a PostPaymentsTransactionsAuthorizeCreated with default headers values
func NewPostPaymentsTransactionsAuthorizeCreated() *PostPaymentsTransactionsAuthorizeCreated {
	return &PostPaymentsTransactionsAuthorizeCreated{}
}

/*PostPaymentsTransactionsAuthorizeCreated handles this case with default header values.

Successfully created the transaction.
*/
type PostPaymentsTransactionsAuthorizeCreated struct {
	/*URI of created transaction.
	 */
	Location string

	Payload *models_secure.TransactionResponse
}

func (o *PostPaymentsTransactionsAuthorizeCreated) Error() string {
	return fmt.Sprintf("[POST /payments/transactions/authorize][%d] postPaymentsTransactionsAuthorizeCreated  %+v", 201, o.Payload)
}

func (o *PostPaymentsTransactionsAuthorizeCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header location
	o.Location = response.GetHeader("location")

	o.Payload = new(models_secure.TransactionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPaymentsTransactionsAuthorizeUnprocessableEntity creates a PostPaymentsTransactionsAuthorizeUnprocessableEntity with default headers values
func NewPostPaymentsTransactionsAuthorizeUnprocessableEntity() *PostPaymentsTransactionsAuthorizeUnprocessableEntity {
	return &PostPaymentsTransactionsAuthorizeUnprocessableEntity{}
}

/*PostPaymentsTransactionsAuthorizeUnprocessableEntity handles this case with default header values.

Transaction not created. Returns an array of error messages explaining the problems with the provided attributes.
*/
type PostPaymentsTransactionsAuthorizeUnprocessableEntity struct {
	Payload *models_secure.Transaction422Response
}

func (o *PostPaymentsTransactionsAuthorizeUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /payments/transactions/authorize][%d] postPaymentsTransactionsAuthorizeUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostPaymentsTransactionsAuthorizeUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.Transaction422Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
