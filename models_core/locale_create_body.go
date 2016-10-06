package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*LocaleCreateBody Data to create a locale.

swagger:model localeCreateBody
*/
type LocaleCreateBody struct {

	/* Category identifier. Identifier to retrieve the category.

	Required: true
	*/
	CategoryID *int64 `json:"category_id"`

	/* Name of locale.

	Required: true
	*/
	Name *string `json:"name"`

	/* Slug of locale.

	Required: true
	*/
	Slug *string `json:"slug"`

	/* Westfield locale.

	Required: true
	Pattern: ^[a-z]{2}(_[A-Z]{2})?(\|[a-z]+)?$
	*/
	WestfieldLocale *string `json:"westfield_locale"`
}

// Validate validates this locale create body
func (m *LocaleCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategoryID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSlug(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateWestfieldLocale(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LocaleCreateBody) validateCategoryID(formats strfmt.Registry) error {

	if err := validate.Required("category_id", "body", m.CategoryID); err != nil {
		return err
	}

	return nil
}

func (m *LocaleCreateBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *LocaleCreateBody) validateSlug(formats strfmt.Registry) error {

	if err := validate.Required("slug", "body", m.Slug); err != nil {
		return err
	}

	return nil
}

func (m *LocaleCreateBody) validateWestfieldLocale(formats strfmt.Registry) error {

	if err := validate.Required("westfield_locale", "body", m.WestfieldLocale); err != nil {
		return err
	}

	if err := validate.Pattern("westfield_locale", "body", string(*m.WestfieldLocale), `^[a-z]{2}(_[A-Z]{2})?(\|[a-z]+)?$`); err != nil {
		return err
	}

	return nil
}