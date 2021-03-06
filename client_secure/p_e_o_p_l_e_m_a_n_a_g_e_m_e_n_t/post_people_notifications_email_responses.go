package p_e_o_p_l_e_m_a_n_a_g_e_m_e_n_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"restclient/models_secure"
)

// PostPeopleNotificationsEmailReader is a Reader for the PostPeopleNotificationsEmail structure.
type PostPeopleNotificationsEmailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PostPeopleNotificationsEmailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPostPeopleNotificationsEmailNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostPeopleNotificationsEmailBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPostPeopleNotificationsEmailUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPeopleNotificationsEmailNoContent creates a PostPeopleNotificationsEmailNoContent with default headers values
func NewPostPeopleNotificationsEmailNoContent() *PostPeopleNotificationsEmailNoContent {
	return &PostPeopleNotificationsEmailNoContent{}
}

/*PostPeopleNotificationsEmailNoContent handles this case with default header values.

Message published succesfully
*/
type PostPeopleNotificationsEmailNoContent struct {
}

func (o *PostPeopleNotificationsEmailNoContent) Error() string {
	return fmt.Sprintf("[POST /people/notifications/email][%d] postPeopleNotificationsEmailNoContent ", 204)
}

func (o *PostPeopleNotificationsEmailNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostPeopleNotificationsEmailBadRequest creates a PostPeopleNotificationsEmailBadRequest with default headers values
func NewPostPeopleNotificationsEmailBadRequest() *PostPeopleNotificationsEmailBadRequest {
	return &PostPeopleNotificationsEmailBadRequest{}
}

/*PostPeopleNotificationsEmailBadRequest handles this case with default header values.

Bad request
*/
type PostPeopleNotificationsEmailBadRequest struct {
	Payload *models_secure.Person400Response
}

func (o *PostPeopleNotificationsEmailBadRequest) Error() string {
	return fmt.Sprintf("[POST /people/notifications/email][%d] postPeopleNotificationsEmailBadRequest  %+v", 400, o.Payload)
}

func (o *PostPeopleNotificationsEmailBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.Person400Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostPeopleNotificationsEmailUnprocessableEntity creates a PostPeopleNotificationsEmailUnprocessableEntity with default headers values
func NewPostPeopleNotificationsEmailUnprocessableEntity() *PostPeopleNotificationsEmailUnprocessableEntity {
	return &PostPeopleNotificationsEmailUnprocessableEntity{}
}

/*PostPeopleNotificationsEmailUnprocessableEntity handles this case with default header values.

Centre or Person identitifer does not exist
*/
type PostPeopleNotificationsEmailUnprocessableEntity struct {
	Payload *models_secure.Notifications422Response
}

func (o *PostPeopleNotificationsEmailUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /people/notifications/email][%d] postPeopleNotificationsEmailUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostPeopleNotificationsEmailUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.Notifications422Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
