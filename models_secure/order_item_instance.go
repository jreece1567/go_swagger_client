package models_secure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*OrderItemInstance Food order item

swagger:model orderItemInstance
*/
type OrderItemInstance struct {

	/* Name of the item.

	Required: true
	*/
	Name *string `json:"name"`

	/* options

	Required: true
	*/
	Options []*OrderOptionInstance `json:"options"`

	/* Price of the item.

	Required: true
	*/
	Price *string `json:"price"`

	/* Ordered quantity for the item.
	 */
	Quantity float64 `json:"quantity,omitempty"`

	/* Price multiplied by the item quantity.

	Required: true
	*/
	Subtotal *string `json:"subtotal"`

	/* The total price of this item, including options.

	Required: true
	*/
	TotalPrice *string `json:"total_price"`
}

// Validate validates this order item instance
func (m *OrderItemInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOptions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePrice(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSubtotal(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTotalPrice(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrderItemInstance) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *OrderItemInstance) validateOptions(formats strfmt.Registry) error {

	if err := validate.Required("options", "body", m.Options); err != nil {
		return err
	}

	for i := 0; i < len(m.Options); i++ {

		if swag.IsZero(m.Options[i]) { // not required
			continue
		}

		if m.Options[i] != nil {

			if err := m.Options[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *OrderItemInstance) validatePrice(formats strfmt.Registry) error {

	if err := validate.Required("price", "body", m.Price); err != nil {
		return err
	}

	return nil
}

func (m *OrderItemInstance) validateSubtotal(formats strfmt.Registry) error {

	if err := validate.Required("subtotal", "body", m.Subtotal); err != nil {
		return err
	}

	return nil
}

func (m *OrderItemInstance) validateTotalPrice(formats strfmt.Registry) error {

	if err := validate.Required("total_price", "body", m.TotalPrice); err != nil {
		return err
	}

	return nil
}
