package p_e_o_p_l_e_c_o_n_s_u_m_e_r

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"restclient/models_secure"
)

// GetAccountCreditCardsReader is a Reader for the GetAccountCreditCards structure.
type GetAccountCreditCardsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetAccountCreditCardsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAccountCreditCardsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetAccountCreditCardsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAccountCreditCardsOK creates a GetAccountCreditCardsOK with default headers values
func NewGetAccountCreditCardsOK() *GetAccountCreditCardsOK {
	return &GetAccountCreditCardsOK{}
}

/*GetAccountCreditCardsOK handles this case with default header values.

Successfully retrieved the list of user credit cards.
*/
type GetAccountCreditCardsOK struct {
	Payload *models_secure.CreditCardsListResponse
}

func (o *GetAccountCreditCardsOK) Error() string {
	return fmt.Sprintf("[GET /account/credit_cards][%d] getAccountCreditCardsOK  %+v", 200, o.Payload)
}

func (o *GetAccountCreditCardsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.CreditCardsListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountCreditCardsUnauthorized creates a GetAccountCreditCardsUnauthorized with default headers values
func NewGetAccountCreditCardsUnauthorized() *GetAccountCreditCardsUnauthorized {
	return &GetAccountCreditCardsUnauthorized{}
}

/*GetAccountCreditCardsUnauthorized handles this case with default header values.

Unauthorized request.
*/
type GetAccountCreditCardsUnauthorized struct {
	Payload *models_secure.UnauthorizedResponse
}

func (o *GetAccountCreditCardsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /account/credit_cards][%d] getAccountCreditCardsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAccountCreditCardsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.UnauthorizedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
