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

// GetCategoriesGroupsGroupIDReader is a Reader for the GetCategoriesGroupsGroupID structure.
type GetCategoriesGroupsGroupIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetCategoriesGroupsGroupIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCategoriesGroupsGroupIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetCategoriesGroupsGroupIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCategoriesGroupsGroupIDOK creates a GetCategoriesGroupsGroupIDOK with default headers values
func NewGetCategoriesGroupsGroupIDOK() *GetCategoriesGroupsGroupIDOK {
	return &GetCategoriesGroupsGroupIDOK{}
}

/*GetCategoriesGroupsGroupIDOK handles this case with default header values.

Successfully retrieved the individual group.
*/
type GetCategoriesGroupsGroupIDOK struct {
	Payload *models_core.GroupResponse
}

func (o *GetCategoriesGroupsGroupIDOK) Error() string {
	return fmt.Sprintf("[GET /categories/groups/{group_id}][%d] getCategoriesGroupsGroupIdOK  %+v", 200, o.Payload)
}

func (o *GetCategoriesGroupsGroupIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.GroupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCategoriesGroupsGroupIDNotFound creates a GetCategoriesGroupsGroupIDNotFound with default headers values
func NewGetCategoriesGroupsGroupIDNotFound() *GetCategoriesGroupsGroupIDNotFound {
	return &GetCategoriesGroupsGroupIDNotFound{}
}

/*GetCategoriesGroupsGroupIDNotFound handles this case with default header values.

Category not found.
*/
type GetCategoriesGroupsGroupIDNotFound struct {
	Payload *models_core.Http404Response
}

func (o *GetCategoriesGroupsGroupIDNotFound) Error() string {
	return fmt.Sprintf("[GET /categories/groups/{group_id}][%d] getCategoriesGroupsGroupIdNotFound  %+v", 404, o.Payload)
}

func (o *GetCategoriesGroupsGroupIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.Http404Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
