package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*CommentListResponse List of comments for a deal.

swagger:model commentListResponse
*/
type CommentListResponse struct {

	/* data

	Required: true
	*/
	Data []*CommentInstance `json:"data"`

	/* errors

	Required: true
	*/
	Errors EmptyObject `json:"errors"`

	/* meta

	Required: true
	*/
	Meta *MetaResponse `json:"meta"`
}

// Validate validates this comment list response
func (m *CommentListResponse) Validate(formats strfmt.Registry) error {
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

func (m *CommentListResponse) validateData(formats strfmt.Registry) error {

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

func (m *CommentListResponse) validateErrors(formats strfmt.Registry) error {

	return nil
}

func (m *CommentListResponse) validateMeta(formats strfmt.Registry) error {

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
