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

/*WishlistItemInstance wishlist item instance

swagger:model wishlistItemInstance
*/
type WishlistItemInstance struct {

	/* Date and Time when item was added to the Wishlist.

	Required: true
	*/
	CreatedAt *strfmt.DateTime `json:"created_at"`

	/* data

	Required: true
	*/
	Data *WishlistItemData `json:"data"`

	/* Wishlist item's external resource kind.

	Required: true
	*/
	Kind *string `json:"kind"`

	/* Platform source where it is being saved
	 */
	PlatformSource string `json:"platform_source,omitempty"`

	/* Wishlist item's external resource identifier.

	Required: true
	*/
	ResourceID *string `json:"resource_id"`
}

// Validate validates this wishlist item instance
func (m *WishlistItemInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateData(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateKind(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePlatformSource(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateResourceID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WishlistItemInstance) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *WishlistItemInstance) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	if m.Data != nil {

		if err := m.Data.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

var wishlistItemInstanceTypeKindPropEnum []interface{}

// prop value enum
func (m *WishlistItemInstance) validateKindEnum(path, location string, value string) error {
	if wishlistItemInstanceTypeKindPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["product","deal","event","retailer"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			wishlistItemInstanceTypeKindPropEnum = append(wishlistItemInstanceTypeKindPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, wishlistItemInstanceTypeKindPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WishlistItemInstance) validateKind(formats strfmt.Registry) error {

	if err := validate.Required("kind", "body", m.Kind); err != nil {
		return err
	}

	// value enum
	if err := m.validateKindEnum("kind", "body", *m.Kind); err != nil {
		return err
	}

	return nil
}

var wishlistItemInstanceTypePlatformSourcePropEnum []interface{}

// prop value enum
func (m *WishlistItemInstance) validatePlatformSourceEnum(path, location string, value string) error {
	if wishlistItemInstanceTypePlatformSourcePropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["android","ios","mobile_web","web","concierge"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			wishlistItemInstanceTypePlatformSourcePropEnum = append(wishlistItemInstanceTypePlatformSourcePropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, wishlistItemInstanceTypePlatformSourcePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *WishlistItemInstance) validatePlatformSource(formats strfmt.Registry) error {

	if swag.IsZero(m.PlatformSource) { // not required
		return nil
	}

	// value enum
	if err := m.validatePlatformSourceEnum("platform_source", "body", m.PlatformSource); err != nil {
		return err
	}

	return nil
}

func (m *WishlistItemInstance) validateResourceID(formats strfmt.Registry) error {

	if err := validate.Required("resource_id", "body", m.ResourceID); err != nil {
		return err
	}

	return nil
}
