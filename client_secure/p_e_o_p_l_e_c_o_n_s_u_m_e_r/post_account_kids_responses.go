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

// PostAccountKidsReader is a Reader for the PostAccountKids structure.
type PostAccountKidsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PostAccountKidsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostAccountKidsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostAccountKidsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostAccountKidsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPostAccountKidsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostAccountKidsOK creates a PostAccountKidsOK with default headers values
func NewPostAccountKidsOK() *PostAccountKidsOK {
	return &PostAccountKidsOK{}
}

/*PostAccountKidsOK handles this case with default header values.

Successfully added the list of kids into account.
*/
type PostAccountKidsOK struct {
}

func (o *PostAccountKidsOK) Error() string {
	return fmt.Sprintf("[POST /account/kids][%d] postAccountKidsOK ", 200)
}

func (o *PostAccountKidsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAccountKidsBadRequest creates a PostAccountKidsBadRequest with default headers values
func NewPostAccountKidsBadRequest() *PostAccountKidsBadRequest {
	return &PostAccountKidsBadRequest{}
}

/*PostAccountKidsBadRequest handles this case with default header values.

Bad request
*/
type PostAccountKidsBadRequest struct {
	Payload *models_secure.Person400Response
}

func (o *PostAccountKidsBadRequest) Error() string {
	return fmt.Sprintf("[POST /account/kids][%d] postAccountKidsBadRequest  %+v", 400, o.Payload)
}

func (o *PostAccountKidsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.Person400Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAccountKidsNotFound creates a PostAccountKidsNotFound with default headers values
func NewPostAccountKidsNotFound() *PostAccountKidsNotFound {
	return &PostAccountKidsNotFound{}
}

/*PostAccountKidsNotFound handles this case with default header values.

Account not found
*/
type PostAccountKidsNotFound struct {
	Payload *models_secure.Http404Response
}

func (o *PostAccountKidsNotFound) Error() string {
	return fmt.Sprintf("[POST /account/kids][%d] postAccountKidsNotFound  %+v", 404, o.Payload)
}

func (o *PostAccountKidsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.Http404Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAccountKidsUnprocessableEntity creates a PostAccountKidsUnprocessableEntity with default headers values
func NewPostAccountKidsUnprocessableEntity() *PostAccountKidsUnprocessableEntity {
	return &PostAccountKidsUnprocessableEntity{}
}

/*PostAccountKidsUnprocessableEntity handles this case with default header values.

Missing or invalid kid data.
*/
type PostAccountKidsUnprocessableEntity struct {
	Payload *models_secure.AddKids422Response
}

func (o *PostAccountKidsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /account/kids][%d] postAccountKidsUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostAccountKidsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.AddKids422Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
