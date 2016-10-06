package p_e_o_p_l_e_c_o_n_s_u_m_e_r

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"restclient/models_secure"
)

// GetAccountParkingSessionsSummaryReader is a Reader for the GetAccountParkingSessionsSummary structure.
type GetAccountParkingSessionsSummaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetAccountParkingSessionsSummaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAccountParkingSessionsSummaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetAccountParkingSessionsSummaryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAccountParkingSessionsSummaryOK creates a GetAccountParkingSessionsSummaryOK with default headers values
func NewGetAccountParkingSessionsSummaryOK() *GetAccountParkingSessionsSummaryOK {
	return &GetAccountParkingSessionsSummaryOK{}
}

/*GetAccountParkingSessionsSummaryOK handles this case with default header values.

Successfully retrieved the list of parking sessions summary.
*/
type GetAccountParkingSessionsSummaryOK struct {
	Payload *models_secure.AccountParkingSessionsSummaryResponse
}

func (o *GetAccountParkingSessionsSummaryOK) Error() string {
	return fmt.Sprintf("[GET /account/parking/sessions/summary][%d] getAccountParkingSessionsSummaryOK  %+v", 200, o.Payload)
}

func (o *GetAccountParkingSessionsSummaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.AccountParkingSessionsSummaryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountParkingSessionsSummaryUnauthorized creates a GetAccountParkingSessionsSummaryUnauthorized with default headers values
func NewGetAccountParkingSessionsSummaryUnauthorized() *GetAccountParkingSessionsSummaryUnauthorized {
	return &GetAccountParkingSessionsSummaryUnauthorized{}
}

/*GetAccountParkingSessionsSummaryUnauthorized handles this case with default header values.

Unauthorized request.
*/
type GetAccountParkingSessionsSummaryUnauthorized struct {
	Payload *models_secure.UnauthorizedResponse
}

func (o *GetAccountParkingSessionsSummaryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /account/parking/sessions/summary][%d] getAccountParkingSessionsSummaryUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAccountParkingSessionsSummaryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_secure.UnauthorizedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
