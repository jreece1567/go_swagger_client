package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

/*ParkingRates Parking rates

Each n-th hour range with parking rates

swagger:model parkingRates
*/
type ParkingRates struct {

	/* Starting hour range
	 */
	HourlyRateFrom string `json:"hourly_rate_from,omitempty"`

	/* Ending hour range
	 */
	HourlyRateTo string `json:"hourly_rate_to,omitempty"`

	/* Cost for this range on weekdays
	 */
	HourlyRateWeekday string `json:"hourly_rate_weekday,omitempty"`

	/* Cost for this range on weekends
	 */
	HourlyRateWeekend string `json:"hourly_rate_weekend,omitempty"`
}

// Validate validates this parking rates
func (m *ParkingRates) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
