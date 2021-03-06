package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*CentreInstance Centre

swagger:model centreInstance
*/
type CentreInstance struct {

	/* Hash of links

	Required: true
	*/
	Links *CentreLinks `json:"_links"`

	/* Business unit code

	Required: true
	*/
	BusinessUnitCode *string `json:"business_unit_code"`

	/* Centre identifier

	Required: true
	*/
	CentreID *string `json:"centre_id"`

	/* Country

	Required: true
	*/
	Country *string `json:"country"`

	/* Deleted date
	 */
	DeletedAt strfmt.DateTime `json:"deleted_at,omitempty"`

	/* Date centre will-be/was disabled

	Required: true
	*/
	DisabledAt *strfmt.DateTime `json:"disabled_at"`

	/* Email address

	Required: true
	*/
	EmailAddress *string `json:"email_address"`

	/* Centre is enabled (true/false)

	Required: true
	*/
	Enabled *bool `json:"enabled"`

	/* Date centre is/was enabled

	Required: true
	*/
	EnabledAt *strfmt.DateTime `json:"enabled_at"`

	/* features

	Required: true
	*/
	Features []string `json:"features"`

	/* Hash of geofence information with possible keys 'building', 'parking', or 'area'

	Required: true
	*/
	Geofences *GeofenceList `json:"geofences"`

	/* Latitude

	Required: true
	*/
	Latitude *float64 `json:"latitude"`

	/* Longitude

	Required: true
	*/
	Longitude *float64 `json:"longitude"`

	/* Map location identifier

	Required: true
	*/
	MapID *string `json:"map_id"`

	/* Name

	Required: true
	*/
	Name *string `json:"name"`

	/* Number of floor levels

	Required: true
	*/
	NumberOfLevels *int64 `json:"number_of_levels"`

	/* Owner of centre

	Required: true
	*/
	Owner *string `json:"owner"`

	/* Phone number

	Required: true
	*/
	PhoneNumber *string `json:"phone_number"`

	/* Postal Code

	Required: true
	*/
	Postcode *string `json:"postcode"`

	/* Short name

	Required: true
	*/
	ShortName *string `json:"short_name"`

	/* Social media links

	Required: true
	*/
	SocialMedia *SocialMediaList `json:"social_media"`

	/* State

	Required: true
	*/
	State *string `json:"state"`

	/* Street address

	Required: true
	*/
	StreetAddress *string `json:"street_address"`

	/* City/Suburb

	Required: true
	*/
	Suburb *string `json:"suburb"`

	/* Tier of Centre

	Required: true
	*/
	Theme *string `json:"theme"`

	/* Centre's time zone

	Required: true
	*/
	TimeZone *string `json:"time_zone"`

	/* Type of asset

	Required: true
	*/
	Type *string `json:"type"`

	/* Updated date

	Required: true
	*/
	UpdatedAt *strfmt.DateTime `json:"updated_at"`
}

// Validate validates this centre instance
func (m *CentreInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateBusinessUnitCode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCentreID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCountry(formats); err != nil {
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

	if err := m.validateGeofences(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLatitude(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLongitude(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMapID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNumberOfLevels(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePhoneNumber(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePostcode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateShortName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSocialMedia(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStreetAddress(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSuburb(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTheme(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTimeZone(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
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

func (m *CentreInstance) validateLinks(formats strfmt.Registry) error {

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

func (m *CentreInstance) validateBusinessUnitCode(formats strfmt.Registry) error {

	if err := validate.Required("business_unit_code", "body", m.BusinessUnitCode); err != nil {
		return err
	}

	return nil
}

func (m *CentreInstance) validateCentreID(formats strfmt.Registry) error {

	if err := validate.Required("centre_id", "body", m.CentreID); err != nil {
		return err
	}

	return nil
}

var centreInstanceTypeCountryPropEnum []interface{}

// prop value enum
func (m *CentreInstance) validateCountryEnum(path, location string, value string) error {
	if centreInstanceTypeCountryPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["AU","NZ","UK","US"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			centreInstanceTypeCountryPropEnum = append(centreInstanceTypeCountryPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, centreInstanceTypeCountryPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CentreInstance) validateCountry(formats strfmt.Registry) error {

	if err := validate.Required("country", "body", m.Country); err != nil {
		return err
	}

	// value enum
	if err := m.validateCountryEnum("country", "body", *m.Country); err != nil {
		return err
	}

	return nil
}

func (m *CentreInstance) validateDisabledAt(formats strfmt.Registry) error {

	if err := validate.Required("disabled_at", "body", m.DisabledAt); err != nil {
		return err
	}

	return nil
}

func (m *CentreInstance) validateEmailAddress(formats strfmt.Registry) error {

	if err := validate.Required("email_address", "body", m.EmailAddress); err != nil {
		return err
	}

	return nil
}

func (m *CentreInstance) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *CentreInstance) validateEnabledAt(formats strfmt.Registry) error {

	if err := validate.Required("enabled_at", "body", m.EnabledAt); err != nil {
		return err
	}

	return nil
}

func (m *CentreInstance) validateFeatures(formats strfmt.Registry) error {

	if err := validate.Required("features", "body", m.Features); err != nil {
		return err
	}

	return nil
}

func (m *CentreInstance) validateGeofences(formats strfmt.Registry) error {

	if err := validate.Required("geofences", "body", m.Geofences); err != nil {
		return err
	}

	if m.Geofences != nil {

		if err := m.Geofences.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *CentreInstance) validateLatitude(formats strfmt.Registry) error {

	if err := validate.Required("latitude", "body", m.Latitude); err != nil {
		return err
	}

	return nil
}

func (m *CentreInstance) validateLongitude(formats strfmt.Registry) error {

	if err := validate.Required("longitude", "body", m.Longitude); err != nil {
		return err
	}

	return nil
}

func (m *CentreInstance) validateMapID(formats strfmt.Registry) error {

	if err := validate.Required("map_id", "body", m.MapID); err != nil {
		return err
	}

	return nil
}

func (m *CentreInstance) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CentreInstance) validateNumberOfLevels(formats strfmt.Registry) error {

	if err := validate.Required("number_of_levels", "body", m.NumberOfLevels); err != nil {
		return err
	}

	return nil
}

func (m *CentreInstance) validateOwner(formats strfmt.Registry) error {

	if err := validate.Required("owner", "body", m.Owner); err != nil {
		return err
	}

	return nil
}

func (m *CentreInstance) validatePhoneNumber(formats strfmt.Registry) error {

	if err := validate.Required("phone_number", "body", m.PhoneNumber); err != nil {
		return err
	}

	return nil
}

func (m *CentreInstance) validatePostcode(formats strfmt.Registry) error {

	if err := validate.Required("postcode", "body", m.Postcode); err != nil {
		return err
	}

	return nil
}

func (m *CentreInstance) validateShortName(formats strfmt.Registry) error {

	if err := validate.Required("short_name", "body", m.ShortName); err != nil {
		return err
	}

	return nil
}

func (m *CentreInstance) validateSocialMedia(formats strfmt.Registry) error {

	if err := validate.Required("social_media", "body", m.SocialMedia); err != nil {
		return err
	}

	if m.SocialMedia != nil {

		if err := m.SocialMedia.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *CentreInstance) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *CentreInstance) validateStreetAddress(formats strfmt.Registry) error {

	if err := validate.Required("street_address", "body", m.StreetAddress); err != nil {
		return err
	}

	return nil
}

func (m *CentreInstance) validateSuburb(formats strfmt.Registry) error {

	if err := validate.Required("suburb", "body", m.Suburb); err != nil {
		return err
	}

	return nil
}

func (m *CentreInstance) validateTheme(formats strfmt.Registry) error {

	if err := validate.Required("theme", "body", m.Theme); err != nil {
		return err
	}

	return nil
}

func (m *CentreInstance) validateTimeZone(formats strfmt.Registry) error {

	if err := validate.Required("time_zone", "body", m.TimeZone); err != nil {
		return err
	}

	return nil
}

func (m *CentreInstance) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *CentreInstance) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	return nil
}
