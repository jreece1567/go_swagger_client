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

/*ScheduleCreateBody Parameters for creating a schedule.

swagger:model scheduleCreateBody
*/
type ScheduleCreateBody struct {

	/* days of week
	 */
	DaysOfWeek map[string][]ScheduleEvent `json:"days_of_week,omitempty"`

	/* Date-time the schedule starts.
	 */
	EndsAt strfmt.DateTime `json:"ends_at,omitempty"`

	/* events
	 */
	Events []*ScheduleEvent `json:"events,omitempty"`

	/* Schedule name.

	Required: true
	*/
	Name *string `json:"name"`

	/* Schedule participant identifier.

	Required: true
	*/
	ParticipantID *string `json:"participant_id"`

	/* Type of participant. Used with participant_kind_id.
	 */
	ParticipantKind string `json:"participant_kind,omitempty"`

	/* Participant kind identifier. Identifier for participant of type participant_kind.
	 */
	ParticipantKindID string `json:"participant_kind_id,omitempty"`

	/* Date-time the schedule ends.
	 */
	StartsAt strfmt.DateTime `json:"starts_at,omitempty"`

	/* Schedule's time zone.

	Required: true
	*/
	TimeZone *string `json:"time_zone"`

	/* Type of schedule.
	 */
	Type string `json:"type,omitempty"`
}

// Validate validates this schedule create body
func (m *ScheduleCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDaysOfWeek(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEvents(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateParticipantID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateParticipantKind(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTimeZone(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScheduleCreateBody) validateDaysOfWeek(formats strfmt.Registry) error {

	if swag.IsZero(m.DaysOfWeek) { // not required
		return nil
	}

	if err := validate.Required("days_of_week", "body", m.DaysOfWeek); err != nil {
		return err
	}

	return nil
}

func (m *ScheduleCreateBody) validateEvents(formats strfmt.Registry) error {

	if swag.IsZero(m.Events) { // not required
		return nil
	}

	for i := 0; i < len(m.Events); i++ {

		if swag.IsZero(m.Events[i]) { // not required
			continue
		}

		if m.Events[i] != nil {

			if err := m.Events[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ScheduleCreateBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ScheduleCreateBody) validateParticipantID(formats strfmt.Registry) error {

	if err := validate.Required("participant_id", "body", m.ParticipantID); err != nil {
		return err
	}

	return nil
}

var scheduleCreateBodyTypeParticipantKindPropEnum []interface{}

// prop value enum
func (m *ScheduleCreateBody) validateParticipantKindEnum(path, location string, value string) error {
	if scheduleCreateBodyTypeParticipantKindPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Centre","Store"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			scheduleCreateBodyTypeParticipantKindPropEnum = append(scheduleCreateBodyTypeParticipantKindPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, scheduleCreateBodyTypeParticipantKindPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ScheduleCreateBody) validateParticipantKind(formats strfmt.Registry) error {

	if swag.IsZero(m.ParticipantKind) { // not required
		return nil
	}

	// value enum
	if err := m.validateParticipantKindEnum("participant_kind", "body", m.ParticipantKind); err != nil {
		return err
	}

	return nil
}

func (m *ScheduleCreateBody) validateTimeZone(formats strfmt.Registry) error {

	if err := validate.Required("time_zone", "body", m.TimeZone); err != nil {
		return err
	}

	return nil
}

var scheduleCreateBodyTypeTypePropEnum []interface{}

// prop value enum
func (m *ScheduleCreateBody) validateTypeEnum(path, location string, value string) error {
	if scheduleCreateBodyTypeTypePropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["special_trading","standard_trading"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			scheduleCreateBodyTypeTypePropEnum = append(scheduleCreateBodyTypeTypePropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, scheduleCreateBodyTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ScheduleCreateBody) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}
