package d_e_a_l_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"restclient/models_core"
)

// GetDealsCommentsCommentIDReader is a Reader for the GetDealsCommentsCommentID structure.
type GetDealsCommentsCommentIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetDealsCommentsCommentIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDealsCommentsCommentIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetDealsCommentsCommentIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDealsCommentsCommentIDOK creates a GetDealsCommentsCommentIDOK with default headers values
func NewGetDealsCommentsCommentIDOK() *GetDealsCommentsCommentIDOK {
	return &GetDealsCommentsCommentIDOK{}
}

/*GetDealsCommentsCommentIDOK handles this case with default header values.

Successfully retrieved the individual deal comment.
*/
type GetDealsCommentsCommentIDOK struct {
	Payload *models_core.CommentResponse
}

func (o *GetDealsCommentsCommentIDOK) Error() string {
	return fmt.Sprintf("[GET /deals/comments/{comment_id}][%d] getDealsCommentsCommentIdOK  %+v", 200, o.Payload)
}

func (o *GetDealsCommentsCommentIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.CommentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDealsCommentsCommentIDNotFound creates a GetDealsCommentsCommentIDNotFound with default headers values
func NewGetDealsCommentsCommentIDNotFound() *GetDealsCommentsCommentIDNotFound {
	return &GetDealsCommentsCommentIDNotFound{}
}

/*GetDealsCommentsCommentIDNotFound handles this case with default header values.

Deal comment not found.
*/
type GetDealsCommentsCommentIDNotFound struct {
	Payload *models_core.Http404Response
}

func (o *GetDealsCommentsCommentIDNotFound) Error() string {
	return fmt.Sprintf("[GET /deals/comments/{comment_id}][%d] getDealsCommentsCommentIdNotFound  %+v", 404, o.Payload)
}

func (o *GetDealsCommentsCommentIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.Http404Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}