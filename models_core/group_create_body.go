package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*GroupCreateBody Data to create a category group.

swagger:model groupCreateBody
*/
type GroupCreateBody struct {

	/* Identifiers for categories in this group.

	Required: true
	*/
	CategoryIds []int64 `json:"category_ids"`

	/* Country. One per group.

	Required: true
	*/
	Country *string `json:"country"`

	/* Group name.

	Required: true
	*/
	Name *string `json:"name"`
}

// Validate validates this group create body
func (m *GroupCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategoryIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCountry(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupCreateBody) validateCategoryIds(formats strfmt.Registry) error {

	if err := validate.Required("category_ids", "body", m.CategoryIds); err != nil {
		return err
	}

	return nil
}

var groupCreateBodyTypeCountryPropEnum []interface{}

// prop value enum
func (m *GroupCreateBody) validateCountryEnum(path, location string, value string) error {
	if groupCreateBodyTypeCountryPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["au","nz","uk","us"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			groupCreateBodyTypeCountryPropEnum = append(groupCreateBodyTypeCountryPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, groupCreateBodyTypeCountryPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GroupCreateBody) validateCountry(formats strfmt.Registry) error {

	if err := validate.Required("country", "body", m.Country); err != nil {
		return err
	}

	// value enum
	if err := m.validateCountryEnum("country", "body", *m.Country); err != nil {
		return err
	}

	return nil
}

func (m *GroupCreateBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}
