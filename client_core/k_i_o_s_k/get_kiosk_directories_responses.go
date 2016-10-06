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

// GetKioskDirectoriesReader is a Reader for the GetKioskDirectories structure.
type GetKioskDirectoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetKioskDirectoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetKioskDirectoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetKioskDirectoriesOK creates a GetKioskDirectoriesOK with default headers values
func NewGetKioskDirectoriesOK() *GetKioskDirectoriesOK {
	return &GetKioskDirectoriesOK{}
}

/*GetKioskDirectoriesOK handles this case with default header values.

Successfully retrieved the list of directories.
*/
type GetKioskDirectoriesOK struct {
	Payload *models_core.DirectoryListResponse
}

func (o *GetKioskDirectoriesOK) Error() string {
	return fmt.Sprintf("[GET /kiosk/directories][%d] getKioskDirectoriesOK  %+v", 200, o.Payload)
}

func (o *GetKioskDirectoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.DirectoryListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
