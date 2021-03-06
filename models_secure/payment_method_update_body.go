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

/*PaymentMethodUpdateBody Data to update the details associated with a payment method.

swagger:model paymentMethodUpdateBody
*/
type PaymentMethodUpdateBody struct {

	/* The country code used to verify the payment method details of the customer. This country code is read-only and cannot be changed
	 */
	Country *string `json:"country,omitempty"`

	/* Customer first name.
	 */
	FirstName string `json:"first_name,omitempty"`

	/* Customer last name.
	 */
	LastName string `json:"last_name,omitempty"`

	/* Expiration month.
	 */
	Month int64 `json:"month,omitempty"`

	/* Customer postal code.
	 */
	PostalCode string `json:"postal_code,omitempty"`

	/* Expiration year.
	 */
	Year int64 `json:"year,omitempty"`
}

// Validate validates this payment method update body
func (m *PaymentMethodUpdateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCountry(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var paymentMethodUpdateBodyTypeCountryPropEnum []interface{}

// prop value enum
func (m *PaymentMethodUpdateBody) validateCountryEnum(path, location string, value string) error {
	if paymentMethodUpdateBodyTypeCountryPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["AU","NZ","UK","US"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			paymentMethodUpdateBodyTypeCountryPropEnum = append(paymentMethodUpdateBodyTypeCountryPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, paymentMethodUpdateBodyTypeCountryPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PaymentMethodUpdateBody) validateCountry(formats strfmt.Registry) error {

	if swag.IsZero(m.Country) { // not required
		return nil
	}

	// value enum
	if err := m.validateCountryEnum("country", "body", *m.Country); err != nil {
		return err
	}

	return nil
}
