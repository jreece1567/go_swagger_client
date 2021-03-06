package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*EventCreateBody Parameters for creating an event

swagger:model eventCreateBody
*/
type EventCreateBody struct {

	/* Identifiers of categories the event belongs to
	 */
	CategoryIds []int64 `json:"category_ids,omitempty"`

	/* Centre the event belongs to

	Required: true
	*/
	CentreID *string `json:"centre_id"`

	/* Event description
	 */
	Description string `json:"description,omitempty"`

	/* Ending date and time of the last occurrence
	 */
	DisabledAt strfmt.DateTime `json:"disabled_at,omitempty"`

	/* External URL to buy tickets to event
	 */
	ExternalBuyURL string `json:"external_buy_url,omitempty"`

	/* External URL for event information
	 */
	ExternalURL string `json:"external_url,omitempty"`

	/* External URL information description
	 */
	ExternalURLDescription string `json:"external_url_description,omitempty"`

	/* Featured
	 */
	Featured bool `json:"featured,omitempty"`

	/* Image filename of event location

	Required: true
	*/
	ImageRef *string `json:"image_ref"`

	/* Description of event location
	 */
	Location string `json:"location,omitempty"`

	/* locations
	 */
	Locations []*Location `json:"locations,omitempty"`

	/* Event name

	Required: true
	*/
	Name *string `json:"name"`

	/* One or more occurrences of event

	Required: true
	*/
	Occurrences []*Occurrence `json:"occurrences"`

	/* The date-time the event was published

	Required: true
	*/
	PublishedAt *strfmt.DateTime `json:"published_at"`

	/* Retailers and stores associated with the event
	 */
	Retailers []*EventRetailer `json:"retailers,omitempty"`
}

// Validate validates this event create body
func (m *EventCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategoryIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCentreID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateImageRef(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLocations(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOccurrences(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePublishedAt(formats); err != nil {
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

func (m *EventCreateBody) validateCategoryIds(formats strfmt.Registry) error {

	if swag.IsZero(m.CategoryIds) { // not required
		return nil
	}

	return nil
}

func (m *EventCreateBody) validateCentreID(formats strfmt.Registry) error {

	if err := validate.Required("centre_id", "body", m.CentreID); err != nil {
		return err
	}

	return nil
}

func (m *EventCreateBody) validateImageRef(formats strfmt.Registry) error {

	if err := validate.Required("image_ref", "body", m.ImageRef); err != nil {
		return err
	}

	return nil
}

func (m *EventCreateBody) validateLocations(formats strfmt.Registry) error {

	if swag.IsZero(m.Locations) { // not required
		return nil
	}

	for i := 0; i < len(m.Locations); i++ {

		if swag.IsZero(m.Locations[i]) { // not required
			continue
		}

		if m.Locations[i] != nil {

			if err := m.Locations[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *EventCreateBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *EventCreateBody) validateOccurrences(formats strfmt.Registry) error {

	if err := validate.Required("occurrences", "body", m.Occurrences); err != nil {
		return err
	}

	for i := 0; i < len(m.Occurrences); i++ {

		if swag.IsZero(m.Occurrences[i]) { // not required
			continue
		}

		if m.Occurrences[i] != nil {

			if err := m.Occurrences[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *EventCreateBody) validatePublishedAt(formats strfmt.Registry) error {

	if err := validate.Required("published_at", "body", m.PublishedAt); err != nil {
		return err
	}

	return nil
}

func (m *EventCreateBody) validateRetailers(formats strfmt.Registry) error {

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
