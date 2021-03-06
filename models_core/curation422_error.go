package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

/*Curation422Error List of errors returned while processing the request

swagger:model curation422Error
*/
type Curation422Error struct {

	/* banner narrow image ref
	 */
	BannerNarrowImageRef ValidationError `json:"banner_narrow_image_ref,omitempty"`

	/* banner wide image ref
	 */
	BannerWideImageRef ValidationError `json:"banner_wide_image_ref,omitempty"`

	/* A list of errors that do not apply to an attribute.
	 */
	Base ValidationError `json:"base,omitempty"`

	/* Errors with `category_ids` attribute
	 */
	CategoryIds []int64 `json:"category_ids,omitempty"`

	/* centre ids
	 */
	CentreIds ValidationError `json:"centre_ids,omitempty"`

	/* code
	 */
	Code ValidationError `json:"code,omitempty"`

	/* deleted at
	 */
	DeletedAt ValidationError `json:"deleted_at,omitempty"`

	/* ends at
	 */
	EndsAt ValidationError `json:"ends_at,omitempty"`

	/* image ref
	 */
	ImageRef ValidationError `json:"image_ref,omitempty"`

	/* name
	 */
	Name ValidationError `json:"name,omitempty"`

	/* Errors with `product_category_ids` attribute
	 */
	ProductCategoryIds []int64 `json:"product_category_ids,omitempty"`

	/* product ids
	 */
	ProductIds ValidationError `json:"product_ids,omitempty"`

	/* seo description
	 */
	SeoDescription ValidationError `json:"seo_description,omitempty"`

	/* seo title
	 */
	SeoTitle ValidationError `json:"seo_title,omitempty"`

	/* starts at
	 */
	StartsAt ValidationError `json:"starts_at,omitempty"`

	/* state
	 */
	State ValidationError `json:"state,omitempty"`
}

// Validate validates this curation422 error
func (m *Curation422Error) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategoryIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProductCategoryIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Curation422Error) validateCategoryIds(formats strfmt.Registry) error {

	if swag.IsZero(m.CategoryIds) { // not required
		return nil
	}

	return nil
}

func (m *Curation422Error) validateProductCategoryIds(formats strfmt.Registry) error {

	if swag.IsZero(m.ProductCategoryIds) { // not required
		return nil
	}

	return nil
}
