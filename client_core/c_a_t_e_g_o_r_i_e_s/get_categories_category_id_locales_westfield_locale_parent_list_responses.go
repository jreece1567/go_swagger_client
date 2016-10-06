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

// GetCategoriesCategoryIDLocalesWestfieldLocaleParentListReader is a Reader for the GetCategoriesCategoryIDLocalesWestfieldLocaleParentList structure.
type GetCategoriesCategoryIDLocalesWestfieldLocaleParentListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetCategoriesCategoryIDLocalesWestfieldLocaleParentListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCategoriesCategoryIDLocalesWestfieldLocaleParentListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetCategoriesCategoryIDLocalesWestfieldLocaleParentListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCategoriesCategoryIDLocalesWestfieldLocaleParentListOK creates a GetCategoriesCategoryIDLocalesWestfieldLocaleParentListOK with default headers values
func NewGetCategoriesCategoryIDLocalesWestfieldLocaleParentListOK() *GetCategoriesCategoryIDLocalesWestfieldLocaleParentListOK {
	return &GetCategoriesCategoryIDLocalesWestfieldLocaleParentListOK{}
}

/*GetCategoriesCategoryIDLocalesWestfieldLocaleParentListOK handles this case with default header values.

Successfully retrieved parents of the category with locale.
*/
type GetCategoriesCategoryIDLocalesWestfieldLocaleParentListOK struct {
	Payload *models_core.CategoryChildrenListResponse
}

func (o *GetCategoriesCategoryIDLocalesWestfieldLocaleParentListOK) Error() string {
	return fmt.Sprintf("[GET /categories/{category_id}/locales/{westfield_locale}/parent-list][%d] getCategoriesCategoryIdLocalesWestfieldLocaleParentListOK  %+v", 200, o.Payload)
}

func (o *GetCategoriesCategoryIDLocalesWestfieldLocaleParentListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.CategoryChildrenListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCategoriesCategoryIDLocalesWestfieldLocaleParentListNotFound creates a GetCategoriesCategoryIDLocalesWestfieldLocaleParentListNotFound with default headers values
func NewGetCategoriesCategoryIDLocalesWestfieldLocaleParentListNotFound() *GetCategoriesCategoryIDLocalesWestfieldLocaleParentListNotFound {
	return &GetCategoriesCategoryIDLocalesWestfieldLocaleParentListNotFound{}
}

/*GetCategoriesCategoryIDLocalesWestfieldLocaleParentListNotFound handles this case with default header values.

Category locale not found.
*/
type GetCategoriesCategoryIDLocalesWestfieldLocaleParentListNotFound struct {
	Payload *models_core.Http404Response
}

func (o *GetCategoriesCategoryIDLocalesWestfieldLocaleParentListNotFound) Error() string {
	return fmt.Sprintf("[GET /categories/{category_id}/locales/{westfield_locale}/parent-list][%d] getCategoriesCategoryIdLocalesWestfieldLocaleParentListNotFound  %+v", 404, o.Payload)
}

func (o *GetCategoriesCategoryIDLocalesWestfieldLocaleParentListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.Http404Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}