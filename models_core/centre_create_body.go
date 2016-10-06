package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*CentreCreateBody Centre create parameters in message body

swagger:model centreCreateBody
*/
type CentreCreateBody struct {

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

	/* Medium-sized dining image filename
	 */
	DiningImageMediumRef string `json:"dining_image_medium_ref,omitempty"`

	/* Dining image filename
	 */
	DiningImageRef string `json:"dining_image_ref,omitempty"`

	/* Small-sized dining image filename
	 */
	DiningImageSmallRef string `json:"dining_image_small_ref,omitempty"`

	/* Date centre will-be/was disabled
	 */
	DisabledAt strfmt.DateTime `json:"disabled_at,omitempty"`

	/* Email address

	Required: true
	*/
	EmailAddress *string `json:"email_address"`

	/* Centre is enabled (true/false)
	 */
	Enabled bool `json:"enabled,omitempty"`

	/* Date centre is/was enabled
	 */
	EnabledAt strfmt.DateTime `json:"enabled_at,omitempty"`

	/* features
	 */
	Features []string `json:"features,omitempty"`

	/* Hash of geofence information with possible keys 'building', 'parking'', or 'area'
	 */
	Geofences *GeofenceList `json:"geofences,omitempty"`

	/* Hero image fielname
	 */
	HeroImageRef string `json:"hero_image_ref,omitempty"`

	/* Latitude

	Required: true
	*/
	Latitude *float64 `json:"latitude"`

	/* Longitude

	Required: true
	*/
	Longitude *float64 `json:"longitude"`

	/* Map location identifier
	 */
	MapID string `json:"map_id,omitempty"`

	/* Name

	Required: true
	*/
	Name *string `json:"name"`

	/* Number of floor levels
	 */
	NumberOfLevels int64 `json:"number_of_levels,omitempty"`

	/* Owner of centre
	 */
	Owner string `json:"owner,omitempty"`

	/* Phone number

	Required: true
	*/
	PhoneNumber *string `json:"phone_number"`

	/* Postal Code

	Required: true
	*/
	Postcode *string `json:"postcode"`

	/* Short name
	 */
	ShortName string `json:"short_name,omitempty"`

	/* Social media links
	 */
	SocialMedia *SocialMediaList `json:"social_media,omitempty"`

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

	/* Time zone of centre

	Required: true
	*/
	TimeZone *string `json:"time_zone"`

	/* Type of asset
	 */
	Type string `json:"type,omitempty"`
}

// Validate validates this centre create body
func (m *CentreCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

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

	if err := m.validateEmailAddress(formats); err != nil {
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

	if err := m.validateName(formats); err != nil {
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

	if err := m.validateTimeZone(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CentreCreateBody) validateBusinessUnitCode(formats strfmt.Registry) error {

	if err := validate.Required("business_unit_code", "body", m.BusinessUnitCode); err != nil {
		return err
	}

	return nil
}

func (m *CentreCreateBody) validateCentreID(formats strfmt.Registry) error {

	if err := validate.Required("centre_id", "body", m.CentreID); err != nil {
		return err
	}

	return nil
}

var centreCreateBodyTypeCountryPropEnum []interface{}

// prop value enum
func (m *CentreCreateBody) validateCountryEnum(path, location string, value string) error {
	if centreCreateBodyTypeCountryPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["AU","NZ","UK","US"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			centreCreateBodyTypeCountryPropEnum = append(centreCreateBodyTypeCountryPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, centreCreateBodyTypeCountryPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CentreCreateBody) validateCountry(formats strfmt.Registry) error {

	if err := validate.Required("country", "body", m.Country); err != nil {
		return err
	}

	// value enum
	if err := m.validateCountryEnum("country", "body", *m.Country); err != nil {
		return err
	}

	return nil
}

func (m *CentreCreateBody) validateEmailAddress(formats strfmt.Registry) error {

	if err := validate.Required("email_address", "body", m.EmailAddress); err != nil {
		return err
	}

	return nil
}

func (m *CentreCreateBody) validateFeatures(formats strfmt.Registry) error {

	if swag.IsZero(m.Features) { // not required
		return nil
	}

	return nil
}

func (m *CentreCreateBody) validateGeofences(formats strfmt.Registry) error {

	if swag.IsZero(m.Geofences) { // not required
		return nil
	}

	if m.Geofences != nil {

		if err := m.Geofences.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *CentreCreateBody) validateLatitude(formats strfmt.Registry) error {

	if err := validate.Required("latitude", "body", m.Latitude); err != nil {
		return err
	}

	return nil
}

func (m *CentreCreateBody) validateLongitude(formats strfmt.Registry) error {

	if err := validate.Required("longitude", "body", m.Longitude); err != nil {
		return err
	}

	return nil
}

func (m *CentreCreateBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CentreCreateBody) validatePhoneNumber(formats strfmt.Registry) error {

	if err := validate.Required("phone_number", "body", m.PhoneNumber); err != nil {
		return err
	}

	return nil
}

func (m *CentreCreateBody) validatePostcode(formats strfmt.Registry) error {

	if err := validate.Required("postcode", "body", m.Postcode); err != nil {
		return err
	}

	return nil
}

func (m *CentreCreateBody) validateSocialMedia(formats strfmt.Registry) error {

	if swag.IsZero(m.SocialMedia) { // not required
		return nil
	}

	if m.SocialMedia != nil {

		if err := m.SocialMedia.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *CentreCreateBody) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *CentreCreateBody) validateStreetAddress(formats strfmt.Registry) error {

	if err := validate.Required("street_address", "body", m.StreetAddress); err != nil {
		return err
	}

	return nil
}

func (m *CentreCreateBody) validateSuburb(formats strfmt.Registry) error {

	if err := validate.Required("suburb", "body", m.Suburb); err != nil {
		return err
	}

	return nil
}

func (m *CentreCreateBody) validateTimeZone(formats strfmt.Registry) error {

	if err := validate.Required("time_zone", "body", m.TimeZone); err != nil {
		return err
	}

	return nil
}