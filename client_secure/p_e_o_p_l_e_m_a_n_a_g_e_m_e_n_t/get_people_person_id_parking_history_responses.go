package p_e_o_p_l_e_m_a_n_a_g_e_m_e_n_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"restclient/models_secure"
)

// GetPeoplePersonIDParkingHistoryReader is a Reader for the GetPeoplePersonIDParkingHistory structure.
type GetPeoplePersonIDParkingHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetPeoplePersonIDParkingHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPeoplePersonIDParkingHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetPeoplePersonIDParkingHistoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPeoplePersonIDParkingHistoryOK creates a GetPeoplePersonIDParkingHistoryOK with default headers values
func NewGetPeoplePersonIDParkingHistoryOK() *GetPeoplePersonIDParkingHistoryOK {
	return &GetPeoplePersonIDParkingHistoryOK{}
}

/*GetPeoplePersonIDParkingHistoryOK handles this case with default header values.

Successfully retrieved the list of parking sessions history.
*/
type GetPeoplePersonIDParkingHistoryOK struct {
	Payload *models_secure.ParkingHistoryResponse
}

func (o *GetPeoplePersonIDParkingHistoryOK) Error() string {
	return fmt.Sprintf("[GET /people/{person_id}/parking/history][%d] getPeoplePersonIdParkingHistoryOK  %+v", 200, o.Payload)
}

func (o *GetPeoplePersonIDParkingHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.ParkingHistoryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPeoplePersonIDParkingHistoryUnauthorized creates a GetPeoplePersonIDParkingHistoryUnauthorized with default headers values
func NewGetPeoplePersonIDParkingHistoryUnauthorized() *GetPeoplePersonIDParkingHistoryUnauthorized {
	return &GetPeoplePersonIDParkingHistoryUnauthorized{}
}

/*GetPeoplePersonIDParkingHistoryUnauthorized handles this case with default header values.

Unauthorized request.
*/
type GetPeoplePersonIDParkingHistoryUnauthorized struct {
	Payload *models_secure.UnauthorizedResponse
}

func (o *GetPeoplePersonIDParkingHistoryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /people/{person_id}/parking/history][%d] getPeoplePersonIdParkingHistoryUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPeoplePersonIDParkingHistoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.UnauthorizedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}