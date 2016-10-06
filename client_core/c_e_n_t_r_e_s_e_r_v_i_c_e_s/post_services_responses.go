package c_e_n_t_r_e_s_e_r_v_i_c_e_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"restclient/models_core"
)

// PostServicesReader is a Reader for the PostServices structure.
type PostServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PostServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostServicesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostServicesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPostServicesUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostServicesCreated creates a PostServicesCreated with default headers values
func NewPostServicesCreated() *PostServicesCreated {
	return &PostServicesCreated{}
}

/*PostServicesCreated handles this case with default header values.

Successfully created the service.
*/
type PostServicesCreated struct {
	/*URI of created service.
	 */
	Location string

	Payload *models_core.ServiceResponse
}

func (o *PostServicesCreated) Error() string {
	return fmt.Sprintf("[POST /services][%d] postServicesCreated  %+v", 201, o.Payload)
}

func (o *PostServicesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header location
	o.Location = response.GetHeader("location")

	o.Payload = new(models_core.ServiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostServicesUnauthorized creates a PostServicesUnauthorized with default headers values
func NewPostServicesUnauthorized() *PostServicesUnauthorized {
	return &PostServicesUnauthorized{}
}

/*PostServicesUnauthorized handles this case with default header values.

Unauthorized request.
*/
type PostServicesUnauthorized struct {
	Payload *models_core.Http401Response
}

func (o *PostServicesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /services][%d] postServicesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostServicesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.Http401Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostServicesUnprocessableEntity creates a PostServicesUnprocessableEntity with default headers values
func NewPostServicesUnprocessableEntity() *PostServicesUnprocessableEntity {
	return &PostServicesUnprocessableEntity{}
}

/*PostServicesUnprocessableEntity handles this case with default header values.

Service not created. Returns an array of error messages explaining the problems with the provided attributes.
*/
type PostServicesUnprocessableEntity struct {
	Payload *models_core.Service422Response
}

func (o *PostServicesUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /services][%d] postServicesUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostServicesUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.Service422Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}