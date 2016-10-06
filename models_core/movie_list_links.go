package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*MovieListLinks movie list links

swagger:model movieListLinks
*/
type MovieListLinks struct {

	/* Link to purchase tickets for this movie

	Required: true
	*/
	BookingURL *GenericLink `json:"booking_url"`

	/* Cloudinary image details

	Required: true
	*/
	Image *GenericLink `json:"image"`

	/* Canonical link to this movie

	Required: true
	*/
	Self *GenericLink `json:"self"`
}

// Validate validates this movie list links
func (m *MovieListLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBookingURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MovieListLinks) validateBookingURL(formats strfmt.Registry) error {

	if err := validate.Required("booking_url", "body", m.BookingURL); err != nil {
		return err
	}

	if m.BookingURL != nil {

		if err := m.BookingURL.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *MovieListLinks) validateImage(formats strfmt.Registry) error {

	if err := validate.Required("image", "body", m.Image); err != nil {
		return err
	}

	if m.Image != nil {

		if err := m.Image.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *MovieListLinks) validateSelf(formats strfmt.Registry) error {

	if err := validate.Required("self", "body", m.Self); err != nil {
		return err
	}

	if m.Self != nil {

		if err := m.Self.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}
