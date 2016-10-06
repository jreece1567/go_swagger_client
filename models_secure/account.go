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

/*Account Account including associated identification data.

swagger:model account
*/
type Account struct {

	/* DateTime that the shopper was created

	Required: true
	*/
	CreatedAt *string `json:"created_at"`

	/* Email address for the shopper

	Required: true
	*/
	Email *string `json:"email"`

	/* Value is true if user is subscribed to marketing

	Required: true
	*/
	EmailMarketingOptIn *bool `json:"email_marketing_opt_in"`

	/* First name for the shopper

	Required: true
	*/
	FirstName *string `json:"first_name"`

	/* User gender
	 */
	Gender string `json:"gender,omitempty"`

	/* Last name for the shopper
	 */
	LastName string `json:"last_name,omitempty"`

	/* messaging preferences
	 */
	MessagingPreferences *MessagingPreferenceInstance `json:"messaging_preferences,omitempty"`

	/* List of centre identifiers in response associated with the newsletter subscriptions
	 */
	NewsletterCentreIds []string `json:"newsletter_centre_ids,omitempty"`

	/* List of newsletter identifiers in response that the user is subscribed to
	 */
	NewsletterSubscriptions []string `json:"newsletter_subscriptions,omitempty"`

	/* Identifier of the account

	Required: true
	*/
	PersonID *string `json:"person_id"`

	/* User's mobile phone number following the E.164 standard
	 */
	PhoneNumber string `json:"phone_number,omitempty"`

	/* Identifier of the user primary center

	Required: true
	*/
	PrimaryCentreID *string `json:"primary_centre_id"`

	/* Value is true if user signed up for parking
	 */
	RegisteredForParking bool `json:"registered_for_parking,omitempty"`

	/* Value is true if user is subscribed to sms marketing
	 */
	SmsMarketingOptIn bool `json:"sms_marketing_opt_in,omitempty"`

	/* User timezone
	 */
	Timezone string `json:"timezone,omitempty"`

	/* DateTime that the shopper last updated

	Required: true
	*/
	UpdatedAt *string `json:"updated_at"`
}

// Validate validates this account
func (m *Account) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEmailMarketingOptIn(formats); err != nil {
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

	if err := m.validateMessagingPreferences(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNewsletterCentreIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNewsletterSubscriptions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePersonID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePrimaryCentreID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Account) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *Account) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *Account) validateEmailMarketingOptIn(formats strfmt.Registry) error {

	if err := validate.Required("email_marketing_opt_in", "body", m.EmailMarketingOptIn); err != nil {
		return err
	}

	return nil
}

func (m *Account) validateFirstName(formats strfmt.Registry) error {

	if err := validate.Required("first_name", "body", m.FirstName); err != nil {
		return err
	}

	return nil
}

var accountTypeGenderPropEnum []interface{}

// prop value enum
func (m *Account) validateGenderEnum(path, location string, value string) error {
	if accountTypeGenderPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["female","male"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			accountTypeGenderPropEnum = append(accountTypeGenderPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, accountTypeGenderPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Account) validateGender(formats strfmt.Registry) error {

	if swag.IsZero(m.Gender) { // not required
		return nil
	}

	// value enum
	if err := m.validateGenderEnum("gender", "body", m.Gender); err != nil {
		return err
	}

	return nil
}

func (m *Account) validateMessagingPreferences(formats strfmt.Registry) error {

	if swag.IsZero(m.MessagingPreferences) { // not required
		return nil
	}

	if m.MessagingPreferences != nil {

		if err := m.MessagingPreferences.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *Account) validateNewsletterCentreIds(formats strfmt.Registry) error {

	if swag.IsZero(m.NewsletterCentreIds) { // not required
		return nil
	}

	return nil
}

func (m *Account) validateNewsletterSubscriptions(formats strfmt.Registry) error {

	if swag.IsZero(m.NewsletterSubscriptions) { // not required
		return nil
	}

	return nil
}

func (m *Account) validatePersonID(formats strfmt.Registry) error {

	if err := validate.Required("person_id", "body", m.PersonID); err != nil {
		return err
	}

	return nil
}

func (m *Account) validatePrimaryCentreID(formats strfmt.Registry) error {

	if err := validate.Required("primary_centre_id", "body", m.PrimaryCentreID); err != nil {
		return err
	}

	return nil
}

func (m *Account) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	return nil
}
