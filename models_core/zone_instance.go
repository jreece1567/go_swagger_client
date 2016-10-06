package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*ZoneInstance Zone record

swagger:model zoneInstance
*/
type ZoneInstance struct {

	/* Hash of links

	Required: true
	*/
	Links *ZoneLinks `json:"_links"`

	/* Centre identifier

	Required: true
	*/
	CentreID *string `json:"centre_id"`

	/* Zone color

	Required: true
	*/
	Color *Color `json:"color"`

	/* Created date

	Required: true
	*/
	CreatedAt *strfmt.DateTime `json:"created_at"`

	/* Zone name

	Required: true
	Max Length: 24
	*/
	Name *string `json:"name"`

	/* Primary waypoint

	Required: true
	*/
	PrimaryWaypoint *int64 `json:"primary_waypoint"`

	/* status

	Required: true
	*/
	Status *string `json:"status"`

	/* Updated date

	Required: true
	*/
	UpdatedAt *strfmt.DateTime `json:"updated_at"`

	/* waypoints

	Required: true
	Min Items: 3
	*/
	Waypoints []int64 `json:"waypoints"`

	/* Zone identifier. Identifier to retrieve this zone.

	Required: true
	*/
	ZoneID *int64 `json:"zone_id"`
}

// Validate validates this zone instance
func (m *ZoneInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCentreID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateColor(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePrimaryWaypoint(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateWaypoints(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateZoneID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ZoneInstance) validateLinks(formats strfmt.Registry) error {

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

func (m *ZoneInstance) validateCentreID(formats strfmt.Registry) error {

	if err := validate.Required("centre_id", "body", m.CentreID); err != nil {
		return err
	}

	return nil
}

func (m *ZoneInstance) validateColor(formats strfmt.Registry) error {

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

func (m *ZoneInstance) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *ZoneInstance) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 24); err != nil {
		return err
	}

	return nil
}

func (m *ZoneInstance) validatePrimaryWaypoint(formats strfmt.Registry) error {

	if err := validate.Required("primary_waypoint", "body", m.PrimaryWaypoint); err != nil {
		return err
	}

	return nil
}

var zoneInstanceTypeStatusPropEnum []interface{}

// prop value enum
func (m *ZoneInstance) validateStatusEnum(path, location string, value string) error {
	if zoneInstanceTypeStatusPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["active","inactive"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			zoneInstanceTypeStatusPropEnum = append(zoneInstanceTypeStatusPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, zoneInstanceTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ZoneInstance) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *ZoneInstance) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	return nil
}

func (m *ZoneInstance) validateWaypoints(formats strfmt.Registry) error {

	if err := validate.Required("waypoints", "body", m.Waypoints); err != nil {
		return err
	}

	iWaypointsSize := int64(len(m.Waypoints))

	if err := validate.MinItems("waypoints", "body", iWaypointsSize, 3); err != nil {
		return err
	}

	return nil
}

func (m *ZoneInstance) validateZoneID(formats strfmt.Registry) error {

	if err := validate.Required("zone_id", "body", m.ZoneID); err != nil {
		return err
	}

	return nil
}