package models_secure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*CursoredMetaResponse Response metadata.

swagger:model cursoredMetaResponse
*/
type CursoredMetaResponse struct {

	/* API version. Version number for the API.

	Required: true
	*/
	APIVersion *string `json:"api_version"`

	/* End-of-life information about the deprecation of the endpoint. If this has content you must act on it.

	Required: true
	*/
	DeprecationInformation *DeprecationInformation `json:"deprecation_information"`

	/* Next page identifier to retrieve the next page of results.
	 */
	NextPageID string `json:"next_page_id,omitempty"`

	/* Total number of results for all pages.
	 */
	Total int64 `json:"total,omitempty"`
}

// Validate validates this cursored meta response
func (m *CursoredMetaResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIVersion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDeprecationInformation(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CursoredMetaResponse) validateAPIVersion(formats strfmt.Registry) error {

	if err := validate.Required("api_version", "body", m.APIVersion); err != nil {
		return err
	}

	return nil
}

func (m *CursoredMetaResponse) validateDeprecationInformation(formats strfmt.Registry) error {

	if err := validate.Required("deprecation_information", "body", m.DeprecationInformation); err != nil {
		return err
	}

	if m.DeprecationInformation != nil {

		if err := m.DeprecationInformation.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
