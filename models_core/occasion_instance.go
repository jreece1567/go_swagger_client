package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*OccasionInstance Occasion instance.

swagger:model occasionInstance
*/
type OccasionInstance struct {

	/* Centre identifier. Identifier to retrieve the centre this occasion is related to. Lowercase code name for a specific centre. Examples: ['sanfrancisco','sydney','london'].

	Required: true
	*/
	CentreID *string `json:"centre_id"`

	/* Date-time the occasion was marked as deleted.

	Required: true
	*/
	DeletedAt *strfmt.DateTime `json:"deleted_at"`

	/* Date-time the occasion ends.

	Required: true
	*/
	EndsAt *strfmt.DateTime `json:"ends_at"`

	/* Occasion name.

	Required: true
	*/
	Name *string `json:"name"`

	/* Occasion identifier. Identifier to retrieve this occasion.

	Required: true
	*/
	OccasionID *int64 `json:"occasion_id"`

	/* Schedule identifiers. List of schedules to retrieve the schedules related to this occasion.

	Required: true
	*/
	ScheduleIds []int64 `json:"schedule_ids"`

	/* Date-time the occasion starts.

	Required: true
	*/
	StartsAt *strfmt.DateTime `json:"starts_at"`
}

// Validate validates this occasion instance
func (m *OccasionInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCentreID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDeletedAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEndsAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOccasionID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateScheduleIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStartsAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OccasionInstance) validateCentreID(formats strfmt.Registry) error {

	if err := validate.Required("centre_id", "body", m.CentreID); err != nil {
		return err
	}

	return nil
}

func (m *OccasionInstance) validateDeletedAt(formats strfmt.Registry) error {

	if err := validate.Required("deleted_at", "body", m.DeletedAt); err != nil {
		return err
	}

	return nil
}

func (m *OccasionInstance) validateEndsAt(formats strfmt.Registry) error {

	if err := validate.Required("ends_at", "body", m.EndsAt); err != nil {
		return err
	}

	return nil
}

func (m *OccasionInstance) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *OccasionInstance) validateOccasionID(formats strfmt.Registry) error {

	if err := validate.Required("occasion_id", "body", m.OccasionID); err != nil {
		return err
	}

	return nil
}

func (m *OccasionInstance) validateScheduleIds(formats strfmt.Registry) error {

	if err := validate.Required("schedule_ids", "body", m.ScheduleIds); err != nil {
		return err
	}

	return nil
}

func (m *OccasionInstance) validateStartsAt(formats strfmt.Registry) error {

	if err := validate.Required("starts_at", "body", m.StartsAt); err != nil {
		return err
	}

	return nil
}