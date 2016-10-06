package models_secure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

/*EdnaAppsInstance Apps instance for EDNA

swagger:model ednaAppsInstance
*/
type EdnaAppsInstance struct {

	/* Identifier of the app.
	 */
	AppID string `json:"app_id,omitempty"`

	/* Devices associated with the app.
	 */
	Devices []string `json:"devices,omitempty"`

	/* Number of notifications.
	 */
	NotificationBadge int64 `json:"notification_badge,omitempty"`
}

// Validate validates this edna apps instance
func (m *EdnaAppsInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDevices(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EdnaAppsInstance) validateDevices(formats strfmt.Registry) error {

	if swag.IsZero(m.Devices) { // not required
		return nil
	}

	return nil
}