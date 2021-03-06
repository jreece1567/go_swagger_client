package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*RetailerInstance A retailer that has stores in Westfield centres.

swagger:model retailerInstance
*/
type RetailerInstance struct {

	/* links

	Required: true
	*/
	Links *RetailerLinks `json:"_links"`

	/* List of affiliate fields

	Required: true
	*/
	AffiliateFields []*RetailerAffiliateField `json:"affiliate_fields"`

	/* List of Category identifiers associated with the retailer

	Required: true
	*/
	CategoryIds []int64 `json:"category_ids"`

	/* Two letter code of the country the retailer operates in

	Required: true
	*/
	Country *string `json:"country"`

	/* Date and time the retailer was marked as deleted, in `zulu-time` format
	 */
	DeletedAt strfmt.DateTime `json:"deleted_at,omitempty"`

	/* E-mail address of the retailer

	Required: true
	*/
	EmailAddress *string `json:"email_address"`

	/* Flag indicating if the retailer is enabled in the Westfield system

	Required: true
	*/
	Enabled *bool `json:"enabled"`

	/* List of features the retailer has. If 'products' is present the retailer has products on the website. If 'can_syndicate' is present the product data for this retailer can be redistributed by a third party.

	Required: true
	*/
	Features []string `json:"features"`

	/* Identifier of the retailer's lease with Westfield

	Required: true
	*/
	LeaseID *string `json:"lease_id"`

	/* Name of the retailer

	Required: true
	*/
	Name *string `json:"name"`

	/* Tiers for grouping similar retailers

	Required: true
	*/
	ProductTier *int64 `json:"product_tier"`

	/* URL-friendly code for the retailer

	Required: true
	*/
	RetailerCode *string `json:"retailer_code"`

	/* Retailer identifier. Identifier to retrieve this retailer.

	Required: true
	*/
	RetailerID *int64 `json:"retailer_id"`

	/* URL of the retailer's website

	Required: true
	*/
	RetailerWebsite *string `json:"retailer_website"`

	/* List of Service Category identifiers associated with the retailer

	Required: true
	*/
	ServiceCategoryIds []int64 `json:"service_category_ids"`

	/* Long description of the retailer

	Required: true
	Max Length: 1000
	*/
	StoreProfile *string `json:"store_profile"`

	/* Date and time the retailer was last updated, in `zulu-time` format

	Required: true
	*/
	UpdatedAt *strfmt.DateTime `json:"updated_at"`
}

// Validate validates this retailer instance
func (m *RetailerInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAffiliateFields(formats); err != nil {
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

	if err := m.validateEmailAddress(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFeatures(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLeaseID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProductTier(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRetailerCode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRetailerID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRetailerWebsite(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateServiceCategoryIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStoreProfile(formats); err != nil {
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

func (m *RetailerInstance) validateLinks(formats strfmt.Registry) error {

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

func (m *RetailerInstance) validateAffiliateFields(formats strfmt.Registry) error {

	if err := validate.Required("affiliate_fields", "body", m.AffiliateFields); err != nil {
		return err
	}

	for i := 0; i < len(m.AffiliateFields); i++ {

		if swag.IsZero(m.AffiliateFields[i]) { // not required
			continue
		}

		if m.AffiliateFields[i] != nil {

			if err := m.AffiliateFields[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *RetailerInstance) validateCategoryIds(formats strfmt.Registry) error {

	if err := validate.Required("category_ids", "body", m.CategoryIds); err != nil {
		return err
	}

	return nil
}

func (m *RetailerInstance) validateCountry(formats strfmt.Registry) error {

	if err := validate.Required("country", "body", m.Country); err != nil {
		return err
	}

	return nil
}

func (m *RetailerInstance) validateEmailAddress(formats strfmt.Registry) error {

	if err := validate.Required("email_address", "body", m.EmailAddress); err != nil {
		return err
	}

	return nil
}

func (m *RetailerInstance) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

var retailerInstanceFeaturesItemsEnum []interface{}

func (m *RetailerInstance) validateFeaturesItemsEnum(path, location string, value string) error {
	if retailerInstanceFeaturesItemsEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["products","can_syndicate"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			retailerInstanceFeaturesItemsEnum = append(retailerInstanceFeaturesItemsEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, retailerInstanceFeaturesItemsEnum); err != nil {
		return err
	}
	return nil
}

func (m *RetailerInstance) validateFeatures(formats strfmt.Registry) error {

	if err := validate.Required("features", "body", m.Features); err != nil {
		return err
	}

	for i := 0; i < len(m.Features); i++ {

		// value enum
		if err := m.validateFeaturesItemsEnum("features"+"."+strconv.Itoa(i), "body", m.Features[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *RetailerInstance) validateLeaseID(formats strfmt.Registry) error {

	if err := validate.Required("lease_id", "body", m.LeaseID); err != nil {
		return err
	}

	return nil
}

func (m *RetailerInstance) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var retailerInstanceTypeProductTierPropEnum []interface{}

// prop value enum
func (m *RetailerInstance) validateProductTierEnum(path, location string, value int64) error {
	if retailerInstanceTypeProductTierPropEnum == nil {
		var res []int64
		if err := json.Unmarshal([]byte(`[1,2,3]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			retailerInstanceTypeProductTierPropEnum = append(retailerInstanceTypeProductTierPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, retailerInstanceTypeProductTierPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *RetailerInstance) validateProductTier(formats strfmt.Registry) error {

	if err := validate.Required("product_tier", "body", m.ProductTier); err != nil {
		return err
	}

	// value enum
	if err := m.validateProductTierEnum("product_tier", "body", *m.ProductTier); err != nil {
		return err
	}

	return nil
}

func (m *RetailerInstance) validateRetailerCode(formats strfmt.Registry) error {

	if err := validate.Required("retailer_code", "body", m.RetailerCode); err != nil {
		return err
	}

	return nil
}

func (m *RetailerInstance) validateRetailerID(formats strfmt.Registry) error {

	if err := validate.Required("retailer_id", "body", m.RetailerID); err != nil {
		return err
	}

	return nil
}

func (m *RetailerInstance) validateRetailerWebsite(formats strfmt.Registry) error {

	if err := validate.Required("retailer_website", "body", m.RetailerWebsite); err != nil {
		return err
	}

	return nil
}

func (m *RetailerInstance) validateServiceCategoryIds(formats strfmt.Registry) error {

	if err := validate.Required("service_category_ids", "body", m.ServiceCategoryIds); err != nil {
		return err
	}

	return nil
}

func (m *RetailerInstance) validateStoreProfile(formats strfmt.Registry) error {

	if err := validate.Required("store_profile", "body", m.StoreProfile); err != nil {
		return err
	}

	if err := validate.MaxLength("store_profile", "body", string(*m.StoreProfile), 1000); err != nil {
		return err
	}

	return nil
}

func (m *RetailerInstance) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	return nil
}
