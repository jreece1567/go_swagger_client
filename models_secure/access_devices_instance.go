package models_secure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*AccessDevicesInstance AccessDevices

swagger:model accessDevicesInstance
*/
type AccessDevicesInstance struct {

	/* Device identifier.

	Required: true
	*/
	DeviceID *string `json:"device_id"`

	/* Device status

	Required: true
	*/
	DeviceStatus *string `json:"device_status"`

	/* Device type

	Required: true
	*/
	Type *string `json:"type"`
}

// Validate validates this access devices instance
func (m *AccessDevicesInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeviceID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDeviceStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccessDevicesInstance) validateDeviceID(formats strfmt.Registry) error {

	if err := validate.Required("device_id", "body", m.DeviceID); err != nil {
		return err
	}

	return nil
}

var accessDevicesInstanceTypeDeviceStatusPropEnum []interface{}

// prop value enum
func (m *AccessDevicesInstance) validateDeviceStatusEnum(path, location string, value string) error {
	if accessDevicesInstanceTypeDeviceStatusPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["ok","removed"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			accessDevicesInstanceTypeDeviceStatusPropEnum = append(accessDevicesInstanceTypeDeviceStatusPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, accessDevicesInstanceTypeDeviceStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AccessDevicesInstance) validateDeviceStatus(formats strfmt.Registry) error {

	if err := validate.Required("device_status", "body", m.DeviceStatus); err != nil {
		return err
	}

	// value enum
	if err := m.validateDeviceStatusEnum("device_status", "body", *m.DeviceStatus); err != nil {
		return err
	}

	return nil
}

func (m *AccessDevicesInstance) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}