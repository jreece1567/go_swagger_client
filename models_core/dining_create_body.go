package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*DiningCreateBody Parameters for the dining section in a store

swagger:model diningCreateBody
*/
type DiningCreateBody struct {

	/* Description of allergen information at this store

	Required: true
	Max Length: 2000
	*/
	AllergenInfo *string `json:"allergen_info"`

	/* Cuisine offered by the store

	Required: true
	*/
	Cuisine *string `json:"cuisine"`

	/* Storage provider identifier of the dining image for the store

	Required: true
	*/
	ImageRef *string `json:"image_ref"`

	/* Storage provider identifier of the menu for the store

	Required: true
	*/
	MenuRef *string `json:"menu_ref"`

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

// Validate validates this dining create body
func (m *DiningCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllergenInfo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCuisine(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateImageRef(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMenuRef(formats); err != nil {
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

func (m *DiningCreateBody) validateAllergenInfo(formats strfmt.Registry) error {

	if err := validate.Required("allergen_info", "body", m.AllergenInfo); err != nil {
		return err
	}

	if err := validate.MaxLength("allergen_info", "body", string(*m.AllergenInfo), 2000); err != nil {
		return err
	}

	return nil
}

func (m *DiningCreateBody) validateCuisine(formats strfmt.Registry) error {

	if err := validate.Required("cuisine", "body", m.Cuisine); err != nil {
		return err
	}

	return nil
}

func (m *DiningCreateBody) validateImageRef(formats strfmt.Registry) error {

	if err := validate.Required("image_ref", "body", m.ImageRef); err != nil {
		return err
	}

	return nil
}

func (m *DiningCreateBody) validateMenuRef(formats strfmt.Registry) error {

	if err := validate.Required("menu_ref", "body", m.MenuRef); err != nil {
		return err
	}

	return nil
}

func (m *DiningCreateBody) validatePriceGuide(formats strfmt.Registry) error {

	if err := validate.Required("price_guide", "body", m.PriceGuide); err != nil {
		return err
	}

	if err := validate.MaximumInt("price_guide", "body", int64(*m.PriceGuide), 4, false); err != nil {
		return err
	}

	return nil
}

func (m *DiningCreateBody) validateVendorID(formats strfmt.Registry) error {

	if err := validate.Required("vendor_id", "body", m.VendorID); err != nil {
		return err
	}

	return nil
}
