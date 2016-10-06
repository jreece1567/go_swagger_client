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

/*ServiceTypeUpdateBody service type update body

swagger:model serviceTypeUpdateBody
*/
type ServiceTypeUpdateBody struct {

	/* Country where the service type is available
	 */
	Country string `json:"country,omitempty"`

	/* Name of service type
	 */
	Name string `json:"name,omitempty"`

	/* Class of the service type
	 */
	ServiceClass string `json:"service_class,omitempty"`
}

// Validate validates this service type update body
func (m *ServiceTypeUpdateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCountry(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateServiceClass(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var serviceTypeUpdateBodyTypeCountryPropEnum []interface{}

// prop value enum
func (m *ServiceTypeUpdateBody) validateCountryEnum(path, location string, value string) error {
	if serviceTypeUpdateBodyTypeCountryPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["au","nz","uk","us"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			serviceTypeUpdateBodyTypeCountryPropEnum = append(serviceTypeUpdateBodyTypeCountryPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, serviceTypeUpdateBodyTypeCountryPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ServiceTypeUpdateBody) validateCountry(formats strfmt.Registry) error {

	if swag.IsZero(m.Country) { // not required
		return nil
	}

	// value enum
	if err := m.validateCountryEnum("country", "body", m.Country); err != nil {
		return err
	}

	return nil
}

var serviceTypeUpdateBodyTypeServiceClassPropEnum []interface{}

// prop value enum
func (m *ServiceTypeUpdateBody) validateServiceClassEnum(path, location string, value string) error {
	if serviceTypeUpdateBodyTypeServiceClassPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["physical","digital"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			serviceTypeUpdateBodyTypeServiceClassPropEnum = append(serviceTypeUpdateBodyTypeServiceClassPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, serviceTypeUpdateBodyTypeServiceClassPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ServiceTypeUpdateBody) validateServiceClass(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceClass) { // not required
		return nil
	}

	// value enum
	if err := m.validateServiceClassEnum("service_class", "body", m.ServiceClass); err != nil {
		return err
	}

	return nil
}
