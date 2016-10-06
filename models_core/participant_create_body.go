package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*ParticipantCreateBody Parameters for creating a participant.

swagger:model participantCreateBody
*/
type ParticipantCreateBody struct {

	/* Participant type

	Required: true
	*/
	Kind *string `json:"kind"`

	/* Identifier for participant.

	Required: true
	*/
	KindID *string `json:"kind_id"`
}

// Validate validates this participant create body
func (m *ParticipantCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKind(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateKindID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var participantCreateBodyTypeKindPropEnum []interface{}

// prop value enum
func (m *ParticipantCreateBody) validateKindEnum(path, location string, value string) error {
	if participantCreateBodyTypeKindPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Centre","Store"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			participantCreateBodyTypeKindPropEnum = append(participantCreateBodyTypeKindPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, participantCreateBodyTypeKindPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ParticipantCreateBody) validateKind(formats strfmt.Registry) error {

	if err := validate.Required("kind", "body", m.Kind); err != nil {
		return err
	}

	// value enum
	if err := m.validateKindEnum("kind", "body", *m.Kind); err != nil {
		return err
	}

	return nil
}

func (m *ParticipantCreateBody) validateKindID(formats strfmt.Registry) error {

	if err := validate.Required("kind_id", "body", m.KindID); err != nil {
		return err
	}

	return nil
}
