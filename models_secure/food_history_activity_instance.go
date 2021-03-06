package models_secure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*FoodHistoryActivityInstance Food order

swagger:model foodHistoryActivityInstance
*/
type FoodHistoryActivityInstance struct {

	/* The coupons that are applied to the order

	Required: true
	*/
	ActiveCoupons []*OrderCouponInstance `json:"active_coupons"`

	/* Creation date for the given order

	Required: true
	*/
	CreatedAt *strfmt.DateTime `json:"created_at"`

	/* credit card

	Required: true
	*/
	CreditCard *CreditCardInstance `json:"credit_card"`

	/* Delivery time for the order.
	 */
	DeliveryTime strfmt.DateTime `json:"delivery_time,omitempty"`

	/* Food Order identifier.

	Required: true
	*/
	FoodOrderID *string `json:"food_order_id"`

	/* Items for the food order

	Required: true
	*/
	OrderItems []*OrderItemInstance `json:"order_items"`

	/* Name of the restaurant

	Required: true
	*/
	Restaurant *string `json:"restaurant"`

	/* Human friendly order number.

	Required: true
	*/
	ShortID *int64 `json:"short_id"`

	/* Textual order status.

	Required: true
	*/
	Status *string `json:"status"`

	/* The order subtotal.

	Required: true
	*/
	Subtotal *string `json:"subtotal"`

	/* The order tax amount.

	Required: true
	*/
	Tax *string `json:"tax"`

	/* The total amount of the order.

	Required: true
	*/
	TotalAmount *string `json:"total_amount"`
}

// Validate validates this food history activity instance
func (m *FoodHistoryActivityInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActiveCoupons(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCreditCard(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFoodOrderID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOrderItems(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRestaurant(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateShortID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSubtotal(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTax(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTotalAmount(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FoodHistoryActivityInstance) validateActiveCoupons(formats strfmt.Registry) error {

	if err := validate.Required("active_coupons", "body", m.ActiveCoupons); err != nil {
		return err
	}

	for i := 0; i < len(m.ActiveCoupons); i++ {

		if swag.IsZero(m.ActiveCoupons[i]) { // not required
			continue
		}

		if m.ActiveCoupons[i] != nil {

			if err := m.ActiveCoupons[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *FoodHistoryActivityInstance) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (m *FoodHistoryActivityInstance) validateCreditCard(formats strfmt.Registry) error {

	if err := validate.Required("credit_card", "body", m.CreditCard); err != nil {
		return err
	}

	if m.CreditCard != nil {

		if err := m.CreditCard.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *FoodHistoryActivityInstance) validateFoodOrderID(formats strfmt.Registry) error {

	if err := validate.Required("food_order_id", "body", m.FoodOrderID); err != nil {
		return err
	}

	return nil
}

func (m *FoodHistoryActivityInstance) validateOrderItems(formats strfmt.Registry) error {

	if err := validate.Required("order_items", "body", m.OrderItems); err != nil {
		return err
	}

	for i := 0; i < len(m.OrderItems); i++ {

		if swag.IsZero(m.OrderItems[i]) { // not required
			continue
		}

		if m.OrderItems[i] != nil {

			if err := m.OrderItems[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *FoodHistoryActivityInstance) validateRestaurant(formats strfmt.Registry) error {

	if err := validate.Required("restaurant", "body", m.Restaurant); err != nil {
		return err
	}

	return nil
}

func (m *FoodHistoryActivityInstance) validateShortID(formats strfmt.Registry) error {

	if err := validate.Required("short_id", "body", m.ShortID); err != nil {
		return err
	}

	return nil
}

func (m *FoodHistoryActivityInstance) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *FoodHistoryActivityInstance) validateSubtotal(formats strfmt.Registry) error {

	if err := validate.Required("subtotal", "body", m.Subtotal); err != nil {
		return err
	}

	return nil
}

func (m *FoodHistoryActivityInstance) validateTax(formats strfmt.Registry) error {

	if err := validate.Required("tax", "body", m.Tax); err != nil {
		return err
	}

	return nil
}

func (m *FoodHistoryActivityInstance) validateTotalAmount(formats strfmt.Registry) error {

	if err := validate.Required("total_amount", "body", m.TotalAmount); err != nil {
		return err
	}

	return nil
}
