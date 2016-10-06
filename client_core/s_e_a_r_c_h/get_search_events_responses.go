package s_e_a_r_c_h

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"restclient/models_core"
)

// GetSearchEventsReader is a Reader for the GetSearchEvents structure.
type GetSearchEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetSearchEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSearchEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSearchEventsOK creates a GetSearchEventsOK with default headers values
func NewGetSearchEventsOK() *GetSearchEventsOK {
	return &GetSearchEventsOK{}
}

/*GetSearchEventsOK handles this case with default header values.

Successfully retrieved the search results.
*/
type GetSearchEventsOK struct {
	Payload *models_core.SearchEventResponse
}

func (o *GetSearchEventsOK) Error() string {
	return fmt.Sprintf("[GET /search/events][%d] getSearchEventsOK  %+v", 200, o.Payload)
}

func (o *GetSearchEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.SearchEventResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}