package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*DiningInstance Dining information for a store.

swagger:model diningInstance
*/
type DiningInstance struct {

	/* links

	Required: true
	*/
	Links *DiningLinks `json:"_links"`

	/* Description of allergen information at this store

	Required: true
	Max Length: 2000
	*/
	AllergenInfo *string `json:"allergen_info"`

	/* Cuisine offered by the store if it provides dining facilities

	Required: true
	*/
	Cuisine *string `json:"cuisine"`

	/* Number indicating the relative price of dining at the store

	Required: true
	Maximum: 4
	*/
	PriceGuide *int64 `json:"price_guide"`

	/* Dining vendor identifier if the store is a dining facility that participates in Open Dining Network; an empty string otherwise.

	Required: true
	*/
	VendorID *string `json:"vendor_id"`
}

// Validate validates this dining instance
func (m *DiningInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAllergenInfo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCuisine(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePriceGuide(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVendorID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DiningInstance) validateLinks(formats strfmt.Registry) error {

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

func (m *DiningInstance) validateAllergenInfo(formats strfmt.Registry) error {

	if err := validate.Required("allergen_info", "body", m.AllergenInfo); err != nil {
		return err
	}

	if err := validate.MaxLength("allergen_info", "body", string(*m.AllergenInfo), 2000); err != nil {
		return err
	}

	return nil
}

func (m *DiningInstance) validateCuisine(formats strfmt.Registry) error {

	if err := validate.Required("cuisine", "body", m.Cuisine); err != nil {
		return err
	}

	return nil
}

func (m *DiningInstance) validatePriceGuide(formats strfmt.Registry) error {

	if err := validate.Required("price_guide", "body", m.PriceGuide); err != nil {
		return err
	}

	if err := validate.MaximumInt("price_guide", "body", int64(*m.PriceGuide), 4, false); err != nil {
		return err
	}

	return nil
}

func (m *DiningInstance) validateVendorID(formats strfmt.Registry) error {

	if err := validate.Required("vendor_id", "body", m.VendorID); err != nil {
		return err
	}

	return nil
}
