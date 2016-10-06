package models_secure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

/*WishlistEvent Event data.

swagger:model wishlistEvent
*/
type WishlistEvent struct {

	/* image
	 */
	Image *GenericLink `json:"image,omitempty"`

	/* name
	 */
	Name string `json:"name,omitempty"`

	/* A list of datetimes that the event occurs.
	 */
	Occurences []*WishlistEventOccurrence `json:"occurences,omitempty"`

	/* Retailers and stores associated with the event
	 */
	Retailers []*EventRetailer `json:"retailers,omitempty"`
}

// Validate validates this wishlist event
func (m *WishlistEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOccurences(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRetailers(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WishlistEvent) validateImage(formats strfmt.Registry) error {

	if swag.IsZero(m.Image) { // not required
		return nil
	}

	if m.Image != nil {

		if err := m.Image.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *WishlistEvent) validateOccurences(formats strfmt.Registry) error {

	if swag.IsZero(m.Occurences) { // not required
		return nil
	}

	for i := 0; i < len(m.Occurences); i++ {

		if swag.IsZero(m.Occurences[i]) { // not required
			continue
		}

		if m.Occurences[i] != nil {

			if err := m.Occurences[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *WishlistEvent) validateRetailers(formats strfmt.Registry) error {

	if swag.IsZero(m.Retailers) { // not required
		return nil
	}

	for i := 0; i < len(m.Retailers); i++ {

		if swag.IsZero(m.Retailers[i]) { // not required
			continue
		}

		if m.Retailers[i] != nil {

			if err := m.Retailers[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}