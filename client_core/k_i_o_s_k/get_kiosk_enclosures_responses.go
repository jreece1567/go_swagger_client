package k_i_o_s_k

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"restclient/models_core"
)

// GetKioskEnclosuresReader is a Reader for the GetKioskEnclosures structure.
type GetKioskEnclosuresReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetKioskEnclosuresReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetKioskEnclosuresOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetKioskEnclosuresOK creates a GetKioskEnclosuresOK with default headers values
func NewGetKioskEnclosuresOK() *GetKioskEnclosuresOK {
	return &GetKioskEnclosuresOK{}
}

/*GetKioskEnclosuresOK handles this case with default header values.

Successfully retrieved the list of enclosures.
*/
type GetKioskEnclosuresOK struct {
	Payload *models_core.EnclosuresListResponse
}

func (o *GetKioskEnclosuresOK) Error() string {
	return fmt.Sprintf("[GET /kiosk/enclosures][%d] getKioskEnclosuresOK  %+v", 200, o.Payload)
}

func (o *GetKioskEnclosuresOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.EnclosuresListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}