package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*MovieLinks movie links

swagger:model movieLinks
*/
type MovieLinks struct {

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

	/* List of trailers for this movie

	Required: true
	*/
	TrailerHighRes *TrailerLink `json:"trailer_high_res"`

	/* List of trailers for this movie

	Required: true
	*/
	TrailerLowRes *TrailerLink `json:"trailer_low_res"`

	/* List of trailers for this movie

	Required: true
	*/
	Trailers []*TrailerLink `json:"trailers"`
}

// Validate validates this movie links
func (m *MovieLinks) Validate(formats strfmt.Registry) error {
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

	if err := m.validateTrailerHighRes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTrailerLowRes(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTrailers(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MovieLinks) validateBookingURL(formats strfmt.Registry) error {

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

func (m *MovieLinks) validateImage(formats strfmt.Registry) error {

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

func (m *MovieLinks) validateSelf(formats strfmt.Registry) error {

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

func (m *MovieLinks) validateTrailerHighRes(formats strfmt.Registry) error {

	if err := validate.Required("trailer_high_res", "body", m.TrailerHighRes); err != nil {
		return err
	}

	if m.TrailerHighRes != nil {

		if err := m.TrailerHighRes.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *MovieLinks) validateTrailerLowRes(formats strfmt.Registry) error {

	if err := validate.Required("trailer_low_res", "body", m.TrailerLowRes); err != nil {
		return err
	}

	if m.TrailerLowRes != nil {

		if err := m.TrailerLowRes.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *MovieLinks) validateTrailers(formats strfmt.Registry) error {

	if err := validate.Required("trailers", "body", m.Trailers); err != nil {
		return err
	}

	for i := 0; i < len(m.Trailers); i++ {

		if swag.IsZero(m.Trailers[i]) { // not required
			continue
		}

		if m.Trailers[i] != nil {

			if err := m.Trailers[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}