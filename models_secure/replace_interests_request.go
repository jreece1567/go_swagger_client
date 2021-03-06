package models_secure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*ReplaceInterestsRequest Data to replace interest list of account

swagger:model replaceInterestsRequest
*/
type ReplaceInterestsRequest struct {

	/* Interest identitfiers. A list of unique identitfiers for the interests

	Required: true
	*/
	InterestIds []string `json:"interest_ids"`

	/* Source from which the interests are being updated

	Required: true
	*/
	InterestsUpdatedFrom *string `json:"interests_updated_from"`
}

// Validate validates this replace interests request
func (m *ReplaceInterestsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInterestIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInterestsUpdatedFrom(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReplaceInterestsRequest) validateInterestIds(formats strfmt.Registry) error {

	if err := validate.Required("interest_ids", "body", m.InterestIds); err != nil {
		return err
	}

	return nil
}

var replaceInterestsRequestTypeInterestsUpdatedFromPropEnum []interface{}

// prop value enum
func (m *ReplaceInterestsRequest) validateInterestsUpdatedFromEnum(path, location string, value string) error {
	if replaceInterestsRequestTypeInterestsUpdatedFromPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["android","ios","mobile_web","web","concierge"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			replaceInterestsRequestTypeInterestsUpdatedFromPropEnum = append(replaceInterestsRequestTypeInterestsUpdatedFromPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, replaceInterestsRequestTypeInterestsUpdatedFromPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ReplaceInterestsRequest) validateInterestsUpdatedFrom(formats strfmt.Registry) error {

	if err := validate.Required("interests_updated_from", "body", m.InterestsUpdatedFrom); err != nil {
		return err
	}

	// value enum
	if err := m.validateInterestsUpdatedFromEnum("interests_updated_from", "body", *m.InterestsUpdatedFrom); err != nil {
		return err
	}

	return nil
}
