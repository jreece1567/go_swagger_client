package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*GroupInstance Category group.

swagger:model groupInstance
*/
type GroupInstance struct {

	/* links

	Required: true
	*/
	Links *GroupLinks `json:"_links"`

	/* Identifiers for categories in this group.

	Required: true
	*/
	CategoryIds []int64 `json:"category_ids"`

	/* Country.

	Required: true
	*/
	Country *string `json:"country"`

	/* Group identifier. Identifier to retrieve this group.

	Required: true
	*/
	GroupID *int64 `json:"group_id"`

	/* Group name.

	Required: true
	*/
	Name *string `json:"name"`

	/* Date-time the locale was last updated.

	Required: true
	*/
	UpdatedAt *strfmt.DateTime `json:"updated_at"`
}

// Validate validates this group instance
func (m *GroupInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCategoryIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCountry(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGroupID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupInstance) validateLinks(formats strfmt.Registry) error {

	if err := validate.Required("_links", "body", m.Links); err != nil {
		return err
	}

	if m.Links != nil {

		if err := m.Links.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *GroupInstance) validateCategoryIds(formats strfmt.Registry) error {

	if err := validate.Required("category_ids", "body", m.CategoryIds); err != nil {
		return err
	}

	return nil
}

func (m *GroupInstance) validateCountry(formats strfmt.Registry) error {

	if err := validate.Required("country", "body", m.Country); err != nil {
		return err
	}

	return nil
}

func (m *GroupInstance) validateGroupID(formats strfmt.Registry) error {

	if err := validate.Required("group_id", "body", m.GroupID); err != nil {
		return err
	}

	return nil
}

func (m *GroupInstance) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *GroupInstance) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	return nil
}