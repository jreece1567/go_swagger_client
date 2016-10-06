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

// PatchCategoriesCategoryIDLocalesWestfieldLocaleReader is a Reader for the PatchCategoriesCategoryIDLocalesWestfieldLocale structure.
type PatchCategoriesCategoryIDLocalesWestfieldLocaleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PatchCategoriesCategoryIDLocalesWestfieldLocaleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPatchCategoriesCategoryIDLocalesWestfieldLocaleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPatchCategoriesCategoryIDLocalesWestfieldLocaleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchCategoriesCategoryIDLocalesWestfieldLocaleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPatchCategoriesCategoryIDLocalesWestfieldLocaleUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchCategoriesCategoryIDLocalesWestfieldLocaleNoContent creates a PatchCategoriesCategoryIDLocalesWestfieldLocaleNoContent with default headers values
func NewPatchCategoriesCategoryIDLocalesWestfieldLocaleNoContent() *PatchCategoriesCategoryIDLocalesWestfieldLocaleNoContent {
	return &PatchCategoriesCategoryIDLocalesWestfieldLocaleNoContent{}
}

/*PatchCategoriesCategoryIDLocalesWestfieldLocaleNoContent handles this case with default header values.

Successfully updated the locale.
*/
type PatchCategoriesCategoryIDLocalesWestfieldLocaleNoContent struct {
}

func (o *PatchCategoriesCategoryIDLocalesWestfieldLocaleNoContent) Error() string {
	return fmt.Sprintf("[PATCH /categories/{category_id}/locales/{westfield_locale}][%d] patchCategoriesCategoryIdLocalesWestfieldLocaleNoContent ", 204)
}

func (o *PatchCategoriesCategoryIDLocalesWestfieldLocaleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchCategoriesCategoryIDLocalesWestfieldLocaleUnauthorized creates a PatchCategoriesCategoryIDLocalesWestfieldLocaleUnauthorized with default headers values
func NewPatchCategoriesCategoryIDLocalesWestfieldLocaleUnauthorized() *PatchCategoriesCategoryIDLocalesWestfieldLocaleUnauthorized {
	return &PatchCategoriesCategoryIDLocalesWestfieldLocaleUnauthorized{}
}

/*PatchCategoriesCategoryIDLocalesWestfieldLocaleUnauthorized handles this case with default header values.

Unauthorized request.
*/
type PatchCategoriesCategoryIDLocalesWestfieldLocaleUnauthorized struct {
	Payload *models_core.Http401Response
}

func (o *PatchCategoriesCategoryIDLocalesWestfieldLocaleUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /categories/{category_id}/locales/{westfield_locale}][%d] patchCategoriesCategoryIdLocalesWestfieldLocaleUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchCategoriesCategoryIDLocalesWestfieldLocaleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.Http401Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchCategoriesCategoryIDLocalesWestfieldLocaleNotFound creates a PatchCategoriesCategoryIDLocalesWestfieldLocaleNotFound with default headers values
func NewPatchCategoriesCategoryIDLocalesWestfieldLocaleNotFound() *PatchCategoriesCategoryIDLocalesWestfieldLocaleNotFound {
	return &PatchCategoriesCategoryIDLocalesWestfieldLocaleNotFound{}
}

/*PatchCategoriesCategoryIDLocalesWestfieldLocaleNotFound handles this case with default header values.

Locale not found.
*/
type PatchCategoriesCategoryIDLocalesWestfieldLocaleNotFound struct {
	Payload *models_core.Http404Response
}

func (o *PatchCategoriesCategoryIDLocalesWestfieldLocaleNotFound) Error() string {
	return fmt.Sprintf("[PATCH /categories/{category_id}/locales/{westfield_locale}][%d] patchCategoriesCategoryIdLocalesWestfieldLocaleNotFound  %+v", 404, o.Payload)
}

func (o *PatchCategoriesCategoryIDLocalesWestfieldLocaleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.Http404Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchCategoriesCategoryIDLocalesWestfieldLocaleUnprocessableEntity creates a PatchCategoriesCategoryIDLocalesWestfieldLocaleUnprocessableEntity with default headers values
func NewPatchCategoriesCategoryIDLocalesWestfieldLocaleUnprocessableEntity() *PatchCategoriesCategoryIDLocalesWestfieldLocaleUnprocessableEntity {
	return &PatchCategoriesCategoryIDLocalesWestfieldLocaleUnprocessableEntity{}
}

/*PatchCategoriesCategoryIDLocalesWestfieldLocaleUnprocessableEntity handles this case with default header values.

Locale not updated. Returns an array of error messages explaining the problems with the provided attributes.
*/
type PatchCategoriesCategoryIDLocalesWestfieldLocaleUnprocessableEntity struct {
	Payload *models_core.Locale422Response
}

func (o *PatchCategoriesCategoryIDLocalesWestfieldLocaleUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /categories/{category_id}/locales/{westfield_locale}][%d] patchCategoriesCategoryIdLocalesWestfieldLocaleUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PatchCategoriesCategoryIDLocalesWestfieldLocaleUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.Locale422Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}