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

// GetCategoriesCategoryIDLocalesWestfieldLocaleChildrenTreeReader is a Reader for the GetCategoriesCategoryIDLocalesWestfieldLocaleChildrenTree structure.
type GetCategoriesCategoryIDLocalesWestfieldLocaleChildrenTreeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetCategoriesCategoryIDLocalesWestfieldLocaleChildrenTreeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCategoriesCategoryIDLocalesWestfieldLocaleChildrenTreeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetCategoriesCategoryIDLocalesWestfieldLocaleChildrenTreeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCategoriesCategoryIDLocalesWestfieldLocaleChildrenTreeOK creates a GetCategoriesCategoryIDLocalesWestfieldLocaleChildrenTreeOK with default headers values
func NewGetCategoriesCategoryIDLocalesWestfieldLocaleChildrenTreeOK() *GetCategoriesCategoryIDLocalesWestfieldLocaleChildrenTreeOK {
	return &GetCategoriesCategoryIDLocalesWestfieldLocaleChildrenTreeOK{}
}

/*GetCategoriesCategoryIDLocalesWestfieldLocaleChildrenTreeOK handles this case with default header values.

Successfully retrieved the category tree with locale.
*/
type GetCategoriesCategoryIDLocalesWestfieldLocaleChildrenTreeOK struct {
	Payload *models_core.CategoryChildrenTreeResponse
}

func (o *GetCategoriesCategoryIDLocalesWestfieldLocaleChildrenTreeOK) Error() string {
	return fmt.Sprintf("[GET /categories/{category_id}/locales/{westfield_locale}/children-tree][%d] getCategoriesCategoryIdLocalesWestfieldLocaleChildrenTreeOK  %+v", 200, o.Payload)
}

func (o *GetCategoriesCategoryIDLocalesWestfieldLocaleChildrenTreeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.CategoryChildrenTreeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCategoriesCategoryIDLocalesWestfieldLocaleChildrenTreeNotFound creates a GetCategoriesCategoryIDLocalesWestfieldLocaleChildrenTreeNotFound with default headers values
func NewGetCategoriesCategoryIDLocalesWestfieldLocaleChildrenTreeNotFound() *GetCategoriesCategoryIDLocalesWestfieldLocaleChildrenTreeNotFound {
	return &GetCategoriesCategoryIDLocalesWestfieldLocaleChildrenTreeNotFound{}
}

/*GetCategoriesCategoryIDLocalesWestfieldLocaleChildrenTreeNotFound handles this case with default header values.

Category locale not found.
*/
type GetCategoriesCategoryIDLocalesWestfieldLocaleChildrenTreeNotFound struct {
	Payload *models_core.Http404Response
}

func (o *GetCategoriesCategoryIDLocalesWestfieldLocaleChildrenTreeNotFound) Error() string {
	return fmt.Sprintf("[GET /categories/{category_id}/locales/{westfield_locale}/children-tree][%d] getCategoriesCategoryIdLocalesWestfieldLocaleChildrenTreeNotFound  %+v", 404, o.Payload)
}

func (o *GetCategoriesCategoryIDLocalesWestfieldLocaleChildrenTreeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.Http404Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
