package models_secure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*AccountStatusResponse Account status

swagger:model accountStatusResponse
*/
type AccountStatusResponse struct {

	/* data
	 */
	Data string `json:"data,omitempty"`

	/* errors
	 */
	Errors EmptyObject `json:"errors,omitempty"`

	/* meta
	 */
	Meta *MetaResponse `json:"meta,omitempty"`
}

// Validate validates this account status response
func (m *AccountStatusResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMeta(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var accountStatusResponseTypeDataPropEnum []interface{}

// prop value enum
func (m *AccountStatusResponse) validateDataEnum(path, location string, value string) error {
	if accountStatusResponseTypeDataPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["full","partial","none","deleted","locked"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			accountStatusResponseTypeDataPropEnum = append(accountStatusResponseTypeDataPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, accountStatusResponseTypeDataPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AccountStatusResponse) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	// value enum
	if err := m.validateDataEnum("data", "body", m.Data); err != nil {
		return err
	}

	return nil
}

func (m *AccountStatusResponse) validateMeta(formats strfmt.Registry) error {

	if swag.IsZero(m.Meta) { // not required
		return nil
	}

	if m.Meta != nil {

		if err := m.Meta.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}