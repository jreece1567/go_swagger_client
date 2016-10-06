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

/*AccountTokenRequestBody Account-token request parameters.

swagger:model accountTokenRequestBody
*/
type AccountTokenRequestBody struct {

	/* Country code of the account.

	Max Length: 2
	Min Length: 2
	*/
	Country string `json:"country,omitempty"`

	/* Grant Type

	Required: true
	*/
	GrantType *string `json:"grant_type"`

	/* User password
	 */
	Password string `json:"password,omitempty"`

	/* Refresh token
	 */
	RefreshToken string `json:"refresh_token,omitempty"`

	/* User email
	 */
	Username string `json:"username,omitempty"`
}

// Validate validates this account token request body
func (m *AccountTokenRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCountry(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGrantType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var accountTokenRequestBodyTypeCountryPropEnum []interface{}

// prop value enum
func (m *AccountTokenRequestBody) validateCountryEnum(path, location string, value string) error {
	if accountTokenRequestBodyTypeCountryPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["US","UK","AU","NZ"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			accountTokenRequestBodyTypeCountryPropEnum = append(accountTokenRequestBodyTypeCountryPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, accountTokenRequestBodyTypeCountryPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AccountTokenRequestBody) validateCountry(formats strfmt.Registry) error {

	if swag.IsZero(m.Country) { // not required
		return nil
	}

	if err := validate.MinLength("country", "body", string(m.Country), 2); err != nil {
		return err
	}

	if err := validate.MaxLength("country", "body", string(m.Country), 2); err != nil {
		return err
	}

	// value enum
	if err := m.validateCountryEnum("country", "body", m.Country); err != nil {
		return err
	}

	return nil
}

var accountTokenRequestBodyTypeGrantTypePropEnum []interface{}

// prop value enum
func (m *AccountTokenRequestBody) validateGrantTypeEnum(path, location string, value string) error {
	if accountTokenRequestBodyTypeGrantTypePropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["password","refresh_token"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			accountTokenRequestBodyTypeGrantTypePropEnum = append(accountTokenRequestBodyTypeGrantTypePropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, accountTokenRequestBodyTypeGrantTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AccountTokenRequestBody) validateGrantType(formats strfmt.Registry) error {

	if err := validate.Required("grant_type", "body", m.GrantType); err != nil {
		return err
	}

	// value enum
	if err := m.validateGrantTypeEnum("grant_type", "body", *m.GrantType); err != nil {
		return err
	}

	return nil
}
