package s_t_o_r_e_i_n_f_o_r_m_a_t_i_o_n

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"restclient/models_core"
)

// GetStoresReader is a Reader for the GetStores structure.
type GetStoresReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetStoresReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetStoresOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetStoresOK creates a GetStoresOK with default headers values
func NewGetStoresOK() *GetStoresOK {
	return &GetStoresOK{}
}

/*GetStoresOK handles this case with default header values.

Successfully retrieved the list of stores.
*/
type GetStoresOK struct {
	Payload *models_core.StoresListResponse
}

func (o *GetStoresOK) Error() string {
	return fmt.Sprintf("[GET /stores][%d] getStoresOK  %+v", 200, o.Payload)
}

func (o *GetStoresOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.StoresListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
