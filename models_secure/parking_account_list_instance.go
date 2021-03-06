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

/*ParkingAccountListInstance Parking Account

swagger:model parkingAccountListInstance
*/
type ParkingAccountListInstance struct {

	/* access devices

	Required: true
	*/
	AccessDevices []*AccessDevicesInstance `json:"access_devices"`

	/* Account status

	Required: true
	*/
	AccountStatus *string `json:"account_status"`

	/* Date-time the account was created.

	Required: true
	*/
	CreatedAt *strfmt.DateTime `json:"created_at"`

	/* Account parking card pan number shadowed.

	Required: true
	*/
	CreditCardNumber *string `json:"credit_card_number"`

	/* Account email.

	Required: true
	*/
	Email *string `json:"email"`

	/* Account first name.

	Required: true
	*/
	FirstName *string `json:"first_name"`

	/* Account last name.

	Required: true
	*/
	LastName *string `json:"last_name"`

	/* Account parking card payment token.

	Required: true
	*/
	PaymentMethodToken *string `json:"payment_method_token"`

	/* Westfield UID for an account.

	Required: true
	*/
	PersonID *string `json:"person_id"`

	/* Account rule group GUID.

	Required: true
	*/
	RuleGroupGUID *string `json:"rule_group_guid"`

	/* Date-time the account was last updated.

	Required: true
	*/
	UpdatedAt *strfmt.DateTime `json:"updated_at"`
}

// Validate validates this parking account list instance
func (m *ParkingAccountListInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessDevices(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAccountStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCreditCardNumber(formats); err != nil {
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

	if err := m.validateLastName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePaymentMethodToken(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePersonID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRuleGroupGUID(formats); err != nil {
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

func (m *ParkingAccountListInstance) validateAccessDevices(formats strfmt.Registry) error {

	if err := validate.Required("access_devices", "body", m.AccessDevices); err != nil {
		return err
	}

	for i := 0; i < len(m.AccessDevices); i++ {

		if swag.IsZero(m.AccessDevices[i]) { // not required
			continue
		}

		if m.AccessDevices[i] != nil {

			if err := m.AccessDevices[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

var parkingAccountListInstanceTypeAccountStatusPropEnum []interface{}

// prop value enum
func (m *ParkingAccountListInstance) validateAccountStatusEnum(path, location string, value string) error {
	if parkingAccountListInstanceTypeAccountStatusPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["open","closed","blocked"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			parkingAccountListInstanceTypeAccountStatusPropEnum = append(parkingAccountListInstanceTypeAccountStatusPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, parkingAccountListInstanceTypeAccountStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ParkingAccountListInstance) validateAccountStatus(formats strfmt.Registry) error {

	if err := validate.Required("account_status", "body", m.AccountStatus); err != nil {
		return err
	}

	// value enum
	if err := m.validateAccountStatusEnum("account_status", "body", *m.AccountStatus); err != nil {
		return err
	}

	return nil
}

func (m *ParkingAccountListInstance) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *ParkingAccountListInstance) validateCreditCardNumber(formats strfmt.Registry) error {

	if err := validate.Required("credit_card_number", "body", m.CreditCardNumber); err != nil {
		return err
	}

	return nil
}

func (m *ParkingAccountListInstance) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *ParkingAccountListInstance) validateFirstName(formats strfmt.Registry) error {

	if err := validate.Required("first_name", "body", m.FirstName); err != nil {
		return err
	}

	return nil
}

func (m *ParkingAccountListInstance) validateLastName(formats strfmt.Registry) error {

	if err := validate.Required("last_name", "body", m.LastName); err != nil {
		return err
	}

	return nil
}

func (m *ParkingAccountListInstance) validatePaymentMethodToken(formats strfmt.Registry) error {

	if err := validate.Required("payment_method_token", "body", m.PaymentMethodToken); err != nil {
		return err
	}

	return nil
}

func (m *ParkingAccountListInstance) validatePersonID(formats strfmt.Registry) error {

	if err := validate.Required("person_id", "body", m.PersonID); err != nil {
		return err
	}

	return nil
}

func (m *ParkingAccountListInstance) validateRuleGroupGUID(formats strfmt.Registry) error {

	if err := validate.Required("rule_group_guid", "body", m.RuleGroupGUID); err != nil {
		return err
	}

	return nil
}

func (m *ParkingAccountListInstance) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	return nil
}
