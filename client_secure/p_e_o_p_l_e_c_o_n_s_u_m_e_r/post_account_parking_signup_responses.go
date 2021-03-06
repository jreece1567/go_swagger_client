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

// PostAccountParkingSignupReader is a Reader for the PostAccountParkingSignup structure.
type PostAccountParkingSignupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PostAccountParkingSignupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostAccountParkingSignupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostAccountParkingSignupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPostAccountParkingSignupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPostAccountParkingSignupUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostAccountParkingSignupOK creates a PostAccountParkingSignupOK with default headers values
func NewPostAccountParkingSignupOK() *PostAccountParkingSignupOK {
	return &PostAccountParkingSignupOK{}
}

/*PostAccountParkingSignupOK handles this case with default header values.

Successfully signed up a user to parking service.
*/
type PostAccountParkingSignupOK struct {
}

func (o *PostAccountParkingSignupOK) Error() string {
	return fmt.Sprintf("[POST /account/parking/signup][%d] postAccountParkingSignupOK ", 200)
}

func (o *PostAccountParkingSignupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAccountParkingSignupBadRequest creates a PostAccountParkingSignupBadRequest with default headers values
func NewPostAccountParkingSignupBadRequest() *PostAccountParkingSignupBadRequest {
	return &PostAccountParkingSignupBadRequest{}
}

/*PostAccountParkingSignupBadRequest handles this case with default header values.

Bad request
*/
type PostAccountParkingSignupBadRequest struct {
	Payload *models_secure.Person400Response
}

func (o *PostAccountParkingSignupBadRequest) Error() string {
	return fmt.Sprintf("[POST /account/parking/signup][%d] postAccountParkingSignupBadRequest  %+v", 400, o.Payload)
}

func (o *PostAccountParkingSignupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.Person400Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAccountParkingSignupUnauthorized creates a PostAccountParkingSignupUnauthorized with default headers values
func NewPostAccountParkingSignupUnauthorized() *PostAccountParkingSignupUnauthorized {
	return &PostAccountParkingSignupUnauthorized{}
}

/*PostAccountParkingSignupUnauthorized handles this case with default header values.

Unauthorized request.
*/
type PostAccountParkingSignupUnauthorized struct {
	Payload *models_secure.UnauthorizedResponse
}

func (o *PostAccountParkingSignupUnauthorized) Error() string {
	return fmt.Sprintf("[POST /account/parking/signup][%d] postAccountParkingSignupUnauthorized  %+v", 401, o.Payload)
}

func (o *PostAccountParkingSignupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.UnauthorizedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAccountParkingSignupUnprocessableEntity creates a PostAccountParkingSignupUnprocessableEntity with default headers values
func NewPostAccountParkingSignupUnprocessableEntity() *PostAccountParkingSignupUnprocessableEntity {
	return &PostAccountParkingSignupUnprocessableEntity{}
}

/*PostAccountParkingSignupUnprocessableEntity handles this case with default header values.

Account already registered on parking service.
*/
type PostAccountParkingSignupUnprocessableEntity struct {
	Payload *models_secure.ParkingSignup422Response
}

func (o *PostAccountParkingSignupUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /account/parking/signup][%d] postAccountParkingSignupUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostAccountParkingSignupUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.ParkingSignup422Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
