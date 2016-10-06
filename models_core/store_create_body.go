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

/*StoreCreateBody Parameters for creating a store

swagger:model storeCreateBody
*/
type StoreCreateBody struct {

	/* List of Category IDs associated with the store
	 */
	CategoryIds []int64 `json:"category_ids,omitempty"`

	/* Identifier of the centre where the store is located.

	Required: true
	*/
	CentreID *string `json:"centre_id"`

	/* Long description of the store's retailer
	 */
	Description string `json:"description,omitempty"`

	/* Dining information, if this store has a dining facility. Not provided otherwise.
	 */
	Dining *DiningCreateBody `json:"dining,omitempty"`

	/* [DEPRECATED]. Cuisine offered by the store if it provides dining facilities
	 */
	DiningCuisine string `json:"dining_cuisine,omitempty"`

	/* [DEPRECATED]. Storage provider identifier of the dining image of the store if it provides dining facilities
	 */
	DiningImageRef string `json:"dining_image_ref,omitempty"`

	/* [DEPRECATED]. Storage provider identifier of the menu of the store if it provides dining facilities
	 */
	DiningMenuRef string `json:"dining_menu_ref,omitempty"`

	/* [DEPRECATED]. Number indicating the relative price of dining at the store

	Maximum: 4
	*/
	DiningPriceGuide int64 `json:"dining_price_guide,omitempty"`

	/* [DEPRECATED]. Dining vendor identifier if the store is a dining facility that participates in Open Dining Network; an empty string otherwise.
	 */
	DiningVendorID string `json:"dining_vendor_id,omitempty"`

	/* Date the store was disabled in the Westfield system
	 */
	DisabledAt strfmt.DateTime `json:"disabled_at,omitempty"`

	/* E-mail address of the store
	 */
	EmailAddress string `json:"email_address,omitempty"`

	/* Flag indicating if the store is enabled in the Westfield system
	 */
	Enabled bool `json:"enabled,omitempty"`

	/* Date the store was enabled in the Westfield system
	 */
	EnabledAt strfmt.DateTime `json:"enabled_at,omitempty"`

	/* List of features the store has. If 'foodordering' is present the store supports ordering food via the Food App. If 'giftcards' is present the store accepts giftcards.
	 */
	Features []string `json:"features,omitempty"`

	/* Number of the store's lease with Westfield
	 */
	LeaseNumber string `json:"lease_number,omitempty"`

	/* List of level names sourced form Jibestream locations
	 */
	LocationLevels []string `json:"location_levels,omitempty"`

	/* List of jibestream locations and meta data
	 */
	Locations []*StoreLocation `json:"locations,omitempty"`

	/* Name of the store

	Required: true
	*/
	Name *string `json:"name"`

	/* Phone number of the store

	Required: true
	*/
	PhoneNumber *string `json:"phone_number"`

	/* Identifier of the retailer the store belongs to

	Required: true
	*/
	RetailerID *int64 `json:"retailer_id"`

	/* URL of the store's website
	 */
	RetailerWebsite string `json:"retailer_website,omitempty"`

	/* Numeric value used for wayfinding applications

	Required: true
	*/
	Salience *int64 `json:"salience"`

	/* List of Service Category identifiers associated with the store
	 */
	ServiceCategoryIds []int64 `json:"service_category_ids,omitempty"`

	/* The shop number in the Westfield centre the store occupies
	 */
	ShopNumber string `json:"shop_number,omitempty"`

	/* Storage provider identifier of the store's store front image
	 */
	StoreFrontImageRef string `json:"store_front_image_ref,omitempty"`
}

// Validate validates this store create body
func (m *StoreCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategoryIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCentreID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDining(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDiningPriceGuide(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFeatures(formats); err != nil {
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

	if err := m.validateRetailerID(formats); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StoreCreateBody) validateCategoryIds(formats strfmt.Registry) error {

	if swag.IsZero(m.CategoryIds) { // not required
		return nil
	}

	return nil
}

func (m *StoreCreateBody) validateCentreID(formats strfmt.Registry) error {

	if err := validate.Required("centre_id", "body", m.CentreID); err != nil {
		return err
	}

	return nil
}

func (m *StoreCreateBody) validateDining(formats strfmt.Registry) error {

	if swag.IsZero(m.Dining) { // not required
		return nil
	}

	if m.Dining != nil {

		if err := m.Dining.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *StoreCreateBody) validateDiningPriceGuide(formats strfmt.Registry) error {

	if swag.IsZero(m.DiningPriceGuide) { // not required
		return nil
	}

	if err := validate.MaximumInt("dining_price_guide", "body", int64(m.DiningPriceGuide), 4, false); err != nil {
		return err
	}

	return nil
}

var storeCreateBodyFeaturesItemsEnum []interface{}

func (m *StoreCreateBody) validateFeaturesItemsEnum(path, location string, value string) error {
	if storeCreateBodyFeaturesItemsEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["cinema","foodordering","giftcards","products"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			storeCreateBodyFeaturesItemsEnum = append(storeCreateBodyFeaturesItemsEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, storeCreateBodyFeaturesItemsEnum); err != nil {
		return err
	}
	return nil
}

func (m *StoreCreateBody) validateFeatures(formats strfmt.Registry) error {

	if swag.IsZero(m.Features) { // not required
		return nil
	}

	for i := 0; i < len(m.Features); i++ {

		// value enum
		if err := m.validateFeaturesItemsEnum("features"+"."+strconv.Itoa(i), "body", m.Features[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *StoreCreateBody) validateLocationLevels(formats strfmt.Registry) error {

	if swag.IsZero(m.LocationLevels) { // not required
		return nil
	}

	return nil
}

func (m *StoreCreateBody) validateLocations(formats strfmt.Registry) error {

	if swag.IsZero(m.Locations) { // not required
		return nil
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

func (m *StoreCreateBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *StoreCreateBody) validatePhoneNumber(formats strfmt.Registry) error {

	if err := validate.Required("phone_number", "body", m.PhoneNumber); err != nil {
		return err
	}

	return nil
}

func (m *StoreCreateBody) validateRetailerID(formats strfmt.Registry) error {

	if err := validate.Required("retailer_id", "body", m.RetailerID); err != nil {
		return err
	}

	return nil
}

func (m *StoreCreateBody) validateSalience(formats strfmt.Registry) error {

	if err := validate.Required("salience", "body", m.Salience); err != nil {
		return err
	}

	return nil
}

func (m *StoreCreateBody) validateServiceCategoryIds(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceCategoryIds) { // not required
		return nil
	}

	return nil
}
