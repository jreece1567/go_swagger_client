package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*DealStores Stores related to a deal.

swagger:model dealStores
*/
type DealStores struct {

	/* Centre identifier. Identifier to retrieve the centre this store is related to. Lowercase code name for a specific centre. Examples: ['sanfrancisco','sydney','london']

	Required: true
	*/
	CentreID *string `json:"centre_id"`

	/* Retailer identifier. Identifier to retrieve the retailer this store is related to.

	Required: true
	*/
	RetailerID *int64 `json:"retailer_id"`

	/* Store identifier. Identifier to retrieve the store the deal is related to.

	Required: true
	*/
	StoreID *int64 `json:"store_id"`
}

// Validate validates this deal stores
func (m *DealStores) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCentreID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRetailerID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStoreID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DealStores) validateCentreID(formats strfmt.Registry) error {

	if err := validate.Required("centre_id", "body", m.CentreID); err != nil {
		return err
	}

	return nil
}

func (m *DealStores) validateRetailerID(formats strfmt.Registry) error {

	if err := validate.Required("retailer_id", "body", m.RetailerID); err != nil {
		return err
	}

	return nil
}

func (m *DealStores) validateStoreID(formats strfmt.Registry) error {

	if err := validate.Required("store_id", "body", m.StoreID); err != nil {
		return err
	}

	return nil
}
