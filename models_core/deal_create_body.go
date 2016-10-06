package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*DealCreateBody Data to create a deal.

swagger:model dealCreateBody
*/
type DealCreateBody struct {

	/* Campaign identifiers. List of identifiers to retrieve the campaigns related to this deal.
	 */
	CampaignIds []int64 `json:"campaign_ids,omitempty"`

	/* Category identifiers. List of identifiers to retrieve the categories related to this deal.
	 */
	CategoryIds []int64 `json:"category_ids,omitempty"`

	/* comments
	 */
	Comments *DealComment `json:"comments,omitempty"`

	/* Description of deal.
	 */
	Description string `json:"description,omitempty"`

	/* Date-time that the deal ends.

	Required: true
	*/
	EndsAt *strfmt.DateTime `json:"ends_at"`

	/* Featured deals. Deal is featured (true) or not featured (false) deals. Default is false.
	 */
	Featured bool `json:"featured,omitempty"`

	/* Image reference use internally to retrieve the related image file.
	 */
	ImageRef string `json:"image_ref,omitempty"`

	/* Date-time the deal is published.

	Required: true
	*/
	PublishedAt *strfmt.DateTime `json:"published_at"`

	/* Date-time the deal starts.

	Required: true
	*/
	StartsAt *strfmt.DateTime `json:"starts_at"`

	/* Event to set the state of the deal.
	 */
	StateEvent string `json:"state_event,omitempty"`

	/* stores

	Required: true
	*/
	Stores *DealStores `json:"stores"`

	/* Subtitle of the deal.
	 */
	Subtitle string `json:"subtitle,omitempty"`

	/* Terms and conditions of the deal.

	Required: true
	*/
	TermsAndConditions *string `json:"terms_and_conditions"`

	/* Title of the deal.

	Required: true
	Max Length: 100
	*/
	Title *string `json:"title"`
}

// Validate validates this deal create body
func (m *DealCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCampaignIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCategoryIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateComments(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEndsAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePublishedAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStartsAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStateEvent(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStores(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTermsAndConditions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DealCreateBody) validateCampaignIds(formats strfmt.Registry) error {

	if swag.IsZero(m.CampaignIds) { // not required
		return nil
	}

	return nil
}

func (m *DealCreateBody) validateCategoryIds(formats strfmt.Registry) error {

	if swag.IsZero(m.CategoryIds) { // not required
		return nil
	}

	return nil
}

func (m *DealCreateBody) validateComments(formats strfmt.Registry) error {

	if swag.IsZero(m.Comments) { // not required
		return nil
	}

	if m.Comments != nil {

		if err := m.Comments.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *DealCreateBody) validateEndsAt(formats strfmt.Registry) error {

	if err := validate.Required("ends_at", "body", m.EndsAt); err != nil {
		return err
	}

	return nil
}

func (m *DealCreateBody) validatePublishedAt(formats strfmt.Registry) error {

	if err := validate.Required("published_at", "body", m.PublishedAt); err != nil {
		return err
	}

	return nil
}

func (m *DealCreateBody) validateStartsAt(formats strfmt.Registry) error {

	if err := validate.Required("starts_at", "body", m.StartsAt); err != nil {
		return err
	}

	return nil
}

var dealCreateBodyTypeStateEventPropEnum []interface{}

// prop value enum
func (m *DealCreateBody) validateStateEventEnum(path, location string, value string) error {
	if dealCreateBodyTypeStateEventPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["submit","approve","reject","withdraw"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			dealCreateBodyTypeStateEventPropEnum = append(dealCreateBodyTypeStateEventPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, dealCreateBodyTypeStateEventPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DealCreateBody) validateStateEvent(formats strfmt.Registry) error {

	if swag.IsZero(m.StateEvent) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEventEnum("state_event", "body", m.StateEvent); err != nil {
		return err
	}

	return nil
}

func (m *DealCreateBody) validateStores(formats strfmt.Registry) error {

	if err := validate.Required("stores", "body", m.Stores); err != nil {
		return err
	}

	if m.Stores != nil {

		if err := m.Stores.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *DealCreateBody) validateTermsAndConditions(formats strfmt.Registry) error {

	if err := validate.Required("terms_and_conditions", "body", m.TermsAndConditions); err != nil {
		return err
	}

	return nil
}

func (m *DealCreateBody) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	if err := validate.MaxLength("title", "body", string(*m.Title), 100); err != nil {
		return err
	}

	return nil
}
