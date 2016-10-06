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

// GetAccountNewslettersReader is a Reader for the GetAccountNewsletters structure.
type GetAccountNewslettersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetAccountNewslettersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAccountNewslettersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetAccountNewslettersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetAccountNewslettersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAccountNewslettersOK creates a GetAccountNewslettersOK with default headers values
func NewGetAccountNewslettersOK() *GetAccountNewslettersOK {
	return &GetAccountNewslettersOK{}
}

/*GetAccountNewslettersOK handles this case with default header values.

Successfully retrieved newsletter details.
*/
type GetAccountNewslettersOK struct {
	Payload *models_secure.NewslettersSubscribe200Response
}

func (o *GetAccountNewslettersOK) Error() string {
	return fmt.Sprintf("[GET /account/newsletters][%d] getAccountNewslettersOK  %+v", 200, o.Payload)
}

func (o *GetAccountNewslettersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.NewslettersSubscribe200Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountNewslettersUnauthorized creates a GetAccountNewslettersUnauthorized with default headers values
func NewGetAccountNewslettersUnauthorized() *GetAccountNewslettersUnauthorized {
	return &GetAccountNewslettersUnauthorized{}
}

/*GetAccountNewslettersUnauthorized handles this case with default header values.

The account is full and the given token is not associated with the account
*/
type GetAccountNewslettersUnauthorized struct {
	Payload *models_secure.GetNewsletters401Response
}

func (o *GetAccountNewslettersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /account/newsletters][%d] getAccountNewslettersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAccountNewslettersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.GetNewsletters401Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountNewslettersNotFound creates a GetAccountNewslettersNotFound with default headers values
func NewGetAccountNewslettersNotFound() *GetAccountNewslettersNotFound {
	return &GetAccountNewslettersNotFound{}
}

/*GetAccountNewslettersNotFound handles this case with default header values.

The account is not found
*/
type GetAccountNewslettersNotFound struct {
	Payload *models_secure.GetNewslettershttp404Response
}

func (o *GetAccountNewslettersNotFound) Error() string {
	return fmt.Sprintf("[GET /account/newsletters][%d] getAccountNewslettersNotFound  %+v", 404, o.Payload)
}

func (o *GetAccountNewslettersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.GetNewslettershttp404Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}