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

/*NewAccount new account

swagger:model newAccount
*/
type NewAccount struct {

	/* Source that is requesting the user creation

	Required: true
	*/
	CreationSource *string `json:"creation_source"`

	/* email

	Required: true
	*/
	Email *string `json:"email"`

	/* Value is true if user is subscribed to marketing
	 */
	EmailMarketingOptIn bool `json:"email_marketing_opt_in,omitempty"`

	/* first name

	Required: true
	Max Length: 50
	*/
	FirstName *string `json:"first_name"`

	/* User gender
	 */
	Gender string `json:"gender,omitempty"`

	/* last name

	Max Length: 50
	*/
	LastName string `json:"last_name,omitempty"`

	/* List of newsletter identifiers in response that the user is subscribed to
	 */
	NewsletterSubscriptions []string `json:"newsletter_subscriptions,omitempty"`

	/* Identifier of parking
	 */
	ParkingID string `json:"parking_id,omitempty"`

	/* Account password.

	Max Length: 30
	Min Length: 8
	*/
	Password string `json:"password,omitempty"`

	/* User's mobile phone number following the E.164 standard
	 */
	PhoneNumber string `json:"phone_number,omitempty"`

	/* Identifier of the user primary center

	Required: true
	*/
	PrimaryCentreID *string `json:"primary_centre_id"`

	/* User timezone
	 */
	Timezone string `json:"timezone,omitempty"`

	/* Custom data related to the user provided by wifi vendor
	 */
	WifiCustomData string `json:"wifi_custom_data,omitempty"`

	/* User identifier on wifi vendor database
	 */
	WifiID string `json:"wifi_id,omitempty"`
}

// Validate validates this new account
func (m *NewAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationSource(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFirstName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGender(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLastName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNewsletterSubscriptions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
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

var newAccountTypeCreationSourcePropEnum []interface{}

// prop value enum
func (m *NewAccount) validateCreationSourceEnum(path, location string, value string) error {
	if newAccountTypeCreationSourcePropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["mobile_signup","web_signup","wifi","newsletter","wfss","contest","in_centre","gift_card","parking"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			newAccountTypeCreationSourcePropEnum = append(newAccountTypeCreationSourcePropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, newAccountTypeCreationSourcePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NewAccount) validateCreationSource(formats strfmt.Registry) error {

	if err := validate.Required("creation_source", "body", m.CreationSource); err != nil {
		return err
	}

	// value enum
	if err := m.validateCreationSourceEnum("creation_source", "body", *m.CreationSource); err != nil {
		return err
	}

	return nil
}

func (m *NewAccount) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *NewAccount) validateFirstName(formats strfmt.Registry) error {

	if err := validate.Required("first_name", "body", m.FirstName); err != nil {
		return err
	}

	if err := validate.MaxLength("first_name", "body", string(*m.FirstName), 50); err != nil {
		return err
	}

	return nil
}

var newAccountTypeGenderPropEnum []interface{}

// prop value enum
func (m *NewAccount) validateGenderEnum(path, location string, value string) error {
	if newAccountTypeGenderPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["female","male"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			newAccountTypeGenderPropEnum = append(newAccountTypeGenderPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, newAccountTypeGenderPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NewAccount) validateGender(formats strfmt.Registry) error {

	if swag.IsZero(m.Gender) { // not required
		return nil
	}

	// value enum
	if err := m.validateGenderEnum("gender", "body", m.Gender); err != nil {
		return err
	}

	return nil
}

func (m *NewAccount) validateLastName(formats strfmt.Registry) error {

	if swag.IsZero(m.LastName) { // not required
		return nil
	}

	if err := validate.MaxLength("last_name", "body", string(m.LastName), 50); err != nil {
		return err
	}

	return nil
}

func (m *NewAccount) validateNewsletterSubscriptions(formats strfmt.Registry) error {

	if swag.IsZero(m.NewsletterSubscriptions) { // not required
		return nil
	}

	return nil
}

func (m *NewAccount) validatePassword(formats strfmt.Registry) error {

	if swag.IsZero(m.Password) { // not required
		return nil
	}

	if err := validate.MinLength("password", "body", string(m.Password), 8); err != nil {
		return err
	}

	if err := validate.MaxLength("password", "body", string(m.Password), 30); err != nil {
		return err
	}

	return nil
}

func (m *NewAccount) validatePrimaryCentreID(formats strfmt.Registry) error {

	if err := validate.Required("primary_centre_id", "body", m.PrimaryCentreID); err != nil {
		return err
	}

	return nil
}
