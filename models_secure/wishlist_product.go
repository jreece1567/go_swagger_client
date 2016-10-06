package models_secure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*WishlistProduct Product data.

swagger:model wishlistProduct
*/
type WishlistProduct struct {

	/* An image associated with this product (first image of first variant).
	 */
	Image string `json:"image,omitempty"`

	/* Display name of product.
	 */
	Name string `json:"name,omitempty"`

	/* Price of this product.
	 */
	Price float64 `json:"price,omitempty"`

	/* Unique id associated with retailer of this product.
	 */
	RetailerCode string `json:"retailer_code,omitempty"`

	/* Name associated with retailer of this product.
	 */
	RetailerName string `json:"retailer_name,omitempty"`

	/* Sale price of this product.
	 */
	SalePrice float64 `json:"sale_price,omitempty"`
}

// Validate validates this wishlist product
func (m *WishlistProduct) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}