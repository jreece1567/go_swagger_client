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

// GetCategoriesLocalesWestfieldLocaleReader is a Reader for the GetCategoriesLocalesWestfieldLocale structure.
type GetCategoriesLocalesWestfieldLocaleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetCategoriesLocalesWestfieldLocaleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCategoriesLocalesWestfieldLocaleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetCategoriesLocalesWestfieldLocaleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCategoriesLocalesWestfieldLocaleOK creates a GetCategoriesLocalesWestfieldLocaleOK with default headers values
func NewGetCategoriesLocalesWestfieldLocaleOK() *GetCategoriesLocalesWestfieldLocaleOK {
	return &GetCategoriesLocalesWestfieldLocaleOK{}
}

/*GetCategoriesLocalesWestfieldLocaleOK handles this case with default header values.

Successfully retrieved the individual category locale.
*/
type GetCategoriesLocalesWestfieldLocaleOK struct {
	Payload *models_core.LocaleListResponse
}

func (o *GetCategoriesLocalesWestfieldLocaleOK) Error() string {
	return fmt.Sprintf("[GET /categories/locales/{westfield_locale}][%d] getCategoriesLocalesWestfieldLocaleOK  %+v", 200, o.Payload)
}

func (o *GetCategoriesLocalesWestfieldLocaleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.LocaleListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCategoriesLocalesWestfieldLocaleNotFound creates a GetCategoriesLocalesWestfieldLocaleNotFound with default headers values
func NewGetCategoriesLocalesWestfieldLocaleNotFound() *GetCategoriesLocalesWestfieldLocaleNotFound {
	return &GetCategoriesLocalesWestfieldLocaleNotFound{}
}

/*GetCategoriesLocalesWestfieldLocaleNotFound handles this case with default header values.

Category locales not found.
*/
type GetCategoriesLocalesWestfieldLocaleNotFound struct {
	Payload *models_core.Http404Response
}

func (o *GetCategoriesLocalesWestfieldLocaleNotFound) Error() string {
	return fmt.Sprintf("[GET /categories/locales/{westfield_locale}][%d] getCategoriesLocalesWestfieldLocaleNotFound  %+v", 404, o.Payload)
}

func (o *GetCategoriesLocalesWestfieldLocaleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.Http404Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
