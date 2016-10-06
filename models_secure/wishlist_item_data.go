package models_secure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

/*WishlistItemData Wishlist item's resource data.

swagger:model wishlistItemData
*/
type WishlistItemData struct {

	/* deal
	 */
	Deal *WishlistDeal `json:"deal,omitempty"`

	/* event
	 */
	Event *WishlistEvent `json:"event,omitempty"`

	/* product
	 */
	Product *WishlistProduct `json:"product,omitempty"`

	/* retailer
	 */
	Retailer *WishlistRetailer `json:"retailer,omitempty"`
}

// Validate validates this wishlist item data
func (m *WishlistItemData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeal(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEvent(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProduct(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRetailer(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WishlistItemData) validateDeal(formats strfmt.Registry) error {

	if swag.IsZero(m.Deal) { // not required
		return nil
	}

	if m.Deal != nil {

		if err := m.Deal.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *WishlistItemData) validateEvent(formats strfmt.Registry) error {

	if swag.IsZero(m.Event) { // not required
		return nil
	}

	if m.Event != nil {

		if err := m.Event.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *WishlistItemData) validateProduct(formats strfmt.Registry) error {

	if swag.IsZero(m.Product) { // not required
		return nil
	}

	if m.Product != nil {

		if err := m.Product.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *WishlistItemData) validateRetailer(formats strfmt.Registry) error {

	if swag.IsZero(m.Retailer) { // not required
		return nil
	}

	if m.Retailer != nil {

		if err := m.Retailer.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
