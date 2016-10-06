package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*LocaleInstance Locale

swagger:model localeInstance
*/
type LocaleInstance struct {

	/* links

	Required: true
	*/
	Links *LocaleLinks `json:"_links"`

	/* Category identifier. Identifier to retrieve the category.

	Required: true
	*/
	CategoryID *int64 `json:"category_id"`

	/* Locale identifiers. Array of identifiers to order children.

	Required: true
	*/
	ChildrenSortOrder []int64 `json:"children_sort_order"`

	/* Description of locale.

	Required: true
	*/
	Description *string `json:"description"`

	/* Link to the image.
	 */
	Image string `json:"image,omitempty"`

	/* Indexable status.

	Required: true
	*/
	Indexable *bool `json:"indexable"`

	/* Locale name.

	Required: true
	*/
	Name *string `json:"name"`

	/* Locale slug.

	Required: true
	*/
	Slug *string `json:"slug"`

	/* Locale state.

	Required: true
	*/
	State *string `json:"state"`

	/* Unknown.

	Required: true
	*/
	StateEvents []string `json:"state_events"`

	/* status

	Required: true
	*/
	Status *LocaleStatus `json:"status"`

	/* Date-time the locale was last updated.

	Required: true
	*/
	UpdatedAt *strfmt.DateTime `json:"updated_at"`

	/* Locale identifier.

	Required: true
	Pattern: ^[a-z]{2}(_[A-Z]{2})?(\|[a-z]+)?$
	*/
	WestfieldLocale *string `json:"westfield_locale"`
}

// Validate validates this locale instance
func (m *LocaleInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCategoryID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateChildrenSortOrder(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIndexable(formats); err != nil {
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

	if err := m.validateState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStateEvents(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
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

func (m *LocaleInstance) validateLinks(formats strfmt.Registry) error {

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

func (m *LocaleInstance) validateCategoryID(formats strfmt.Registry) error {

	if err := validate.Required("category_id", "body", m.CategoryID); err != nil {
		return err
	}

	return nil
}

func (m *LocaleInstance) validateChildrenSortOrder(formats strfmt.Registry) error {

	if err := validate.Required("children_sort_order", "body", m.ChildrenSortOrder); err != nil {
		return err
	}

	return nil
}

func (m *LocaleInstance) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *LocaleInstance) validateIndexable(formats strfmt.Registry) error {

	if err := validate.Required("indexable", "body", m.Indexable); err != nil {
		return err
	}

	return nil
}

func (m *LocaleInstance) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *LocaleInstance) validateSlug(formats strfmt.Registry) error {

	if err := validate.Required("slug", "body", m.Slug); err != nil {
		return err
	}

	return nil
}

func (m *LocaleInstance) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *LocaleInstance) validateStateEvents(formats strfmt.Registry) error {

	if err := validate.Required("state_events", "body", m.StateEvents); err != nil {
		return err
	}

	return nil
}

func (m *LocaleInstance) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if m.Status != nil {

		if err := m.Status.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *LocaleInstance) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	return nil
}

func (m *LocaleInstance) validateWestfieldLocale(formats strfmt.Registry) error {

	if err := validate.Required("westfield_locale", "body", m.WestfieldLocale); err != nil {
		return err
	}

	if err := validate.Pattern("westfield_locale", "body", string(*m.WestfieldLocale), `^[a-z]{2}(_[A-Z]{2})?(\|[a-z]+)?$`); err != nil {
		return err
	}

	return nil
}
