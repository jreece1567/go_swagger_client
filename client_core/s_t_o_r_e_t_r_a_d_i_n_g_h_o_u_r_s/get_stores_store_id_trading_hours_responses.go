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

// GetStoresStoreIDTradingHoursReader is a Reader for the GetStoresStoreIDTradingHours structure.
type GetStoresStoreIDTradingHoursReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetStoresStoreIDTradingHoursReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetStoresStoreIDTradingHoursOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetStoresStoreIDTradingHoursOK creates a GetStoresStoreIDTradingHoursOK with default headers values
func NewGetStoresStoreIDTradingHoursOK() *GetStoresStoreIDTradingHoursOK {
	return &GetStoresStoreIDTradingHoursOK{}
}

/*GetStoresStoreIDTradingHoursOK handles this case with default header values.

Successfully retrieved the list of store trading hours.
*/
type GetStoresStoreIDTradingHoursOK struct {
	Payload *models_core.StoreTradingHoursListResponse
}

func (o *GetStoresStoreIDTradingHoursOK) Error() string {
	return fmt.Sprintf("[GET /stores/{store_id}/trading-hours][%d] getStoresStoreIdTradingHoursOK  %+v", 200, o.Payload)
}

func (o *GetStoresStoreIDTradingHoursOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.StoreTradingHoursListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
