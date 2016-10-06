package s_e_a_r_c_h

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new s e a r c h API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for s e a r c h API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetSearch searches all content types

Search results from all APIs.
*/
func (a *Client) GetSearch(params *GetSearchParams, authInfo runtime.ClientAuthInfoWriter) (*GetSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSearch",
		Method:             "GET",
		PathPattern:        "/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSearchReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSearchOK), nil
}

/*
GetSearchDeals searches deals

Search content of the 'deals' content type
*/
func (a *Client) GetSearchDeals(params *GetSearchDealsParams, authInfo runtime.ClientAuthInfoWriter) (*GetSearchDealsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSearchDealsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSearchDeals",
		Method:             "GET",
		PathPattern:        "/search/deals",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSearchDealsReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSearchDealsOK), nil
}

/*
GetSearchEvents searches events

Search content of the 'events' content type
*/
func (a *Client) GetSearchEvents(params *GetSearchEventsParams, authInfo runtime.ClientAuthInfoWriter) (*GetSearchEventsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSearchEventsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSearchEvents",
		Method:             "GET",
		PathPattern:        "/search/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSearchEventsReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSearchEventsOK), nil
}

/*
GetSearchMovies searches movies

Search content of the 'movies' content type
*/
func (a *Client) GetSearchMovies(params *GetSearchMoviesParams, authInfo runtime.ClientAuthInfoWriter) (*GetSearchMoviesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSearchMoviesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSearchMovies",
		Method:             "GET",
		PathPattern:        "/search/movies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSearchMoviesReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSearchMoviesOK), nil
}

/*
GetSearchNotices searches notices

Search content of the 'notices' content type
*/
func (a *Client) GetSearchNotices(params *GetSearchNoticesParams, authInfo runtime.ClientAuthInfoWriter) (*GetSearchNoticesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSearchNoticesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSearchNotices",
		Method:             "GET",
		PathPattern:        "/search/notices",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSearchNoticesReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSearchNoticesOK), nil
}

/*
GetSearchProductcurations searches productcurations

Search content of the 'productcurations' content type
*/
func (a *Client) GetSearchProductcurations(params *GetSearchProductcurationsParams, authInfo runtime.ClientAuthInfoWriter) (*GetSearchProductcurationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSearchProductcurationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSearchProductcurations",
		Method:             "GET",
		PathPattern:        "/search/productcurations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSearchProductcurationsReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSearchProductcurationsOK), nil
}

/*
GetSearchProductcurationsCountry searches productcurations for a specific country

Search content of the 'productcurations' content type for a specific country
*/
func (a *Client) GetSearchProductcurationsCountry(params *GetSearchProductcurationsCountryParams, authInfo runtime.ClientAuthInfoWriter) (*GetSearchProductcurationsCountryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSearchProductcurationsCountryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSearchProductcurationsCountry",
		Method:             "GET",
		PathPattern:        "/search/productcurations/{country}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSearchProductcurationsCountryReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSearchProductcurationsCountryOK), nil
}

/*
GetSearchProducts searches products

Search content of the 'products' content type
*/
func (a *Client) GetSearchProducts(params *GetSearchProductsParams, authInfo runtime.ClientAuthInfoWriter) (*GetSearchProductsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSearchProductsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSearchProducts",
		Method:             "GET",
		PathPattern:        "/search/products",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSearchProductsReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSearchProductsOK), nil
}

/*
GetSearchProductsCountry searches products for a specific country

Search content of the 'products' content type for a specific country
*/
func (a *Client) GetSearchProductsCountry(params *GetSearchProductsCountryParams, authInfo runtime.ClientAuthInfoWriter) (*GetSearchProductsCountryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSearchProductsCountryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSearchProductsCountry",
		Method:             "GET",
		PathPattern:        "/search/products/{country}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSearchProductsCountryReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSearchProductsCountryOK), nil
}

/*
GetSearchProductsCountrySyndicated searches products that syndicate

Search syndicated content of the 'products' content type
*/
func (a *Client) GetSearchProductsCountrySyndicated(params *GetSearchProductsCountrySyndicatedParams, authInfo runtime.ClientAuthInfoWriter) (*GetSearchProductsCountrySyndicatedOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSearchProductsCountrySyndicatedParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSearchProductsCountrySyndicated",
		Method:             "GET",
		PathPattern:        "/search/products/{country}/syndicated",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSearchProductsCountrySyndicatedReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSearchProductsCountrySyndicatedOK), nil
}

/*
GetSearchServices searches services

Search content of the 'services' content type
*/
func (a *Client) GetSearchServices(params *GetSearchServicesParams, authInfo runtime.ClientAuthInfoWriter) (*GetSearchServicesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSearchServicesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSearchServices",
		Method:             "GET",
		PathPattern:        "/search/services",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSearchServicesReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSearchServicesOK), nil
}

/*
GetSearchStores searches stores

Search content of the 'stores' content type
*/
func (a *Client) GetSearchStores(params *GetSearchStoresParams, authInfo runtime.ClientAuthInfoWriter) (*GetSearchStoresOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSearchStoresParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetSearchStores",
		Method:             "GET",
		PathPattern:        "/search/stores",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSearchStoresReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSearchStoresOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
