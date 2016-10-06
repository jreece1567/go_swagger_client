package s_t_o_r_e_t_r_a_d_i_n_g_h_o_u_r_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"restclient/models_core"
)

// PatchOccasionsOccasionIDReader is a Reader for the PatchOccasionsOccasionID structure.
type PatchOccasionsOccasionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PatchOccasionsOccasionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPatchOccasionsOccasionIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPatchOccasionsOccasionIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchOccasionsOccasionIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPatchOccasionsOccasionIDUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchOccasionsOccasionIDNoContent creates a PatchOccasionsOccasionIDNoContent with default headers values
func NewPatchOccasionsOccasionIDNoContent() *PatchOccasionsOccasionIDNoContent {
	return &PatchOccasionsOccasionIDNoContent{}
}

/*PatchOccasionsOccasionIDNoContent handles this case with default header values.

Successfully updated the occasion.
*/
type PatchOccasionsOccasionIDNoContent struct {
}

func (o *PatchOccasionsOccasionIDNoContent) Error() string {
	return fmt.Sprintf("[PATCH /occasions/{occasion_id}][%d] patchOccasionsOccasionIdNoContent ", 204)
}

func (o *PatchOccasionsOccasionIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchOccasionsOccasionIDUnauthorized creates a PatchOccasionsOccasionIDUnauthorized with default headers values
func NewPatchOccasionsOccasionIDUnauthorized() *PatchOccasionsOccasionIDUnauthorized {
	return &PatchOccasionsOccasionIDUnauthorized{}
}

/*PatchOccasionsOccasionIDUnauthorized handles this case with default header values.

Unauthorized request.
*/
type PatchOccasionsOccasionIDUnauthorized struct {
	Payload *models_core.Http401Response
}

func (o *PatchOccasionsOccasionIDUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /occasions/{occasion_id}][%d] patchOccasionsOccasionIdUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchOccasionsOccasionIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.Http401Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchOccasionsOccasionIDNotFound creates a PatchOccasionsOccasionIDNotFound with default headers values
func NewPatchOccasionsOccasionIDNotFound() *PatchOccasionsOccasionIDNotFound {
	return &PatchOccasionsOccasionIDNotFound{}
}

/*PatchOccasionsOccasionIDNotFound handles this case with default header values.

Occasion not found.
*/
type PatchOccasionsOccasionIDNotFound struct {
	Payload *models_core.Http404Response
}

func (o *PatchOccasionsOccasionIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /occasions/{occasion_id}][%d] patchOccasionsOccasionIdNotFound  %+v", 404, o.Payload)
}

func (o *PatchOccasionsOccasionIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.Http404Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchOccasionsOccasionIDUnprocessableEntity creates a PatchOccasionsOccasionIDUnprocessableEntity with default headers values
func NewPatchOccasionsOccasionIDUnprocessableEntity() *PatchOccasionsOccasionIDUnprocessableEntity {
	return &PatchOccasionsOccasionIDUnprocessableEntity{}
}

/*PatchOccasionsOccasionIDUnprocessableEntity handles this case with default header values.

Occasion not updated. Returns an array of error messages explaining the problems with the provided attributes.
*/
type PatchOccasionsOccasionIDUnprocessableEntity struct {
	Payload *models_core.Occasion422Response
}

func (o *PatchOccasionsOccasionIDUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /occasions/{occasion_id}][%d] patchOccasionsOccasionIdUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PatchOccasionsOccasionIDUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.Occasion422Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
