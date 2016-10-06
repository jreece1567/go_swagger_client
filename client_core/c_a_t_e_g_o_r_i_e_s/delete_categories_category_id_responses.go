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

// DeleteCategoriesCategoryIDReader is a Reader for the DeleteCategoriesCategoryID structure.
type DeleteCategoriesCategoryIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeleteCategoriesCategoryIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteCategoriesCategoryIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteCategoriesCategoryIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteCategoriesCategoryIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteCategoriesCategoryIDNoContent creates a DeleteCategoriesCategoryIDNoContent with default headers values
func NewDeleteCategoriesCategoryIDNoContent() *DeleteCategoriesCategoryIDNoContent {
	return &DeleteCategoriesCategoryIDNoContent{}
}

/*DeleteCategoriesCategoryIDNoContent handles this case with default header values.

Successfully deleted the category.
*/
type DeleteCategoriesCategoryIDNoContent struct {
}

func (o *DeleteCategoriesCategoryIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /categories/{category_id}][%d] deleteCategoriesCategoryIdNoContent ", 204)
}

func (o *DeleteCategoriesCategoryIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCategoriesCategoryIDUnauthorized creates a DeleteCategoriesCategoryIDUnauthorized with default headers values
func NewDeleteCategoriesCategoryIDUnauthorized() *DeleteCategoriesCategoryIDUnauthorized {
	return &DeleteCategoriesCategoryIDUnauthorized{}
}

/*DeleteCategoriesCategoryIDUnauthorized handles this case with default header values.

Unauthorized request.
*/
type DeleteCategoriesCategoryIDUnauthorized struct {
	Payload *models_core.Http401Response
}

func (o *DeleteCategoriesCategoryIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /categories/{category_id}][%d] deleteCategoriesCategoryIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteCategoriesCategoryIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.Http401Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCategoriesCategoryIDNotFound creates a DeleteCategoriesCategoryIDNotFound with default headers values
func NewDeleteCategoriesCategoryIDNotFound() *DeleteCategoriesCategoryIDNotFound {
	return &DeleteCategoriesCategoryIDNotFound{}
}

/*DeleteCategoriesCategoryIDNotFound handles this case with default header values.

Category not found.
*/
type DeleteCategoriesCategoryIDNotFound struct {
	Payload *models_core.Http404Response
}

func (o *DeleteCategoriesCategoryIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /categories/{category_id}][%d] deleteCategoriesCategoryIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteCategoriesCategoryIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.Http404Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}