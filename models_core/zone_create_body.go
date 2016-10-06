package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*ZoneCreateBody Zone create parameters in message body

swagger:model zoneCreateBody
*/
type ZoneCreateBody struct {

	/* Centre identifier

	Required: true
	*/
	CentreID *string `json:"centre_id"`

	/* Zone color

	Required: true
	*/
	Color *Color `json:"color"`

	/* Zone name

	Required: true
	Max Length: 24
	*/
	Name *string `json:"name"`

	/* status

	Required: true
	*/
	Status *string `json:"status"`

	/* waypoints

	Required: true
	Min Items: 3
	*/
	Waypoints []int64 `json:"waypoints"`
}

// Validate validates this zone create body
func (m *ZoneCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCentreID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateColor(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateWaypoints(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ZoneCreateBody) validateCentreID(formats strfmt.Registry) error {

	if err := validate.Required("centre_id", "body", m.CentreID); err != nil {
		return err
	}

	return nil
}

func (m *ZoneCreateBody) validateColor(formats strfmt.Registry) error {

	if err := validate.Required("color", "body", m.Color); err != nil {
		return err
	}

	if m.Color != nil {

		if err := m.Color.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *ZoneCreateBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 24); err != nil {
		return err
	}

	return nil
}

var zoneCreateBodyTypeStatusPropEnum []interface{}

// prop value enum
func (m *ZoneCreateBody) validateStatusEnum(path, location string, value string) error {
	if zoneCreateBodyTypeStatusPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["active","inactive"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			zoneCreateBodyTypeStatusPropEnum = append(zoneCreateBodyTypeStatusPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, zoneCreateBodyTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ZoneCreateBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *ZoneCreateBody) validateWaypoints(formats strfmt.Registry) error {

	if err := validate.Required("waypoints", "body", m.Waypoints); err != nil {
		return err
	}

	iWaypointsSize := int64(len(m.Waypoints))

	if err := validate.MinItems("waypoints", "body", iWaypointsSize, 3); err != nil {
		return err
	}

	return nil
}
