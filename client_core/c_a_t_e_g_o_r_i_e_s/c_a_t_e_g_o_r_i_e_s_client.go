package c_a_t_e_g_o_r_i_e_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new c a t e g o r i e s API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for c a t e g o r i e s API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteCategoriesCategoryID deletes a category

Request deletion of a category.
*/
func (a *Client) DeleteCategoriesCategoryID(params *DeleteCategoriesCategoryIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteCategoriesCategoryIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCategoriesCategoryIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteCategoriesCategoryID",
		Method:             "DELETE",
		PathPattern:        "/categories/{category_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteCategoriesCategoryIDReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteCategoriesCategoryIDNoContent), nil
}

/*
DeleteCategoriesCategoryIDLocalesWestfieldLocale deletes a locale

Request deletion of a locale.
*/
func (a *Client) DeleteCategoriesCategoryIDLocalesWestfieldLocale(params *DeleteCategoriesCategoryIDLocalesWestfieldLocaleParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteCategoriesCategoryIDLocalesWestfieldLocaleNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCategoriesCategoryIDLocalesWestfieldLocaleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteCategoriesCategoryIDLocalesWestfieldLocale",
		Method:             "DELETE",
		PathPattern:        "/categories/{category_id}/locales/{westfield_locale}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteCategoriesCategoryIDLocalesWestfieldLocaleReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteCategoriesCategoryIDLocalesWestfieldLocaleNoContent), nil
}

/*
DeleteCategoriesGroupsGroupID deletes a group

Request deletion of a group.
*/
func (a *Client) DeleteCategoriesGroupsGroupID(params *DeleteCategoriesGroupsGroupIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteCategoriesGroupsGroupIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCategoriesGroupsGroupIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteCategoriesGroupsGroupID",
		Method:             "DELETE",
		PathPattern:        "/categories/groups/{group_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteCategoriesGroupsGroupIDReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteCategoriesGroupsGroupIDNoContent), nil
}

/*
GetCategories lists categories

Request an array of categories.
*/
func (a *Client) GetCategories(params *GetCategoriesParams, authInfo runtime.ClientAuthInfoWriter) (*GetCategoriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCategoriesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCategories",
		Method:             "GET",
		PathPattern:        "/categories",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCategoriesReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCategoriesOK), nil
}

/*
GetCategoriesCategoryID individuals category

Request a single category.
*/
func (a *Client) GetCategoriesCategoryID(params *GetCategoriesCategoryIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetCategoriesCategoryIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCategoriesCategoryIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCategoriesCategoryID",
		Method:             "GET",
		PathPattern:        "/categories/{category_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCategoriesCategoryIDReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCategoriesCategoryIDOK), nil
}

/*
GetCategoriesCategoryIDLocales lists category locales

Request an array of category locales.
*/
func (a *Client) GetCategoriesCategoryIDLocales(params *GetCategoriesCategoryIDLocalesParams, authInfo runtime.ClientAuthInfoWriter) (*GetCategoriesCategoryIDLocalesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCategoriesCategoryIDLocalesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCategoriesCategoryIDLocales",
		Method:             "GET",
		PathPattern:        "/categories/{category_id}/locales",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCategoriesCategoryIDLocalesReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCategoriesCategoryIDLocalesOK), nil
}

/*
GetCategoriesCategoryIDLocalesWestfieldLocale individuals category locale

Request a single category locale.
*/
func (a *Client) GetCategoriesCategoryIDLocalesWestfieldLocale(params *GetCategoriesCategoryIDLocalesWestfieldLocaleParams, authInfo runtime.ClientAuthInfoWriter) (*GetCategoriesCategoryIDLocalesWestfieldLocaleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCategoriesCategoryIDLocalesWestfieldLocaleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCategoriesCategoryIDLocalesWestfieldLocale",
		Method:             "GET",
		PathPattern:        "/categories/{category_id}/locales/{westfield_locale}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCategoriesCategoryIDLocalesWestfieldLocaleReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCategoriesCategoryIDLocalesWestfieldLocaleOK), nil
}

/*
GetCategoriesCategoryIDLocalesWestfieldLocaleChildrenList categories list children

Request the children list for a single category and locale.
*/
func (a *Client) GetCategoriesCategoryIDLocalesWestfieldLocaleChildrenList(params *GetCategoriesCategoryIDLocalesWestfieldLocaleChildrenListParams, authInfo runtime.ClientAuthInfoWriter) (*GetCategoriesCategoryIDLocalesWestfieldLocaleChildrenListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCategoriesCategoryIDLocalesWestfieldLocaleChildrenListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCategoriesCategoryIDLocalesWestfieldLocaleChildrenList",
		Method:             "GET",
		PathPattern:        "/categories/{category_id}/locales/{westfield_locale}/children-list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCategoriesCategoryIDLocalesWestfieldLocaleChildrenListReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCategoriesCategoryIDLocalesWestfieldLocaleChildrenListOK), nil
}

/*
GetCategoriesCategoryIDLocalesWestfieldLocaleChildrenTree categories tree

Request the children tree for a single category and locale.
*/
func (a *Client) GetCategoriesCategoryIDLocalesWestfieldLocaleChildrenTree(params *GetCategoriesCategoryIDLocalesWestfieldLocaleChildrenTreeParams, authInfo runtime.ClientAuthInfoWriter) (*GetCategoriesCategoryIDLocalesWestfieldLocaleChildrenTreeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCategoriesCategoryIDLocalesWestfieldLocaleChildrenTreeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCategoriesCategoryIDLocalesWestfieldLocaleChildrenTree",
		Method:             "GET",
		PathPattern:        "/categories/{category_id}/locales/{westfield_locale}/children-tree",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCategoriesCategoryIDLocalesWestfieldLocaleChildrenTreeReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCategoriesCategoryIDLocalesWestfieldLocaleChildrenTreeOK), nil
}

/*
GetCategoriesCategoryIDLocalesWestfieldLocaleParentList categories list parents

Request all parents for a single category as a list. The list is ordered from the top parent down to the category.
*/
func (a *Client) GetCategoriesCategoryIDLocalesWestfieldLocaleParentList(params *GetCategoriesCategoryIDLocalesWestfieldLocaleParentListParams, authInfo runtime.ClientAuthInfoWriter) (*GetCategoriesCategoryIDLocalesWestfieldLocaleParentListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCategoriesCategoryIDLocalesWestfieldLocaleParentListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCategoriesCategoryIDLocalesWestfieldLocaleParentList",
		Method:             "GET",
		PathPattern:        "/categories/{category_id}/locales/{westfield_locale}/parent-list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCategoriesCategoryIDLocalesWestfieldLocaleParentListReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCategoriesCategoryIDLocalesWestfieldLocaleParentListOK), nil
}

/*
GetCategoriesGroups lists groups

Request a category group
*/
func (a *Client) GetCategoriesGroups(params *GetCategoriesGroupsParams, authInfo runtime.ClientAuthInfoWriter) (*GetCategoriesGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCategoriesGroupsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCategoriesGroups",
		Method:             "GET",
		PathPattern:        "/categories/groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCategoriesGroupsReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCategoriesGroupsOK), nil
}

/*
GetCategoriesGroupsGroupID individuals category group

Request a single category grou.
*/
func (a *Client) GetCategoriesGroupsGroupID(params *GetCategoriesGroupsGroupIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetCategoriesGroupsGroupIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCategoriesGroupsGroupIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCategoriesGroupsGroupID",
		Method:             "GET",
		PathPattern:        "/categories/groups/{group_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCategoriesGroupsGroupIDReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCategoriesGroupsGroupIDOK), nil
}

/*
GetCategoriesLocalesWestfieldLocale gets locales for a list of categories

Get list of locales for a list of categories.
*/
func (a *Client) GetCategoriesLocalesWestfieldLocale(params *GetCategoriesLocalesWestfieldLocaleParams, authInfo runtime.ClientAuthInfoWriter) (*GetCategoriesLocalesWestfieldLocaleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCategoriesLocalesWestfieldLocaleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCategoriesLocalesWestfieldLocale",
		Method:             "GET",
		PathPattern:        "/categories/locales/{westfield_locale}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCategoriesLocalesWestfieldLocaleReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCategoriesLocalesWestfieldLocaleOK), nil
}

/*
PatchCategoriesCategoryID updates a category

Request update of a category from JSON data in the request body.
*/
func (a *Client) PatchCategoriesCategoryID(params *PatchCategoriesCategoryIDParams, authInfo runtime.ClientAuthInfoWriter) (*PatchCategoriesCategoryIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchCategoriesCategoryIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchCategoriesCategoryID",
		Method:             "PATCH",
		PathPattern:        "/categories/{category_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchCategoriesCategoryIDReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchCategoriesCategoryIDNoContent), nil
}

/*
PatchCategoriesCategoryIDLocalesWestfieldLocale updates a locale

Request update of a category locale from JSON data in the request body.
*/
func (a *Client) PatchCategoriesCategoryIDLocalesWestfieldLocale(params *PatchCategoriesCategoryIDLocalesWestfieldLocaleParams, authInfo runtime.ClientAuthInfoWriter) (*PatchCategoriesCategoryIDLocalesWestfieldLocaleNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchCategoriesCategoryIDLocalesWestfieldLocaleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchCategoriesCategoryIDLocalesWestfieldLocale",
		Method:             "PATCH",
		PathPattern:        "/categories/{category_id}/locales/{westfield_locale}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchCategoriesCategoryIDLocalesWestfieldLocaleReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchCategoriesCategoryIDLocalesWestfieldLocaleNoContent), nil
}

/*
PatchCategoriesGroupsGroupID updates a group

Request update of a group from JSON data in the request body.
*/
func (a *Client) PatchCategoriesGroupsGroupID(params *PatchCategoriesGroupsGroupIDParams, authInfo runtime.ClientAuthInfoWriter) (*PatchCategoriesGroupsGroupIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchCategoriesGroupsGroupIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchCategoriesGroupsGroupID",
		Method:             "PATCH",
		PathPattern:        "/categories/groups/{group_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchCategoriesGroupsGroupIDReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchCategoriesGroupsGroupIDNoContent), nil
}

/*
PostCategories creates a category

Request creation of a category from JSON data in the request body.
*/
func (a *Client) PostCategories(params *PostCategoriesParams, authInfo runtime.ClientAuthInfoWriter) (*PostCategoriesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCategoriesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostCategories",
		Method:             "POST",
		PathPattern:        "/categories",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostCategoriesReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostCategoriesCreated), nil
}

/*
PostCategoriesCategoryIDLocales creates a locale for category

Request creation of a locale for category from JSON data in the request body.
*/
func (a *Client) PostCategoriesCategoryIDLocales(params *PostCategoriesCategoryIDLocalesParams, authInfo runtime.ClientAuthInfoWriter) (*PostCategoriesCategoryIDLocalesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCategoriesCategoryIDLocalesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostCategoriesCategoryIDLocales",
		Method:             "POST",
		PathPattern:        "/categories/{category_id}/locales",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostCategoriesCategoryIDLocalesReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostCategoriesCategoryIDLocalesCreated), nil
}

/*
PostCategoriesGroups creates a group

Request creation of a category group from JSON data in the request body.
*/
func (a *Client) PostCategoriesGroups(params *PostCategoriesGroupsParams, authInfo runtime.ClientAuthInfoWriter) (*PostCategoriesGroupsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCategoriesGroupsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostCategoriesGroups",
		Method:             "POST",
		PathPattern:        "/categories/groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostCategoriesGroupsReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostCategoriesGroupsCreated), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}