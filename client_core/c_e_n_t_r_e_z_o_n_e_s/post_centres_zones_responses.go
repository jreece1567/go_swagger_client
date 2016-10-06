package c_e_n_t_r_e_z_o_n_e_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"restclient/models_core"
)

// PostCentresZonesReader is a Reader for the PostCentresZones structure.
type PostCentresZonesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PostCentresZonesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostCentresZonesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostCentresZonesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPostCentresZonesUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostCentresZonesCreated creates a PostCentresZonesCreated with default headers values
func NewPostCentresZonesCreated() *PostCentresZonesCreated {
	return &PostCentresZonesCreated{}
}

/*PostCentresZonesCreated handles this case with default header values.

Successfully created the centre zone.
*/
type PostCentresZonesCreated struct {
	/*URI of created zone.
	 */
	Location string

	Payload *models_core.ZoneResponse
}

func (o *PostCentresZonesCreated) Error() string {
	return fmt.Sprintf("[POST /centres/zones][%d] postCentresZonesCreated  %+v", 201, o.Payload)
}

func (o *PostCentresZonesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header location
	o.Location = response.GetHeader("location")

	o.Payload = new(models_core.ZoneResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCentresZonesUnauthorized creates a PostCentresZonesUnauthorized with default headers values
func NewPostCentresZonesUnauthorized() *PostCentresZonesUnauthorized {
	return &PostCentresZonesUnauthorized{}
}

/*PostCentresZonesUnauthorized handles this case with default header values.

Unauthorized request.
*/
type PostCentresZonesUnauthorized struct {
	Payload *models_core.Http401Response
}

func (o *PostCentresZonesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /centres/zones][%d] postCentresZonesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostCentresZonesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.Http401Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCentresZonesUnprocessableEntity creates a PostCentresZonesUnprocessableEntity with default headers values
func NewPostCentresZonesUnprocessableEntity() *PostCentresZonesUnprocessableEntity {
	return &PostCentresZonesUnprocessableEntity{}
}

/*PostCentresZonesUnprocessableEntity handles this case with default header values.

Centre parking not created. Returns an array of error messages explaining the problems with the provided attributes.
*/
type PostCentresZonesUnprocessableEntity struct {
	Payload *models_core.Zone422Response
}

func (o *PostCentresZonesUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /centres/zones][%d] postCentresZonesUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostCentresZonesUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.Zone422Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}