package c_e_n_t_r_e_i_n_f_o_r_m_a_t_i_o_n

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new c e n t r e i n f o r m a t i o n API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for c e n t r e i n f o r m a t i o n API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetCentres lists centres

Request an array of centres filtered by allowed parameters.
*/
func (a *Client) GetCentres(params *GetCentresParams, authInfo runtime.ClientAuthInfoWriter) (*GetCentresOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCentresParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCentres",
		Method:             "GET",
		PathPattern:        "/centres",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCentresReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCentresOK), nil
}

/*
GetCentresCentreID individuals centre

Request a single centre.
*/
func (a *Client) GetCentresCentreID(params *GetCentresCentreIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetCentresCentreIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCentresCentreIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCentresCentreID",
		Method:             "GET",
		PathPattern:        "/centres/{centre_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCentresCentreIDReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCentresCentreIDOK), nil
}

/*
GetCentresNewsletters lists newsletters

List of newsletters associated with each centre.
*/
func (a *Client) GetCentresNewsletters(params *GetCentresNewslettersParams, authInfo runtime.ClientAuthInfoWriter) (*GetCentresNewslettersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCentresNewslettersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCentresNewsletters",
		Method:             "GET",
		PathPattern:        "/centres/newsletters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCentresNewslettersReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCentresNewslettersOK), nil
}

/*
GetCentresNewslettersInfo information about a list of newsletters

Request the centre_ids and countries of a list of newsletters
*/
func (a *Client) GetCentresNewslettersInfo(params *GetCentresNewslettersInfoParams, authInfo runtime.ClientAuthInfoWriter) (*GetCentresNewslettersInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCentresNewslettersInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCentresNewslettersInfo",
		Method:             "GET",
		PathPattern:        "/centres/newsletters/info",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCentresNewslettersInfoReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCentresNewslettersInfoOK), nil
}

/*
GetCountries lists countries

Request an array of countries.
*/
func (a *Client) GetCountries(params *GetCountriesParams, authInfo runtime.ClientAuthInfoWriter) (*GetCountriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCountriesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCountries",
		Method:             "GET",
		PathPattern:        "/countries",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCountriesReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCountriesOK), nil
}

/*
GetStates lists states filtered by allowed parameters

Returns a list of states
*/
func (a *Client) GetStates(params *GetStatesParams, authInfo runtime.ClientAuthInfoWriter) (*GetStatesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStatesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStates",
		Method:             "GET",
		PathPattern:        "/states",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStatesReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetStatesOK), nil
}

/*
PatchCentresCentreID updates a centre

Request update of a centre from JSON data in the request body.
*/
func (a *Client) PatchCentresCentreID(params *PatchCentresCentreIDParams, authInfo runtime.ClientAuthInfoWriter) (*PatchCentresCentreIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchCentresCentreIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchCentresCentreID",
		Method:             "PATCH",
		PathPattern:        "/centres/{centre_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchCentresCentreIDReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchCentresCentreIDNoContent), nil
}

/*
PostCentres creates a centre

Request creation of a centre from JSON data in the request body.
*/
func (a *Client) PostCentres(params *PostCentresParams, authInfo runtime.ClientAuthInfoWriter) (*PostCentresCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCentresParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostCentres",
		Method:             "POST",
		PathPattern:        "/centres",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostCentresReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostCentresCreated), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
