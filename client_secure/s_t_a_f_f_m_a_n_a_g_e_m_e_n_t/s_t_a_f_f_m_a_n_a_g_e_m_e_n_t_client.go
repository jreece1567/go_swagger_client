package s_t_a_f_f_m_a_n_a_g_e_m_e_n_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new s t a f f m a n a g e m e n t API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for s t a f f m a n a g e m e n t API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetStaff staffs with access to westfield apis

Returns a list of staff limited to 100 enabled staff by default
*/
func (a *Client) GetStaff(params *GetStaffParams, authInfo runtime.ClientAuthInfoWriter) (*GetStaffOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStaffParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStaff",
		Method:             "GET",
		PathPattern:        "/staff",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStaffReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetStaffOK), nil
}

/*
GetStaffUUID as specific staff

Returns a staff
*/
func (a *Client) GetStaffUUID(params *GetStaffUUIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetStaffUUIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStaffUUIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetStaffUUID",
		Method:             "GET",
		PathPattern:        "/staff/{uuid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStaffUUIDReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetStaffUUIDOK), nil
}

/*
PostStaff creates new staff
*/
func (a *Client) PostStaff(params *PostStaffParams, authInfo runtime.ClientAuthInfoWriter) (*PostStaffCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostStaffParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostStaff",
		Method:             "POST",
		PathPattern:        "/staff",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostStaffReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostStaffCreated), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
