package s_t_a_f_f_a_u_t_h

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostStaffRevokeReader is a Reader for the PostStaffRevoke structure.
type PostStaffRevokeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PostStaffRevokeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostStaffRevokeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostStaffRevokeOK creates a PostStaffRevokeOK with default headers values
func NewPostStaffRevokeOK() *PostStaffRevokeOK {
	return &PostStaffRevokeOK{}
}

/*PostStaffRevokeOK handles this case with default header values.

Success
*/
type PostStaffRevokeOK struct {
}

func (o *PostStaffRevokeOK) Error() string {
	return fmt.Sprintf("[POST /staff/revoke][%d] postStaffRevokeOK ", 200)
}

func (o *PostStaffRevokeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}