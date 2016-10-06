package c_e_n_t_r_e_p_a_r_k_i_n_g_r_a_t_e_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new c e n t r e p a r k i n g r a t e s API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for c e n t r e p a r k i n g r a t e s API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetCentresCentreIDCarparkReport lists of reportable carpark discounts applied at a centre

Request an array of reportable carpark discounts applied at a centre.
*/
func (a *Client) GetCentresCentreIDCarparkReport(params *GetCentresCentreIDCarparkReportParams, authInfo runtime.ClientAuthInfoWriter) (*GetCentresCentreIDCarparkReportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCentresCentreIDCarparkReportParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCentresCentreIDCarparkReport",
		Method:             "GET",
		PathPattern:        "/centres/{centre_id}/carpark-report",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCentresCentreIDCarparkReportReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCentresCentreIDCarparkReportOK), nil
}

/*
GetCentresCentreIDVehiclesVehicleID individuals vehicle

Request a single vehicle at a centre.
*/
func (a *Client) GetCentresCentreIDVehiclesVehicleID(params *GetCentresCentreIDVehiclesVehicleIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetCentresCentreIDVehiclesVehicleIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCentresCentreIDVehiclesVehicleIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCentresCentreIDVehiclesVehicleID",
		Method:             "GET",
		PathPattern:        "/centres/{centre_id}/vehicles/{vehicle_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCentresCentreIDVehiclesVehicleIDReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCentresCentreIDVehiclesVehicleIDOK), nil
}

/*
GetCentresCentreIDVehiclesVehicleIDDiscounts lists discounts on a vehicle

Request an array of discounts applied to a vehicle at a centre.
*/
func (a *Client) GetCentresCentreIDVehiclesVehicleIDDiscounts(params *GetCentresCentreIDVehiclesVehicleIDDiscountsParams, authInfo runtime.ClientAuthInfoWriter) (*GetCentresCentreIDVehiclesVehicleIDDiscountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCentresCentreIDVehiclesVehicleIDDiscountsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCentresCentreIDVehiclesVehicleIDDiscounts",
		Method:             "GET",
		PathPattern:        "/centres/{centre_id}/vehicles/{vehicle_id}/discounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCentresCentreIDVehiclesVehicleIDDiscountsReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCentresCentreIDVehiclesVehicleIDDiscountsOK), nil
}

/*
GetCentresParkings lists centre parking rates

Request an array of centre parking rates filtered by allowed parameters.
*/
func (a *Client) GetCentresParkings(params *GetCentresParkingsParams, authInfo runtime.ClientAuthInfoWriter) (*GetCentresParkingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCentresParkingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCentresParkings",
		Method:             "GET",
		PathPattern:        "/centres/parkings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCentresParkingsReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCentresParkingsOK), nil
}

/*
GetCentresParkingsParkingID individuals centre parking rates

Request single-centre parking rates.
*/
func (a *Client) GetCentresParkingsParkingID(params *GetCentresParkingsParkingIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetCentresParkingsParkingIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCentresParkingsParkingIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCentresParkingsParkingID",
		Method:             "GET",
		PathPattern:        "/centres/parkings/{parking_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCentresParkingsParkingIDReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCentresParkingsParkingIDOK), nil
}

/*
PatchCentresParkingsParkingID updates centre parking rates

Request update of centre parking rates from JSON data in the request body.
*/
func (a *Client) PatchCentresParkingsParkingID(params *PatchCentresParkingsParkingIDParams, authInfo runtime.ClientAuthInfoWriter) (*PatchCentresParkingsParkingIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchCentresParkingsParkingIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchCentresParkingsParkingID",
		Method:             "PATCH",
		PathPattern:        "/centres/parkings/{parking_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchCentresParkingsParkingIDReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchCentresParkingsParkingIDNoContent), nil
}

/*
PostCentresCentreIDVehiclesVehicleIDDiscounts applies a discount to a vehicle

Request creation of a discount on a vehicle from JSON data in the body.
*/
func (a *Client) PostCentresCentreIDVehiclesVehicleIDDiscounts(params *PostCentresCentreIDVehiclesVehicleIDDiscountsParams, authInfo runtime.ClientAuthInfoWriter) (*PostCentresCentreIDVehiclesVehicleIDDiscountsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCentresCentreIDVehiclesVehicleIDDiscountsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostCentresCentreIDVehiclesVehicleIDDiscounts",
		Method:             "POST",
		PathPattern:        "/centres/{centre_id}/vehicles/{vehicle_id}/discounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostCentresCentreIDVehiclesVehicleIDDiscountsReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostCentresCentreIDVehiclesVehicleIDDiscountsCreated), nil
}

/*
PostCentresParkings creates centre parking rates

Request creation of centre parking rates from JSON data in the request body.
*/
func (a *Client) PostCentresParkings(params *PostCentresParkingsParams, authInfo runtime.ClientAuthInfoWriter) (*PostCentresParkingsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCentresParkingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostCentresParkings",
		Method:             "POST",
		PathPattern:        "/centres/parkings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostCentresParkingsReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostCentresParkingsCreated), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
