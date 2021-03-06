package s_e_a_r_c_h

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetSearchEventsParams creates a new GetSearchEventsParams object
// with the default values initialized.
func NewGetSearchEventsParams() *GetSearchEventsParams {
	var (
		debugDefault             bool  = bool(false)
		limitDefault             int64 = int64(5)
		maxValuesPerFacetDefault int64 = int64(10)
		pageDefault              int64 = int64(1)
	)
	return &GetSearchEventsParams{
		Debug:             &debugDefault,
		Limit:             &limitDefault,
		MaxValuesPerFacet: &maxValuesPerFacetDefault,
		Page:              &pageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSearchEventsParamsWithTimeout creates a new GetSearchEventsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSearchEventsParamsWithTimeout(timeout time.Duration) *GetSearchEventsParams {
	var (
		debugDefault             bool  = bool(false)
		limitDefault             int64 = int64(5)
		maxValuesPerFacetDefault int64 = int64(10)
		pageDefault              int64 = int64(1)
	)
	return &GetSearchEventsParams{
		Debug:             &debugDefault,
		Limit:             &limitDefault,
		MaxValuesPerFacet: &maxValuesPerFacetDefault,
		Page:              &pageDefault,

		timeout: timeout,
	}
}

/*GetSearchEventsParams contains all the parameters to send to the API endpoint
for the get search events operation typically these are written to a http.Request
*/
type GetSearchEventsParams struct {

	/*CentreID
	  centre_id. Facet by centre_id.

	*/
	CentreID *string
	/*Country
	  country. Facet by country.

	*/
	Country *string
	/*Debug
	  Include debug information (such as ranking and analytics data) in the response

	*/
	Debug *bool
	/*ExcludeFromSearch
	  List of attributes you do not want to use for search

	*/
	ExcludeFromSearch []string
	/*Fields
	  Fields in 'hits' response. Array that lists the fields requested. If empty or not supplied, all attributes will be returned.

	*/
	Fields []string
	/*Limit
	  Limit. Maximum number of results to return per result type.

	*/
	Limit *int64
	/*MaxValuesPerFacet
	  Maximum Values Per Facet. Maximum number of values to return per facet.

	*/
	MaxValuesPerFacet *int64
	/*Page
	  Results page number. Specific page of results to request when paginating.

	*/
	Page *int64
	/*Term
	  Search term. Contents of search term.

	*/
	Term *string

	timeout time.Duration
}

// WithCentreID adds the centreId to the get search events params
func (o *GetSearchEventsParams) WithCentreID(CentreID *string) *GetSearchEventsParams {
	o.CentreID = CentreID
	return o
}

// WithCountry adds the country to the get search events params
func (o *GetSearchEventsParams) WithCountry(Country *string) *GetSearchEventsParams {
	o.Country = Country
	return o
}

// WithDebug adds the debug to the get search events params
func (o *GetSearchEventsParams) WithDebug(Debug *bool) *GetSearchEventsParams {
	o.Debug = Debug
	return o
}

// WithExcludeFromSearch adds the excludeFromSearch to the get search events params
func (o *GetSearchEventsParams) WithExcludeFromSearch(ExcludeFromSearch []string) *GetSearchEventsParams {
	o.ExcludeFromSearch = ExcludeFromSearch
	return o
}

// WithFields adds the fields to the get search events params
func (o *GetSearchEventsParams) WithFields(Fields []string) *GetSearchEventsParams {
	o.Fields = Fields
	return o
}

// WithLimit adds the limit to the get search events params
func (o *GetSearchEventsParams) WithLimit(Limit *int64) *GetSearchEventsParams {
	o.Limit = Limit
	return o
}

// WithMaxValuesPerFacet adds the maxValuesPerFacet to the get search events params
func (o *GetSearchEventsParams) WithMaxValuesPerFacet(MaxValuesPerFacet *int64) *GetSearchEventsParams {
	o.MaxValuesPerFacet = MaxValuesPerFacet
	return o
}

// WithPage adds the page to the get search events params
func (o *GetSearchEventsParams) WithPage(Page *int64) *GetSearchEventsParams {
	o.Page = Page
	return o
}

// WithTerm adds the term to the get search events params
func (o *GetSearchEventsParams) WithTerm(Term *string) *GetSearchEventsParams {
	o.Term = Term
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetSearchEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.CentreID != nil {

		// query param centre_id
		var qrCentreID string
		if o.CentreID != nil {
			qrCentreID = *o.CentreID
		}
		qCentreID := qrCentreID
		if qCentreID != "" {
			if err := r.SetQueryParam("centre_id", qCentreID); err != nil {
				return err
			}
		}

	}

	if o.Country != nil {

		// query param country
		var qrCountry string
		if o.Country != nil {
			qrCountry = *o.Country
		}
		qCountry := qrCountry
		if qCountry != "" {
			if err := r.SetQueryParam("country", qCountry); err != nil {
				return err
			}
		}

	}

	if o.Debug != nil {

		// query param debug
		var qrDebug bool
		if o.Debug != nil {
			qrDebug = *o.Debug
		}
		qDebug := swag.FormatBool(qrDebug)
		if qDebug != "" {
			if err := r.SetQueryParam("debug", qDebug); err != nil {
				return err
			}
		}

	}

	valuesExcludeFromSearch := o.ExcludeFromSearch

	joinedExcludeFromSearch := swag.JoinByFormat(valuesExcludeFromSearch, "csv")
	// query array param excludeFromSearch
	if err := r.SetQueryParam("excludeFromSearch", joinedExcludeFromSearch...); err != nil {
		return err
	}

	valuesFields := o.Fields

	joinedFields := swag.JoinByFormat(valuesFields, "csv")
	// query array param fields
	if err := r.SetQueryParam("fields", joinedFields...); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.MaxValuesPerFacet != nil {

		// query param maxValuesPerFacet
		var qrMaxValuesPerFacet int64
		if o.MaxValuesPerFacet != nil {
			qrMaxValuesPerFacet = *o.MaxValuesPerFacet
		}
		qMaxValuesPerFacet := swag.FormatInt64(qrMaxValuesPerFacet)
		if qMaxValuesPerFacet != "" {
			if err := r.SetQueryParam("maxValuesPerFacet", qMaxValuesPerFacet); err != nil {
				return err
			}
		}

	}

	if o.Page != nil {

		// query param page
		var qrPage int64
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.Term != nil {

		// query param term
		var qrTerm string
		if o.Term != nil {
			qrTerm = *o.Term
		}
		qTerm := qrTerm
		if qTerm != "" {
			if err := r.SetQueryParam("term", qTerm); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
