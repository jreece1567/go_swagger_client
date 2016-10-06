package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*Color A color

swagger:model color
*/
type Color struct {

	/* code

	Required: true
	Max Length: 6
	Min Length: 6
	*/
	Code *string `json:"code"`

	/* name

	Required: true
	*/
	Name *string `json:"name"`
}

// Validate validates this color
func (m *Color) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
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

func (m *Color) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	if err := validate.MinLength("code", "body", string(*m.Code), 6); err != nil {
		return err
	}

	if err := validate.MaxLength("code", "body", string(*m.Code), 6); err != nil {
		return err
	}

	return nil
}

var colorTypeNamePropEnum []interface{}

// prop value enum
func (m *Color) validateNameEnum(path, location string, value string) error {
	if colorTypeNamePropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["blush","green","lavender","light_blue","med_blue","mint","purple","rose","royal","transparent"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			colorTypeNamePropEnum = append(colorTypeNamePropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, colorTypeNamePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Color) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	// value enum
	if err := m.validateNameEnum("name", "body", *m.Name); err != nil {
		return err
	}

	return nil
}