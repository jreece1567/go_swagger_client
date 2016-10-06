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

// DeleteStoresStoreIDTradingHoursTradingHourIDReader is a Reader for the DeleteStoresStoreIDTradingHoursTradingHourID structure.
type DeleteStoresStoreIDTradingHoursTradingHourIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeleteStoresStoreIDTradingHoursTradingHourIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteStoresStoreIDTradingHoursTradingHourIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteStoresStoreIDTradingHoursTradingHourIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteStoresStoreIDTradingHoursTradingHourIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteStoresStoreIDTradingHoursTradingHourIDNoContent creates a DeleteStoresStoreIDTradingHoursTradingHourIDNoContent with default headers values
func NewDeleteStoresStoreIDTradingHoursTradingHourIDNoContent() *DeleteStoresStoreIDTradingHoursTradingHourIDNoContent {
	return &DeleteStoresStoreIDTradingHoursTradingHourIDNoContent{}
}

/*DeleteStoresStoreIDTradingHoursTradingHourIDNoContent handles this case with default header values.

Successfully deleted the store trading-hour.
*/
type DeleteStoresStoreIDTradingHoursTradingHourIDNoContent struct {
}

func (o *DeleteStoresStoreIDTradingHoursTradingHourIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /stores/{store_id}/trading-hours/{trading_hour_id}][%d] deleteStoresStoreIdTradingHoursTradingHourIdNoContent ", 204)
}

func (o *DeleteStoresStoreIDTradingHoursTradingHourIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteStoresStoreIDTradingHoursTradingHourIDUnauthorized creates a DeleteStoresStoreIDTradingHoursTradingHourIDUnauthorized with default headers values
func NewDeleteStoresStoreIDTradingHoursTradingHourIDUnauthorized() *DeleteStoresStoreIDTradingHoursTradingHourIDUnauthorized {
	return &DeleteStoresStoreIDTradingHoursTradingHourIDUnauthorized{}
}

/*DeleteStoresStoreIDTradingHoursTradingHourIDUnauthorized handles this case with default header values.

Unauthorized request.
*/
type DeleteStoresStoreIDTradingHoursTradingHourIDUnauthorized struct {
	Payload *models_core.Http401Response
}

func (o *DeleteStoresStoreIDTradingHoursTradingHourIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /stores/{store_id}/trading-hours/{trading_hour_id}][%d] deleteStoresStoreIdTradingHoursTradingHourIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteStoresStoreIDTradingHoursTradingHourIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.Http401Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteStoresStoreIDTradingHoursTradingHourIDNotFound creates a DeleteStoresStoreIDTradingHoursTradingHourIDNotFound with default headers values
func NewDeleteStoresStoreIDTradingHoursTradingHourIDNotFound() *DeleteStoresStoreIDTradingHoursTradingHourIDNotFound {
	return &DeleteStoresStoreIDTradingHoursTradingHourIDNotFound{}
}

/*DeleteStoresStoreIDTradingHoursTradingHourIDNotFound handles this case with default header values.

Store trading-hour or store not found.
*/
type DeleteStoresStoreIDTradingHoursTradingHourIDNotFound struct {
	Payload *models_core.Http404Response
}

func (o *DeleteStoresStoreIDTradingHoursTradingHourIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /stores/{store_id}/trading-hours/{trading_hour_id}][%d] deleteStoresStoreIdTradingHoursTradingHourIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteStoresStoreIDTradingHoursTradingHourIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_core.Http404Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}