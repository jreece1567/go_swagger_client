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

/*StoreInstance A store belonging to a retailer.

swagger:model storeInstance
*/
type StoreInstance struct {

	/* links

	Required: true
	*/
	Links *StoreLinks `json:"_links"`

	/* List of category codes

	Required: true
	*/
	CategoryCodes []string `json:"category_codes"`

	/* List of categories the store belongs to

	Required: true
	*/
	CategoryIds []int64 `json:"category_ids"`

	/* Identifier of the centre where the store is located.

	Required: true
	*/
	CentreID *string `json:"centre_id"`

	/* Date and time the store was marked as deleted
	 */
	DeletedAt strfmt.DateTime `json:"deleted_at,omitempty"`

	/* Long description of the store's retailer

	Required: true
	*/
	Description *string `json:"description"`

	/* Dining information, if this store has a dining facility. Nil otherwise.

	Required: true
	*/
	Dining *DiningInstance `json:"dining"`

	/* [DEPRECATED]. Cuisine offered by the store if it provides dining facilities

	Required: true
	*/
	DiningCuisine *string `json:"dining_cuisine"`

	/* [DEPRECATED]. Number indicating the relative price of dining at the store

	Required: true
	Maximum: 4
	*/
	DiningPriceGuide *int64 `json:"dining_price_guide"`

	/* [DEPRECATED]. Dining vendor identifier if the store is a dining facility that participates in Open Dining Network; an empty string otherwise.

	Required: true
	*/
	DiningVendorID *string `json:"dining_vendor_id"`

	/* Date the store was disabled in the Westfield system

	Required: true
	*/
	DisabledAt *strfmt.DateTime `json:"disabled_at"`

	/* E-mail address of the store

	Required: true
	*/
	EmailAddress *string `json:"email_address"`

	/* Flag indicating if the store is enabled in the Westfield system

	Required: true
	*/
	Enabled *bool `json:"enabled"`

	/* Date the store was enabled in the Westfield system

	Required: true
	*/
	EnabledAt *strfmt.DateTime `json:"enabled_at"`

	/* List of features the store has. If 'foodordering' is present the store supports ordering food via the Food App. If 'products' is present the retailer has products on the website. If 'giftcards' is present the store accepts giftcards.

	Required: true
	*/
	Features []string `json:"features"`

	/* Flag indicating if the store has deals associated with it

	Required: true
	*/
	HasDeals *bool `json:"has_deals"`

	/* [DEPRECATED]. Flag indicating if the store is a dining facility that participates in Open Dining Network.

	Required: true
	*/
	HasDining *bool `json:"has_dining"`

	/* Flag indicating if the store has events associated with it

	Required: true
	*/
	HasEvents *bool `json:"has_events"`

	/* Number of the store's lease with Westfield

	Required: true
	*/
	LeaseNumber *string `json:"lease_number"`

	/* List of level names sourced form Jibestream locations

	Required: true
	*/
	LocationLevels []string `json:"location_levels"`

	/* List of Jibestream locations and meta data

	Required: true
	*/
	Locations []*StoreLocation `json:"locations"`

	/* Name of the store

	Required: true
	*/
	Name *string `json:"name"`

	/* Phone number of the store

	Required: true
	*/
	PhoneNumber *string `json:"phone_number"`

	/* Code of the retailer the store belongs to

	Required: true
	*/
	RetailerCode *string `json:"retailer_code"`

	/* URL of the store's website

	Required: true
	*/
	RetailerWebsite *string `json:"retailer_website"`

	/* Numeric value used for wayfinding applications

	Required: true
	*/
	Salience *int64 `json:"salience"`

	/* List of service categories associated with the store

	Required: true
	*/
	ServiceCategoryIds []int64 `json:"service_category_ids"`

	/* The shop number in the Westfield centre the store occupies

	Required: true
	*/
	ShopNumber *string `json:"shop_number"`

	/* Status details for this store.

	Required: true
	*/
	Status *StatusInstance `json:"status"`

	/* Auto generated identifier

	Required: true
	*/
	StoreID *int64 `json:"store_id"`

	/* Date and time the retailer was last updated.

	Required: true
	*/
	UpdatedAt *strfmt.DateTime `json:"updated_at"`
}

// Validate validates this store instance
func (m *StoreInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCategoryCodes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCategoryIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCentreID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDining(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDiningCuisine(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDiningPriceGuide(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDiningVendorID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDisabledAt(formats); err != nil {
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

	if err := m.validateEnabledAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFeatures(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHasDeals(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHasDining(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHasEvents(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLeaseNumber(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLocationLevels(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLocations(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePhoneNumber(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRetailerCode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRetailerWebsite(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSalience(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateServiceCategoryIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateShopNumber(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStoreID(formats); err != nil {
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

func (m *StoreInstance) validateLinks(formats strfmt.Registry) error {

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

func (m *StoreInstance) validateCategoryCodes(formats strfmt.Registry) error {

	if err := validate.Required("category_codes", "body", m.CategoryCodes); err != nil {
		return err
	}

	return nil
}

func (m *StoreInstance) validateCategoryIds(formats strfmt.Registry) error {

	if err := validate.Required("category_ids", "body", m.CategoryIds); err != nil {
		return err
	}

	return nil
}

func (m *StoreInstance) validateCentreID(formats strfmt.Registry) error {

	if err := validate.Required("centre_id", "body", m.CentreID); err != nil {
		return err
	}

	return nil
}

func (m *StoreInstance) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *StoreInstance) validateDining(formats strfmt.Registry) error {

	if err := validate.Required("dining", "body", m.Dining); err != nil {
		return err
	}

	if m.Dining != nil {

		if err := m.Dining.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *StoreInstance) validateDiningCuisine(formats strfmt.Registry) error {

	if err := validate.Required("dining_cuisine", "body", m.DiningCuisine); err != nil {
		return err
	}

	return nil
}

func (m *StoreInstance) validateDiningPriceGuide(formats strfmt.Registry) error {

	if err := validate.Required("dining_price_guide", "body", m.DiningPriceGuide); err != nil {
		return err
	}

	if err := validate.MaximumInt("dining_price_guide", "body", int64(*m.DiningPriceGuide), 4, false); err != nil {
		return err
	}

	return nil
}

func (m *StoreInstance) validateDiningVendorID(formats strfmt.Registry) error {

	if err := validate.Required("dining_vendor_id", "body", m.DiningVendorID); err != nil {
		return err
	}

	return nil
}

func (m *StoreInstance) validateDisabledAt(formats strfmt.Registry) error {

	if err := validate.Required("disabled_at", "body", m.DisabledAt); err != nil {
		return err
	}

	return nil
}

func (m *StoreInstance) validateEmailAddress(formats strfmt.Registry) error {

	if err := validate.Required("email_address", "body", m.EmailAddress); err != nil {
		return err
	}

	return nil
}

func (m *StoreInstance) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *StoreInstance) validateEnabledAt(formats strfmt.Registry) error {

	if err := validate.Required("enabled_at", "body", m.EnabledAt); err != nil {
		return err
	}

	return nil
}

var storeInstanceFeaturesItemsEnum []interface{}

func (m *StoreInstance) validateFeaturesItemsEnum(path, location string, value string) error {
	if storeInstanceFeaturesItemsEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["cinema","foodordering","giftcards","products"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			storeInstanceFeaturesItemsEnum = append(storeInstanceFeaturesItemsEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, storeInstanceFeaturesItemsEnum); err != nil {
		return err
	}
	return nil
}

func (m *StoreInstance) validateFeatures(formats strfmt.Registry) error {

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

func (m *StoreInstance) validateHasDeals(formats strfmt.Registry) error {

	if err := validate.Required("has_deals", "body", m.HasDeals); err != nil {
		return err
	}

	return nil
}

func (m *StoreInstance) validateHasDining(formats strfmt.Registry) error {

	if err := validate.Required("has_dining", "body", m.HasDining); err != nil {
		return err
	}

	return nil
}

func (m *StoreInstance) validateHasEvents(formats strfmt.Registry) error {

	if err := validate.Required("has_events", "body", m.HasEvents); err != nil {
		return err
	}

	return nil
}

func (m *StoreInstance) validateLeaseNumber(formats strfmt.Registry) error {

	if err := validate.Required("lease_number", "body", m.LeaseNumber); err != nil {
		return err
	}

	return nil
}

func (m *StoreInstance) validateLocationLevels(formats strfmt.Registry) error {

	if err := validate.Required("location_levels", "body", m.LocationLevels); err != nil {
		return err
	}

	return nil
}

func (m *StoreInstance) validateLocations(formats strfmt.Registry) error {

	if err := validate.Required("locations", "body", m.Locations); err != nil {
		return err
	}

	for i := 0; i < len(m.Locations); i++ {

		if swag.IsZero(m.Locations[i]) { // not required
			continue
		}

		if m.Locations[i] != nil {

			if err := m.Locations[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *StoreInstance) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *StoreInstance) validatePhoneNumber(formats strfmt.Registry) error {

	if err := validate.Required("phone_number", "body", m.PhoneNumber); err != nil {
		return err
	}

	return nil
}

func (m *StoreInstance) validateRetailerCode(formats strfmt.Registry) error {

	if err := validate.Required("retailer_code", "body", m.RetailerCode); err != nil {
		return err
	}

	return nil
}

func (m *StoreInstance) validateRetailerWebsite(formats strfmt.Registry) error {

	if err := validate.Required("retailer_website", "body", m.RetailerWebsite); err != nil {
		return err
	}

	return nil
}

func (m *StoreInstance) validateSalience(formats strfmt.Registry) error {

	if err := validate.Required("salience", "body", m.Salience); err != nil {
		return err
	}

	return nil
}

func (m *StoreInstance) validateServiceCategoryIds(formats strfmt.Registry) error {

	if err := validate.Required("service_category_ids", "body", m.ServiceCategoryIds); err != nil {
		return err
	}

	return nil
}

func (m *StoreInstance) validateShopNumber(formats strfmt.Registry) error {

	if err := validate.Required("shop_number", "body", m.ShopNumber); err != nil {
		return err
	}

	return nil
}

func (m *StoreInstance) validateStatus(formats strfmt.Registry) error {

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

func (m *StoreInstance) validateStoreID(formats strfmt.Registry) error {

	if err := validate.Required("store_id", "body", m.StoreID); err != nil {
		return err
	}

	return nil
}

func (m *StoreInstance) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	return nil
}