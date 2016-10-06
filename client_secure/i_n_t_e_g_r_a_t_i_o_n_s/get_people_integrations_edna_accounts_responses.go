package i_n_t_e_g_r_a_t_i_o_n_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"restclient/models_secure"
)

// GetPeopleIntegrationsEdnaAccountsReader is a Reader for the GetPeopleIntegrationsEdnaAccounts structure.
type GetPeopleIntegrationsEdnaAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetPeopleIntegrationsEdnaAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPeopleIntegrationsEdnaAccountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetPeopleIntegrationsEdnaAccountsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewGetPeopleIntegrationsEdnaAccountsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPeopleIntegrationsEdnaAccountsOK creates a GetPeopleIntegrationsEdnaAccountsOK with default headers values
func NewGetPeopleIntegrationsEdnaAccountsOK() *GetPeopleIntegrationsEdnaAccountsOK {
	return &GetPeopleIntegrationsEdnaAccountsOK{}
}

/*GetPeopleIntegrationsEdnaAccountsOK handles this case with default header values.

Successfully retrieved the list of accounts.
*/
type GetPeopleIntegrationsEdnaAccountsOK struct {
	Payload *models_secure.GetEdna200Response
}

func (o *GetPeopleIntegrationsEdnaAccountsOK) Error() string {
	return fmt.Sprintf("[GET /people/integrations/edna/accounts][%d] getPeopleIntegrationsEdnaAccountsOK  %+v", 200, o.Payload)
}

func (o *GetPeopleIntegrationsEdnaAccountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.GetEdna200Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPeopleIntegrationsEdnaAccountsUnauthorized creates a GetPeopleIntegrationsEdnaAccountsUnauthorized with default headers values
func NewGetPeopleIntegrationsEdnaAccountsUnauthorized() *GetPeopleIntegrationsEdnaAccountsUnauthorized {
	return &GetPeopleIntegrationsEdnaAccountsUnauthorized{}
}

/*GetPeopleIntegrationsEdnaAccountsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetPeopleIntegrationsEdnaAccountsUnauthorized struct {
	Payload *models_secure.UnauthorizedResponse
}

func (o *GetPeopleIntegrationsEdnaAccountsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /people/integrations/edna/accounts][%d] getPeopleIntegrationsEdnaAccountsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPeopleIntegrationsEdnaAccountsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.UnauthorizedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPeopleIntegrationsEdnaAccountsUnprocessableEntity creates a GetPeopleIntegrationsEdnaAccountsUnprocessableEntity with default headers values
func NewGetPeopleIntegrationsEdnaAccountsUnprocessableEntity() *GetPeopleIntegrationsEdnaAccountsUnprocessableEntity {
	return &GetPeopleIntegrationsEdnaAccountsUnprocessableEntity{}
}

/*GetPeopleIntegrationsEdnaAccountsUnprocessableEntity handles this case with default header values.

Owner identifier or both next_page_id and datetime are missing
*/
type GetPeopleIntegrationsEdnaAccountsUnprocessableEntity struct {
	Payload *models_secure.GetEdna422Response
}

func (o *GetPeopleIntegrationsEdnaAccountsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /people/integrations/edna/accounts][%d] getPeopleIntegrationsEdnaAccountsUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetPeopleIntegrationsEdnaAccountsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.GetEdna422Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
