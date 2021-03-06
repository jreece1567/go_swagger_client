package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*CreateStoreHoursBody Parameters for creating one or more trading hours for a store.

swagger:model createStoreHoursBody
*/
type CreateStoreHoursBody struct {

	/* The centre identifer you'd like the hours for

	Required: true
	*/
	CentreID *string `json:"centre_id"`

	/* Whether the store is open or not
	 */
	Closed []string `json:"closed,omitempty"`

	/* Time the store closes
	 */
	ClosingTime []string `json:"closing_time,omitempty"`

	/* Date the trading hour is for
	 */
	Date []string `json:"date,omitempty"`

	/* 0-indexed day of the week (e.g. 0: Mon, 1: Tue)
	 */
	DayOfWeek []int64 `json:"day_of_week,omitempty"`

	/* Descriptions of trading hours
	 */
	Description []string `json:"description,omitempty"`

	/* Types of trading hours
	 */
	HourType []string `json:"hour_type,omitempty"`

	/* Time the store opens
	 */
	OpeningTime []string `json:"opening_time,omitempty"`
}

// Validate validates this create store hours body
func (m *CreateStoreHoursBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCentreID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateClosed(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateClosingTime(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDayOfWeek(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHourType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOpeningTime(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateStoreHoursBody) validateCentreID(formats strfmt.Registry) error {

	if err := validate.Required("centre_id", "body", m.CentreID); err != nil {
		return err
	}

	return nil
}

func (m *CreateStoreHoursBody) validateClosed(formats strfmt.Registry) error {

	if swag.IsZero(m.Closed) { // not required
		return nil
	}

	return nil
}

func (m *CreateStoreHoursBody) validateClosingTime(formats strfmt.Registry) error {

	if swag.IsZero(m.ClosingTime) { // not required
		return nil
	}

	return nil
}

func (m *CreateStoreHoursBody) validateDate(formats strfmt.Registry) error {

	if swag.IsZero(m.Date) { // not required
		return nil
	}

	return nil
}

func (m *CreateStoreHoursBody) validateDayOfWeek(formats strfmt.Registry) error {

	if swag.IsZero(m.DayOfWeek) { // not required
		return nil
	}

	return nil
}

func (m *CreateStoreHoursBody) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	return nil
}

func (m *CreateStoreHoursBody) validateHourType(formats strfmt.Registry) error {

	if swag.IsZero(m.HourType) { // not required
		return nil
	}

	return nil
}

func (m *CreateStoreHoursBody) validateOpeningTime(formats strfmt.Registry) error {

	if swag.IsZero(m.OpeningTime) { // not required
		return nil
	}

	return nil
}
