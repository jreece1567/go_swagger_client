package models_secure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*MerchantInstance Merchant details.

swagger:model merchantInstance
*/
type MerchantInstance struct {

	/* Value is true if transactions for this merchant are charged on behalf of a connected Stripe account through the platform. Value is false if transactions for this merchant are charged directly on the connected Stripe account.
	 */
	ChargeViaPlatform bool `json:"charge_via_platform,omitempty"`

	/* Unique code used to process transactions on this merchant.

	Required: true
	*/
	Code *string `json:"code"`

	/* Date-time the merchant was created.

	Required: true
	*/
	CreatedAt *strfmt.DateTime `json:"created_at"`

	/* Custom display name which can be assigned to this merchant object.

	Required: true
	*/
	DisplayName *string `json:"display_name"`

	/* List of features the merchant has. If 'food' is present, the merchant is capable of food purchases.

	Required: true
	*/
	Features []string `json:"features"`

	/* Identifier of the gateway that the merchant is related.

	Required: true
	*/
	GatewayID *int64 `json:"gateway_id"`

	/* Custom location description that may end up on bank statement.
	 */
	LocationDescriptor string `json:"location_descriptor,omitempty"`

	/* Merchant identifier. Identifier to retrieve this merchant.

	Required: true
	*/
	MerchantID *int64 `json:"merchant_id"`

	/* Name of the merchant.

	Required: true
	*/
	Name *string `json:"name"`

	/* Custom name description that may end up on bank statement.
	 */
	NameDescriptor string `json:"name_descriptor,omitempty"`

	/* Date-time the merchant was last updated.

	Required: true
	*/
	UpdatedAt *strfmt.DateTime `json:"updated_at"`
}

// Validate validates this merchant instance
func (m *MerchantInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDisplayName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFeatures(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGatewayID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMerchantID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
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

func (m *MerchantInstance) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *MerchantInstance) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *MerchantInstance) validateDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("display_name", "body", m.DisplayName); err != nil {
		return err
	}

	return nil
}

var merchantInstanceFeaturesItemsEnum []interface{}

func (m *MerchantInstance) validateFeaturesItemsEnum(path, location string, value string) error {
	if merchantInstanceFeaturesItemsEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["food"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			merchantInstanceFeaturesItemsEnum = append(merchantInstanceFeaturesItemsEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, merchantInstanceFeaturesItemsEnum); err != nil {
		return err
	}
	return nil
}

func (m *MerchantInstance) validateFeatures(formats strfmt.Registry) error {

	if err := validate.Required("features", "body", m.Features); err != nil {
		return err
	}

	for i := 0; i < len(m.Features); i++ {

		// value enum
		if err := m.validateFeaturesItemsEnum("features"+"."+strconv.Itoa(i), "body", m.Features[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *MerchantInstance) validateGatewayID(formats strfmt.Registry) error {

	if err := validate.Required("gateway_id", "body", m.GatewayID); err != nil {
		return err
	}

	return nil
}

func (m *MerchantInstance) validateMerchantID(formats strfmt.Registry) error {

	if err := validate.Required("merchant_id", "body", m.MerchantID); err != nil {
		return err
	}

	return nil
}

func (m *MerchantInstance) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *MerchantInstance) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	return nil
}
