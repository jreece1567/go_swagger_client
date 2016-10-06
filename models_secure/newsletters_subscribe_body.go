package models_secure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*NewslettersSubscribeBody Data to subscribe account to newsletters

swagger:model newslettersSubscribeBody
*/
type NewslettersSubscribeBody struct {

	/* Country code of the account to subscribe to newsletters.

	Required: true
	Max Length: 2
	Min Length: 2
	*/
	Country *string `json:"country"`

	/* Email of the partial account to subscribe to newsletters.

	Required: true
	*/
	Email *string `json:"email"`

	/* First name of the account.
	 */
	FirstName string `json:"first_name,omitempty"`

	/* List of newsletter identifiers to subscribe the account to

	Required: true
	*/
	NewsletterSubscriptions []string `json:"newsletter_subscriptions"`

	/* Primary centre identifier. Identifier of the primary centre of the account

	Required: true
	*/
	PrimaryCentreID *string `json:"primary_centre_id"`
}

// Validate validates this newsletters subscribe body
func (m *NewslettersSubscribeBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCountry(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNewsletterSubscriptions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePrimaryCentreID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var newslettersSubscribeBodyTypeCountryPropEnum []interface{}

// prop value enum
func (m *NewslettersSubscribeBody) validateCountryEnum(path, location string, value string) error {
	if newslettersSubscribeBodyTypeCountryPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["US","UK","AU","NZ"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			newslettersSubscribeBodyTypeCountryPropEnum = append(newslettersSubscribeBodyTypeCountryPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, newslettersSubscribeBodyTypeCountryPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NewslettersSubscribeBody) validateCountry(formats strfmt.Registry) error {

	if err := validate.Required("country", "body", m.Country); err != nil {
		return err
	}

	if err := validate.MinLength("country", "body", string(*m.Country), 2); err != nil {
		return err
	}

	if err := validate.MaxLength("country", "body", string(*m.Country), 2); err != nil {
		return err
	}

	// value enum
	if err := m.validateCountryEnum("country", "body", *m.Country); err != nil {
		return err
	}

	return nil
}

func (m *NewslettersSubscribeBody) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *NewslettersSubscribeBody) validateNewsletterSubscriptions(formats strfmt.Registry) error {

	if err := validate.Required("newsletter_subscriptions", "body", m.NewsletterSubscriptions); err != nil {
		return err
	}

	return nil
}

func (m *NewslettersSubscribeBody) validatePrimaryCentreID(formats strfmt.Registry) error {

	if err := validate.Required("primary_centre_id", "body", m.PrimaryCentreID); err != nil {
		return err
	}

	return nil
}