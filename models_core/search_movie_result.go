package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*SearchMovieResult search movie result

swagger:model SearchMovieResult
*/
type SearchMovieResult struct {

	/* Is the facet count accurate?
	 */
	ExhaustiveFacetsCount bool `json:"exhaustiveFacetsCount,omitempty"`

	/* The facets available to filter by

	Required: true
	*/
	Facets interface{} `json:"facets"`

	/* statistics about facets
	 */
	FacetsStats interface{} `json:"facets_stats,omitempty"`

	/* An array of items that match this search

	Required: true
	*/
	Hits []*MovieListInstance `json:"hits"`

	/* The number of results per page

	Required: true
	*/
	HitsPerPage *int64 `json:"hitsPerPage"`

	/* The type of results contained in the hits key

	Required: true
	*/
	Index *string `json:"index"`

	/* Total number of results for this type

	Required: true
	*/
	NbHits *int64 `json:"nbHits"`

	/* How many pages there are in total

	Required: true
	*/
	NbPages *int64 `json:"nbPages"`

	/* Which page this response is on

	Required: true
	*/
	Page *int64 `json:"page"`

	/* The query as sent to the search backend

	Required: true
	*/
	Params *string `json:"params"`

	/* How long this spent being processed on the backend

	Required: true
	*/
	ProcessingTimeMS *int64 `json:"processingTimeMS"`

	/* The search term queried for

	Required: true
	*/
	Query *string `json:"query"`
}

// Validate validates this search movie result
func (m *SearchMovieResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFacets(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHits(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHitsPerPage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIndex(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNbHits(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNbPages(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateParams(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProcessingTimeMS(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateQuery(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchMovieResult) validateFacets(formats strfmt.Registry) error {

	return nil
}

func (m *SearchMovieResult) validateHits(formats strfmt.Registry) error {

	if err := validate.Required("hits", "body", m.Hits); err != nil {
		return err
	}

	for i := 0; i < len(m.Hits); i++ {

		if swag.IsZero(m.Hits[i]) { // not required
			continue
		}

		if m.Hits[i] != nil {

			if err := m.Hits[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *SearchMovieResult) validateHitsPerPage(formats strfmt.Registry) error {

	if err := validate.Required("hitsPerPage", "body", m.HitsPerPage); err != nil {
		return err
	}

	return nil
}

func (m *SearchMovieResult) validateIndex(formats strfmt.Registry) error {

	if err := validate.Required("index", "body", m.Index); err != nil {
		return err
	}

	return nil
}

func (m *SearchMovieResult) validateNbHits(formats strfmt.Registry) error {

	if err := validate.Required("nbHits", "body", m.NbHits); err != nil {
		return err
	}

	return nil
}

func (m *SearchMovieResult) validateNbPages(formats strfmt.Registry) error {

	if err := validate.Required("nbPages", "body", m.NbPages); err != nil {
		return err
	}

	return nil
}

func (m *SearchMovieResult) validatePage(formats strfmt.Registry) error {

	if err := validate.Required("page", "body", m.Page); err != nil {
		return err
	}

	return nil
}

func (m *SearchMovieResult) validateParams(formats strfmt.Registry) error {

	if err := validate.Required("params", "body", m.Params); err != nil {
		return err
	}

	return nil
}

func (m *SearchMovieResult) validateProcessingTimeMS(formats strfmt.Registry) error {

	if err := validate.Required("processingTimeMS", "body", m.ProcessingTimeMS); err != nil {
		return err
	}

	return nil
}

func (m *SearchMovieResult) validateQuery(formats strfmt.Registry) error {

	if err := validate.Required("query", "body", m.Query); err != nil {
		return err
	}

	return nil
}
