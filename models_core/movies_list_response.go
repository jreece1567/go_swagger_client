package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*MoviesListResponse List of movies

swagger:model moviesListResponse
*/
type MoviesListResponse struct {

	/* data

	Required: true
	*/
	Data []*MovieListInstance `json:"data"`

	/* errors

	Required: true
	*/
	Errors EmptyObject `json:"errors"`

	/* meta

	Required: true
	*/
	Meta *MetaResponse `json:"meta"`
}

// Validate validates this movies list response
func (m *MoviesListResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateErrors(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMeta(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MoviesListResponse) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	for i := 0; i < len(m.Data); i++ {

		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {

			if err := m.Data[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *MoviesListResponse) validateErrors(formats strfmt.Registry) error {

	return nil
}

func (m *MoviesListResponse) validateMeta(formats strfmt.Registry) error {

	if err := validate.Required("meta", "body", m.Meta); err != nil {
		return err
	}

	if m.Meta != nil {

		if err := m.Meta.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
