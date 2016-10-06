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

/*KioskCentreCreateBody Data to create a kiosk centre.

swagger:model kioskCentreCreateBody
*/
type KioskCentreCreateBody struct {

	/* category ids
	 */
	CategoryIds []int64 `json:"category_ids,omitempty"`

	/* Centre identifier.

	Required: true
	*/
	CentreID *string `json:"centre_id"`

	/* features
	 */
	Features []string `json:"features,omitempty"`

	/* Kiosk centre landing page.
	 */
	LandingPage string `json:"landing_page,omitempty"`

	/* languages
	 */
	Languages []string `json:"languages,omitempty"`
}

// Validate validates this kiosk centre create body
func (m *KioskCentreCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategoryIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCentreID(formats); err != nil {
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

func (m *KioskCentreCreateBody) validateCategoryIds(formats strfmt.Registry) error {

	if swag.IsZero(m.CategoryIds) { // not required
		return nil
	}

	return nil
}

func (m *KioskCentreCreateBody) validateCentreID(formats strfmt.Registry) error {

	if err := validate.Required("centre_id", "body", m.CentreID); err != nil {
		return err
	}

	return nil
}

func (m *KioskCentreCreateBody) validateFeatures(formats strfmt.Registry) error {

	if swag.IsZero(m.Features) { // not required
		return nil
	}

	return nil
}

var kioskCentreCreateBodyLanguagesItemsEnum []interface{}

func (m *KioskCentreCreateBody) validateLanguagesItemsEnum(path, location string, value string) error {
	if kioskCentreCreateBodyLanguagesItemsEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["en_US","en_GB","es_ES","fr_FR","it_IT","de_DE","pt_PT","ja_JP","zh_CN","ar_SA","he_IL","ru_RU"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			kioskCentreCreateBodyLanguagesItemsEnum = append(kioskCentreCreateBodyLanguagesItemsEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, kioskCentreCreateBodyLanguagesItemsEnum); err != nil {
		return err
	}
	return nil
}

func (m *KioskCentreCreateBody) validateLanguages(formats strfmt.Registry) error {

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
