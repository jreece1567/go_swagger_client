package c_a_t_e_g_o_r_i_e_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"restclient/models_core"
)

// GetCategoriesGroupsReader is a Reader for the GetCategoriesGroups structure.
type GetCategoriesGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetCategoriesGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCategoriesGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCategoriesGroupsOK creates a GetCategoriesGroupsOK with default headers values
func NewGetCategoriesGroupsOK() *GetCategoriesGroupsOK {
	return &GetCategoriesGroupsOK{}
}

/*GetCategoriesGroupsOK handles this case with default header values.

Successfully retrieved the list of groups.
*/
type GetCategoriesGroupsOK struct {
	Payload *models_core.GroupListResponse
}

func (o *GetCategoriesGroupsOK) Error() string {
	return fmt.Sprintf("[GET /categories/groups][%d] getCategoriesGroupsOK  %+v", 200, o.Payload)
}

func (o *GetCategoriesGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.GroupListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
