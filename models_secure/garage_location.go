package models_secure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*GarageLocation Data to create a car park.

swagger:model garageLocation
*/
type GarageLocation struct {

	/* Total number of available parking spaces for the location.

	Required: true
	*/
	Available *int64 `json:"available"`

	/* Total number of parking spaces for the location.

	Required: true
	*/
	Capacity *int64 `json:"capacity"`

	/* Condition of the location

	Required: true
	*/
	Condition *string `json:"condition"`

	/* Location identifier. Identifier to uniquely identify this particular location in the garage.

	Required: true
	*/
	Location *string `json:"location"`

	/* Total number of occupied carpark spots for the centre.

	Required: true
	*/
	Occupied *int64 `json:"occupied"`

	/* Total number of reserved carpark spots for the centre.

	Required: true
	*/
	Reserved *int64 `json:"reserved"`
}

// Validate validates this garage location
func (m *GarageLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailable(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCapacity(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCondition(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOccupied(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateReserved(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GarageLocation) validateAvailable(formats strfmt.Registry) error {

	if err := validate.Required("available", "body", m.Available); err != nil {
		return err
	}

	return nil
}

func (m *GarageLocation) validateCapacity(formats strfmt.Registry) error {

	if err := validate.Required("capacity", "body", m.Capacity); err != nil {
		return err
	}

	return nil
}

var garageLocationTypeConditionPropEnum []interface{}

// prop value enum
func (m *GarageLocation) validateConditionEnum(path, location string, value string) error {
	if garageLocationTypeConditionPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["spaces","faulty","full","closed"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			garageLocationTypeConditionPropEnum = append(garageLocationTypeConditionPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, garageLocationTypeConditionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GarageLocation) validateCondition(formats strfmt.Registry) error {

	if err := validate.Required("condition", "body", m.Condition); err != nil {
		return err
	}

	// value enum
	if err := m.validateConditionEnum("condition", "body", *m.Condition); err != nil {
		return err
	}

	return nil
}

func (m *GarageLocation) validateLocation(formats strfmt.Registry) error {

	if err := validate.Required("location", "body", m.Location); err != nil {
		return err
	}

	return nil
}

func (m *GarageLocation) validateOccupied(formats strfmt.Registry) error {

	if err := validate.Required("occupied", "body", m.Occupied); err != nil {
		return err
	}

	return nil
}

func (m *GarageLocation) validateReserved(formats strfmt.Registry) error {

	if err := validate.Required("reserved", "body", m.Reserved); err != nil {
		return err
	}

	return nil
}
