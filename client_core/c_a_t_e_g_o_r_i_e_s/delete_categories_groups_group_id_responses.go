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

// DeleteCategoriesGroupsGroupIDReader is a Reader for the DeleteCategoriesGroupsGroupID structure.
type DeleteCategoriesGroupsGroupIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeleteCategoriesGroupsGroupIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteCategoriesGroupsGroupIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteCategoriesGroupsGroupIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteCategoriesGroupsGroupIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteCategoriesGroupsGroupIDNoContent creates a DeleteCategoriesGroupsGroupIDNoContent with default headers values
func NewDeleteCategoriesGroupsGroupIDNoContent() *DeleteCategoriesGroupsGroupIDNoContent {
	return &DeleteCategoriesGroupsGroupIDNoContent{}
}

/*DeleteCategoriesGroupsGroupIDNoContent handles this case with default header values.

Successfully deleted the group.
*/
type DeleteCategoriesGroupsGroupIDNoContent struct {
}

func (o *DeleteCategoriesGroupsGroupIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /categories/groups/{group_id}][%d] deleteCategoriesGroupsGroupIdNoContent ", 204)
}

func (o *DeleteCategoriesGroupsGroupIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCategoriesGroupsGroupIDUnauthorized creates a DeleteCategoriesGroupsGroupIDUnauthorized with default headers values
func NewDeleteCategoriesGroupsGroupIDUnauthorized() *DeleteCategoriesGroupsGroupIDUnauthorized {
	return &DeleteCategoriesGroupsGroupIDUnauthorized{}
}

/*DeleteCategoriesGroupsGroupIDUnauthorized handles this case with default header values.

Unauthorized request.
*/
type DeleteCategoriesGroupsGroupIDUnauthorized struct {
	Payload *models_core.Http401Response
}

func (o *DeleteCategoriesGroupsGroupIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /categories/groups/{group_id}][%d] deleteCategoriesGroupsGroupIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteCategoriesGroupsGroupIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.Http401Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCategoriesGroupsGroupIDNotFound creates a DeleteCategoriesGroupsGroupIDNotFound with default headers values
func NewDeleteCategoriesGroupsGroupIDNotFound() *DeleteCategoriesGroupsGroupIDNotFound {
	return &DeleteCategoriesGroupsGroupIDNotFound{}
}

/*DeleteCategoriesGroupsGroupIDNotFound handles this case with default header values.

Group not found.
*/
type DeleteCategoriesGroupsGroupIDNotFound struct {
	Payload *models_core.Http404Response
}

func (o *DeleteCategoriesGroupsGroupIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /categories/groups/{group_id}][%d] deleteCategoriesGroupsGroupIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteCategoriesGroupsGroupIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.Http404Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
