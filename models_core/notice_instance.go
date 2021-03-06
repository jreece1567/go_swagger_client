package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*NoticeInstance Notice record

swagger:model noticeInstance
*/
type NoticeInstance struct {

	/* Hash of links

	Required: true
	*/
	Links *NoticeLinks `json:"_links"`

	/* Centre identifier

	Required: true
	*/
	CentreID *string `json:"centre_id"`

	/* Deleted date
	 */
	DeletedAt strfmt.DateTime `json:"deleted_at,omitempty"`

	/* Notice content text

	Required: true
	Min Length: 2
	*/
	Detail *string `json:"detail"`

	/* Notice expiry date

	Required: true
	*/
	ExpiresAt *strfmt.DateTime `json:"expires_at"`

	/* Notice is featured (true/false)

	Required: true
	*/
	Featured *bool `json:"featured"`

	/* Image alternative text

	Required: true
	*/
	ImgAltText *string `json:"img_alt_text"`

	/* Notice name

	Required: true
	Max Length: 40
	Min Length: 2
	*/
	Name *string `json:"name"`

	/* Notice identifier

	Required: true
	*/
	NoticeID *int64 `json:"notice_id"`

	/* Notice publication date

	Required: true
	*/
	PublishedAt *strfmt.DateTime `json:"published_at"`

	/* Value indicating importance of notice

	Required: true
	*/
	Salience *int64 `json:"salience"`

	/* Type

	Required: true
	*/
	Type *string `json:"type"`

	/* Update date

	Required: true
	*/
	UpdatedAt *strfmt.DateTime `json:"updated_at"`
}

// Validate validates this notice instance
func (m *NoticeInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCentreID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDetail(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateExpiresAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFeatured(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateImgAltText(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNoticeID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePublishedAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSalience(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NoticeInstance) validateLinks(formats strfmt.Registry) error {

	if err := validate.Required("_links", "body", m.Links); err != nil {
		return err
	}

	if m.Links != nil {

		if err := m.Links.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *NoticeInstance) validateCentreID(formats strfmt.Registry) error {

	if err := validate.Required("centre_id", "body", m.CentreID); err != nil {
		return err
	}

	return nil
}

func (m *NoticeInstance) validateDetail(formats strfmt.Registry) error {

	if err := validate.Required("detail", "body", m.Detail); err != nil {
		return err
	}

	if err := validate.MinLength("detail", "body", string(*m.Detail), 2); err != nil {
		return err
	}

	return nil
}

func (m *NoticeInstance) validateExpiresAt(formats strfmt.Registry) error {

	if err := validate.Required("expires_at", "body", m.ExpiresAt); err != nil {
		return err
	}

	return nil
}

func (m *NoticeInstance) validateFeatured(formats strfmt.Registry) error {

	if err := validate.Required("featured", "body", m.Featured); err != nil {
		return err
	}

	return nil
}

func (m *NoticeInstance) validateImgAltText(formats strfmt.Registry) error {

	if err := validate.Required("img_alt_text", "body", m.ImgAltText); err != nil {
		return err
	}

	return nil
}

func (m *NoticeInstance) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 2); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 40); err != nil {
		return err
	}

	return nil
}

func (m *NoticeInstance) validateNoticeID(formats strfmt.Registry) error {

	if err := validate.Required("notice_id", "body", m.NoticeID); err != nil {
		return err
	}

	return nil
}

func (m *NoticeInstance) validatePublishedAt(formats strfmt.Registry) error {

	if err := validate.Required("published_at", "body", m.PublishedAt); err != nil {
		return err
	}

	return nil
}

func (m *NoticeInstance) validateSalience(formats strfmt.Registry) error {

	if err := validate.Required("salience", "body", m.Salience); err != nil {
		return err
	}

	return nil
}

var noticeInstanceTypeTypePropEnum []interface{}

// prop value enum
func (m *NoticeInstance) validateTypeEnum(path, location string, value string) error {
	if noticeInstanceTypeTypePropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["standard","emergency"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			noticeInstanceTypeTypePropEnum = append(noticeInstanceTypeTypePropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, noticeInstanceTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NoticeInstance) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *NoticeInstance) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	return nil
}
