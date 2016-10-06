package c_e_n_t_r_e_n_o_t_i_c_e_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"restclient/models_core"
)

// GetNoticesReader is a Reader for the GetNotices structure.
type GetNoticesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetNoticesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNoticesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNoticesOK creates a GetNoticesOK with default headers values
func NewGetNoticesOK() *GetNoticesOK {
	return &GetNoticesOK{}
}

/*GetNoticesOK handles this case with default header values.

Successfully retrieved the list of notices.
*/
type GetNoticesOK struct {
	Payload *models_core.NoticeListResponse
}

func (o *GetNoticesOK) Error() string {
	return fmt.Sprintf("[GET /notices][%d] getNoticesOK  %+v", 200, o.Payload)
}

func (o *GetNoticesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.NoticeListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
