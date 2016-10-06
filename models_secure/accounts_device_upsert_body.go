package models_secure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*AccountsDeviceUpsertBody Data to upsert a device.

swagger:model accountsDeviceUpsertBody
*/
type AccountsDeviceUpsertBody struct {

	/* device locale

	Required: true
	*/
	Locale *string `json:"locale"`

	/* device maker

	Required: true
	*/
	Maker *string `json:"maker"`

	/* device mobile carrier
	 */
	MobileCarrier string `json:"mobile_carrier,omitempty"`

	/* device model

	Required: true
	*/
	Model *string `json:"model"`

	/* Notification badge counter.

	Required: true
	*/
	NotificationBadge *string `json:"notification_badge"`

	/* Token which allows push notifications to be sent to a device.
	 */
	NotificationToken string `json:"notification_token,omitempty"`

	/* device operational system version

	Required: true
	*/
	OsVersion *string `json:"os_version"`

	/* device platform

	Required: true
	*/
	Platform *string `json:"platform"`

	/* device timezone

	Required: true
	*/
	Timezone *string `json:"timezone"`
}

// Validate validates this accounts device upsert body
func (m *AccountsDeviceUpsertBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLocale(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMaker(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateModel(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNotificationBadge(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOsVersion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePlatform(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTimezone(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountsDeviceUpsertBody) validateLocale(formats strfmt.Registry) error {

	if err := validate.Required("locale", "body", m.Locale); err != nil {
		return err
	}

	return nil
}

func (m *AccountsDeviceUpsertBody) validateMaker(formats strfmt.Registry) error {

	if err := validate.Required("maker", "body", m.Maker); err != nil {
		return err
	}

	return nil
}

func (m *AccountsDeviceUpsertBody) validateModel(formats strfmt.Registry) error {

	if err := validate.Required("model", "body", m.Model); err != nil {
		return err
	}

	return nil
}

func (m *AccountsDeviceUpsertBody) validateNotificationBadge(formats strfmt.Registry) error {

	if err := validate.Required("notification_badge", "body", m.NotificationBadge); err != nil {
		return err
	}

	return nil
}

func (m *AccountsDeviceUpsertBody) validateOsVersion(formats strfmt.Registry) error {

	if err := validate.Required("os_version", "body", m.OsVersion); err != nil {
		return err
	}

	return nil
}

var accountsDeviceUpsertBodyTypePlatformPropEnum []interface{}

// prop value enum
func (m *AccountsDeviceUpsertBody) validatePlatformEnum(path, location string, value string) error {
	if accountsDeviceUpsertBodyTypePlatformPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["ios","android"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			accountsDeviceUpsertBodyTypePlatformPropEnum = append(accountsDeviceUpsertBodyTypePlatformPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, accountsDeviceUpsertBodyTypePlatformPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AccountsDeviceUpsertBody) validatePlatform(formats strfmt.Registry) error {

	if err := validate.Required("platform", "body", m.Platform); err != nil {
		return err
	}

	// value enum
	if err := m.validatePlatformEnum("platform", "body", *m.Platform); err != nil {
		return err
	}

	return nil
}

func (m *AccountsDeviceUpsertBody) validateTimezone(formats strfmt.Registry) error {

	if err := validate.Required("timezone", "body", m.Timezone); err != nil {
		return err
	}

	return nil
}