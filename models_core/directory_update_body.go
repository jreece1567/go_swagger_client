package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*DirectoryUpdateBody Data to update a directory.

swagger:model directoryUpdateBody
*/
type DirectoryUpdateBody struct {

	/* Category identifiers. List of identifiers to retrieve categories related to the centre identified by centre_id.
	 */
	CategoryIds []int64 `json:"category_ids,omitempty"`

	/* Centre identifier. Identifier to retrieve the centre related to this kiosk_centre directory.
	 */
	CentreID string `json:"centre_id,omitempty"`

	/* Directory is enabled (true) or not enabled (false).
	 */
	Enabled bool `json:"enabled,omitempty"`

	/* Enclosure identifier. Identifier to retrieve the enclosure related to this kiosk_centre directory.
	 */
	EnclosureID int64 `json:"enclosure_id,omitempty"`

	/* List of features available on this kiosk_centre directory for the centre identified by centre_id.
	 */
	Features []string `json:"features,omitempty"`

	/* Directory heading number.
	 */
	Heading int64 `json:"heading,omitempty"`

	/* Landing page for this kiosk_centre directory.
	 */
	LandingPage string `json:"landing_page,omitempty"`

	/* List of languages available on this kiosk_centre directory.
	 */
	Languages []string `json:"languages,omitempty"`

	/* Directory name.
	 */
	Name string `json:"name,omitempty"`

	/* Directory template.
	 */
	Template string `json:"template,omitempty"`

	/* Waypoint identifier. Identifier to retrieve the waypoint related to the enclosure for this kiosk-centre directory.
	 */
	WaypointID string `json:"waypoint_id,omitempty"`
}

// Validate validates this directory update body
func (m *DirectoryUpdateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategoryIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFeatures(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLanguages(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectoryUpdateBody) validateCategoryIds(formats strfmt.Registry) error {

	if swag.IsZero(m.CategoryIds) { // not required
		return nil
	}

	return nil
}

func (m *DirectoryUpdateBody) validateFeatures(formats strfmt.Registry) error {

	if swag.IsZero(m.Features) { // not required
		return nil
	}

	return nil
}

var directoryUpdateBodyLanguagesItemsEnum []interface{}

func (m *DirectoryUpdateBody) validateLanguagesItemsEnum(path, location string, value string) error {
	if directoryUpdateBodyLanguagesItemsEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["en_US","en_GB","es_ES","fr_FR","it_IT","de_DE","pt_PT","ja_JP","zh_CN","ar_SA","he_IL","ru_RU"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			directoryUpdateBodyLanguagesItemsEnum = append(directoryUpdateBodyLanguagesItemsEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, directoryUpdateBodyLanguagesItemsEnum); err != nil {
		return err
	}
	return nil
}

func (m *DirectoryUpdateBody) validateLanguages(formats strfmt.Registry) error {

	if swag.IsZero(m.Languages) { // not required
		return nil
	}

	for i := 0; i < len(m.Languages); i++ {

		// value enum
		if err := m.validateLanguagesItemsEnum("languages"+"."+strconv.Itoa(i), "body", m.Languages[i]); err != nil {
			return err
		}

	}

	return nil
}
