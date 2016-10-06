package s_t_o_r_e_i_n_f_o_r_m_a_t_i_o_n

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new s t o r e i n f o r m a t i o n API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for s t o r e i n f o r m a t i o n API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetStores lists stores

Request an array of stores filtered by allowed parameters.
*/
func (a *Client) GetStores(params *GetStoresParams, authInfo runtime.ClientAuthInfoWriter) (*GetStoresOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStoresParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStores",
		Method:             "GET",
		PathPattern:        "/stores",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStoresReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetStoresOK), nil
}

/*
GetStoresStoreID individuals store

Request a single store.
*/
func (a *Client) GetStoresStoreID(params *GetStoresStoreIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetStoresStoreIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStoresStoreIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStoresStoreID",
		Method:             "GET",
		PathPattern:        "/stores/{store_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStoresStoreIDReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetStoresStoreIDOK), nil
}

/*
PatchStoresStoreID updates a store

Request update of a store from JSON data in the request body.
*/
func (a *Client) PatchStoresStoreID(params *PatchStoresStoreIDParams, authInfo runtime.ClientAuthInfoWriter) (*PatchStoresStoreIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchStoresStoreIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchStoresStoreID",
		Method:             "PATCH",
		PathPattern:        "/stores/{store_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchStoresStoreIDReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchStoresStoreIDNoContent), nil
}

/*
PatchStoresStoreIDLocationsLocationID updates a store location

Request update of a store location from JSON data in the request body.
*/
func (a *Client) PatchStoresStoreIDLocationsLocationID(params *PatchStoresStoreIDLocationsLocationIDParams, authInfo runtime.ClientAuthInfoWriter) (*PatchStoresStoreIDLocationsLocationIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchStoresStoreIDLocationsLocationIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchStoresStoreIDLocationsLocationID",
		Method:             "PATCH",
		PathPattern:        "/stores/{store_id}/locations/{location_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchStoresStoreIDLocationsLocationIDReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchStoresStoreIDLocationsLocationIDNoContent), nil
}

/*
PostStores creates a store

Request creation of a store from JSON data in the request body.
*/
func (a *Client) PostStores(params *PostStoresParams, authInfo runtime.ClientAuthInfoWriter) (*PostStoresCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostStoresParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostStores",
		Method:             "POST",
		PathPattern:        "/stores",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostStoresReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostStoresCreated), nil
}

/*
PostStoresStoreIDLocations creates a store location

Request creation of a store location from JSON data in the request body.
*/
func (a *Client) PostStoresStoreIDLocations(params *PostStoresStoreIDLocationsParams, authInfo runtime.ClientAuthInfoWriter) (*PostStoresStoreIDLocationsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostStoresStoreIDLocationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostStoresStoreIDLocations",
		Method:             "POST",
		PathPattern:        "/stores/{store_id}/locations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostStoresStoreIDLocationsReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostStoresStoreIDLocationsCreated), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}