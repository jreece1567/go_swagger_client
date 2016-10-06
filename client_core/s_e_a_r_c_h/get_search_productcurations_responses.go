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

// GetSearchProductcurationsReader is a Reader for the GetSearchProductcurations structure.
type GetSearchProductcurationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetSearchProductcurationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSearchProductcurationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSearchProductcurationsOK creates a GetSearchProductcurationsOK with default headers values
func NewGetSearchProductcurationsOK() *GetSearchProductcurationsOK {
	return &GetSearchProductcurationsOK{}
}

/*GetSearchProductcurationsOK handles this case with default header values.

Successfully retrieved the search results.
*/
type GetSearchProductcurationsOK struct {
	Payload *models_core.SearchCollectionResponse
}

func (o *GetSearchProductcurationsOK) Error() string {
	return fmt.Sprintf("[GET /search/productcurations][%d] getSearchProductcurationsOK  %+v", 200, o.Payload)
}

func (o *GetSearchProductcurationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.SearchCollectionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}