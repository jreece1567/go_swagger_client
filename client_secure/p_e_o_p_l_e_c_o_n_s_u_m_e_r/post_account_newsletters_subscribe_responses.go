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

// PostAccountNewslettersSubscribeReader is a Reader for the PostAccountNewslettersSubscribe structure.
type PostAccountNewslettersSubscribeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PostAccountNewslettersSubscribeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostAccountNewslettersSubscribeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostAccountNewslettersSubscribeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostAccountNewslettersSubscribeOK creates a PostAccountNewslettersSubscribeOK with default headers values
func NewPostAccountNewslettersSubscribeOK() *PostAccountNewslettersSubscribeOK {
	return &PostAccountNewslettersSubscribeOK{}
}

/*PostAccountNewslettersSubscribeOK handles this case with default header values.

Successfully subscribed account to newsletters.
*/
type PostAccountNewslettersSubscribeOK struct {
	Payload *models_secure.NewslettersSubscribe200Response
}

func (o *PostAccountNewslettersSubscribeOK) Error() string {
	return fmt.Sprintf("[POST /account/newsletters/subscribe][%d] postAccountNewslettersSubscribeOK  %+v", 200, o.Payload)
}

func (o *PostAccountNewslettersSubscribeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.NewslettersSubscribe200Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAccountNewslettersSubscribeBadRequest creates a PostAccountNewslettersSubscribeBadRequest with default headers values
func NewPostAccountNewslettersSubscribeBadRequest() *PostAccountNewslettersSubscribeBadRequest {
	return &PostAccountNewslettersSubscribeBadRequest{}
}

/*PostAccountNewslettersSubscribeBadRequest handles this case with default header values.

Bad request
*/
type PostAccountNewslettersSubscribeBadRequest struct {
	Payload *models_secure.Person400Response
}

func (o *PostAccountNewslettersSubscribeBadRequest) Error() string {
	return fmt.Sprintf("[POST /account/newsletters/subscribe][%d] postAccountNewslettersSubscribeBadRequest  %+v", 400, o.Payload)
}

func (o *PostAccountNewslettersSubscribeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.Person400Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}