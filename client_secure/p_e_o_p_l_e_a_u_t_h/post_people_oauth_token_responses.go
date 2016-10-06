package p_e_o_p_l_e_a_u_t_h

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"restclient/models_secure"
)

// PostPeopleOauthTokenReader is a Reader for the PostPeopleOauthToken structure.
type PostPeopleOauthTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PostPeopleOauthTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostPeopleOauthTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostPeopleOauthTokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPostPeopleOauthTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 423:
		result := NewPostPeopleOauthTokenLocked()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPeopleOauthTokenOK creates a PostPeopleOauthTokenOK with default headers values
func NewPostPeopleOauthTokenOK() *PostPeopleOauthTokenOK {
	return &PostPeopleOauthTokenOK{}
}

/*PostPeopleOauthTokenOK handles this case with default header values.

Success
*/
type PostPeopleOauthTokenOK struct {
	Payload *models_secure.AccountToken
}

func (o *PostPeopleOauthTokenOK) Error() string {
	return fmt.Sprintf("[POST /people/oauth/token][%d] postPeopleOauthTokenOK  %+v", 200, o.Payload)
}

func (o *PostPeopleOauthTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.AccountToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPeopleOauthTokenBadRequest creates a PostPeopleOauthTokenBadRequest with default headers values
func NewPostPeopleOauthTokenBadRequest() *PostPeopleOauthTokenBadRequest {
	return &PostPeopleOauthTokenBadRequest{}
}

/*PostPeopleOauthTokenBadRequest handles this case with default header values.

Bad request
*/
type PostPeopleOauthTokenBadRequest struct {
	Payload *models_secure.Person400Response
}

func (o *PostPeopleOauthTokenBadRequest) Error() string {
	return fmt.Sprintf("[POST /people/oauth/token][%d] postPeopleOauthTokenBadRequest  %+v", 400, o.Payload)
}

func (o *PostPeopleOauthTokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.Person400Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPeopleOauthTokenUnauthorized creates a PostPeopleOauthTokenUnauthorized with default headers values
func NewPostPeopleOauthTokenUnauthorized() *PostPeopleOauthTokenUnauthorized {
	return &PostPeopleOauthTokenUnauthorized{}
}

/*PostPeopleOauthTokenUnauthorized handles this case with default header values.

Unauthorized
*/
type PostPeopleOauthTokenUnauthorized struct {
	Payload *models_secure.UnauthorizedResponse
}

func (o *PostPeopleOauthTokenUnauthorized) Error() string {
	return fmt.Sprintf("[POST /people/oauth/token][%d] postPeopleOauthTokenUnauthorized  %+v", 401, o.Payload)
}

func (o *PostPeopleOauthTokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.UnauthorizedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPeopleOauthTokenLocked creates a PostPeopleOauthTokenLocked with default headers values
func NewPostPeopleOauthTokenLocked() *PostPeopleOauthTokenLocked {
	return &PostPeopleOauthTokenLocked{}
}

/*PostPeopleOauthTokenLocked handles this case with default header values.

Locked account
*/
type PostPeopleOauthTokenLocked struct {
	Payload *models_secure.LockedResponse
}

func (o *PostPeopleOauthTokenLocked) Error() string {
	return fmt.Sprintf("[POST /people/oauth/token][%d] postPeopleOauthTokenLocked  %+v", 423, o.Payload)
}

func (o *PostPeopleOauthTokenLocked) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.LockedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}