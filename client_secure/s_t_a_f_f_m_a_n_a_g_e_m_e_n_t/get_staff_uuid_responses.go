package s_t_a_f_f_m_a_n_a_g_e_m_e_n_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"restclient/models_secure"
)

// GetStaffUUIDReader is a Reader for the GetStaffUUID structure.
type GetStaffUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetStaffUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetStaffUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetStaffUUIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetStaffUUIDOK creates a GetStaffUUIDOK with default headers values
func NewGetStaffUUIDOK() *GetStaffUUIDOK {
	return &GetStaffUUIDOK{}
}

/*GetStaffUUIDOK handles this case with default header values.

A staff members
*/
type GetStaffUUIDOK struct {
	Payload *models_secure.Staff
}

func (o *GetStaffUUIDOK) Error() string {
	return fmt.Sprintf("[GET /staff/{uuid}][%d] getStaffUuidOK  %+v", 200, o.Payload)
}

func (o *GetStaffUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.Staff)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStaffUUIDUnauthorized creates a GetStaffUUIDUnauthorized with default headers values
func NewGetStaffUUIDUnauthorized() *GetStaffUUIDUnauthorized {
	return &GetStaffUUIDUnauthorized{}
}

/*GetStaffUUIDUnauthorized handles this case with default header values.

Unauthorized.
*/
type GetStaffUUIDUnauthorized struct {
	Payload *models_secure.Staff401Response
}

func (o *GetStaffUUIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /staff/{uuid}][%d] getStaffUuidUnauthorized  %+v", 401, o.Payload)
}

func (o *GetStaffUUIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.Staff401Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
