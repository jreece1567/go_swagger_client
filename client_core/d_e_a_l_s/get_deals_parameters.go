package d_e_a_l_s

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

// NewGetDealsParams creates a new GetDealsParams object
// with the default values initialized.
func NewGetDealsParams() *GetDealsParams {
	var (
		allDefault     bool     = bool(false)
		deletedDefault bool     = bool(false)
		pageDefault    int64    = int64(1)
		perPageDefault int64    = int64(25)
		sortDefault    []string = []string{ "featured desc" }
	)
	return &GetDealsParams{
		All:     &allDefault,
		Deleted: &deletedDefault,
		Page:    &pageDefault,
		PerPage: &perPageDefault,
		Sort:    sortDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDealsParamsWithTimeout creates a new GetDealsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDealsParamsWithTimeout(timeout time.Duration) *GetDealsParams {
	var (
		allDefault     bool     = bool(false)
		deletedDefault bool     = bool(false)
		pageDefault    int64    = int64(1)
		perPageDefault int64    = int64(25)
		sortDefault    []string = []string{ "featured desc" }
	)
	return &GetDealsParams{
		All:     &allDefault,
		Deleted: &deletedDefault,
		Page:    &pageDefault,
		PerPage: &perPageDefault,
		Sort:    sortDefault,

		timeout: timeout,
	}
}

/*GetDealsParams contains all the parameters to send to the API endpoint
for the get deals operation typically these are written to a http.Request
*/
type GetDealsParams struct {

	/*All
	  Request all deals or published through available deals.

	*/
	All *bool
	/*CampaignCode
	  Campaign code. Request deals with campaign_code.

	*/
	CampaignCode *string
	/*CampaignID
	  Campaign identifier. Request deals with campaign_id.

	*/
	CampaignID *int64
	/*CentreID
	  Centre identifier. Request deals with centre_id. Lowercase code name for a specific centre. Must relate to store_id. Examples: ['sanfrancisco','sydney','london']

	*/
	CentreID *string
	/*Country
	  Country code. Request deals with country code. Lowercase two-character code for country. Current valid values are: au, nz, uk or us.

	*/
	Country *string
	/*DealIds
	  List of deal ids to be used for additional filtering of results.

	*/
	DealIds []int64
	/*Deleted
	  Deleted deals. Request to include (true) or not include (false) deals that are deleted.

	*/
	Deleted *bool
	/*Featured
	  Featured deals. Request to include only featured (true) or not featured (false) deals. Default to return all deals.

	*/
	Featured *bool
	/*Fields
	  Fields in response. Array that lists the fields requested. If empty or not supplied, all attributes will be returned.

	*/
	Fields []string
	/*Page
	  Results page number. Specific page of results to request when paginating.

	*/
	Page *int64
	/*PerPage
	  Results per page. Number of items per result set when paginating.

	*/
	PerPage *int64
	/*RetailerID
	  Retailer identifier. Request the deals with retailer_id.

	*/
	RetailerID *int64
	/*Sort
	  Sort results by one or more fields. Change sort order by adding desc or asc after field name.

	*/
	Sort []string
	/*State
	  State. Request deals with state.

	*/
	State *string
	/*Statuses
	  Statuses. Requests deals that include one of the listed statuses.

	*/
	Statuses []string
	/*StoreID
	  Store identifier. Request the deals with store_id.

	*/
	StoreID *int64
	/*UpdatedSince
	  Updated since. Request the deals updated since a specific date and time. ISO-8601 format.

	*/
	UpdatedSince *strfmt.DateTime

	timeout time.Duration
}

// WithAll adds the all to the get deals params
func (o *GetDealsParams) WithAll(All *bool) *GetDealsParams {
	o.All = All
	return o
}

// WithCampaignCode adds the campaignCode to the get deals params
func (o *GetDealsParams) WithCampaignCode(CampaignCode *string) *GetDealsParams {
	o.CampaignCode = CampaignCode
	return o
}

// WithCampaignID adds the campaignId to the get deals params
func (o *GetDealsParams) WithCampaignID(CampaignID *int64) *GetDealsParams {
	o.CampaignID = CampaignID
	return o
}

// WithCentreID adds the centreId to the get deals params
func (o *GetDealsParams) WithCentreID(CentreID *string) *GetDealsParams {
	o.CentreID = CentreID
	return o
}

// WithCountry adds the country to the get deals params
func (o *GetDealsParams) WithCountry(Country *string) *GetDealsParams {
	o.Country = Country
	return o
}

// WithDealIds adds the dealIds to the get deals params
func (o *GetDealsParams) WithDealIds(DealIds []int64) *GetDealsParams {
	o.DealIds = DealIds
	return o
}

// WithDeleted adds the deleted to the get deals params
func (o *GetDealsParams) WithDeleted(Deleted *bool) *GetDealsParams {
	o.Deleted = Deleted
	return o
}

// WithFeatured adds the featured to the get deals params
func (o *GetDealsParams) WithFeatured(Featured *bool) *GetDealsParams {
	o.Featured = Featured
	return o
}

// WithFields adds the fields to the get deals params
func (o *GetDealsParams) WithFields(Fields []string) *GetDealsParams {
	o.Fields = Fields
	return o
}

// WithPage adds the page to the get deals params
func (o *GetDealsParams) WithPage(Page *int64) *GetDealsParams {
	o.Page = Page
	return o
}

// WithPerPage adds the perPage to the get deals params
func (o *GetDealsParams) WithPerPage(PerPage *int64) *GetDealsParams {
	o.PerPage = PerPage
	return o
}

// WithRetailerID adds the retailerId to the get deals params
func (o *GetDealsParams) WithRetailerID(RetailerID *int64) *GetDealsParams {
	o.RetailerID = RetailerID
	return o
}

// WithSort adds the sort to the get deals params
func (o *GetDealsParams) WithSort(Sort []string) *GetDealsParams {
	o.Sort = Sort
	return o
}

// WithState adds the state to the get deals params
func (o *GetDealsParams) WithState(State *string) *GetDealsParams {
	o.State = State
	return o
}

// WithStatuses adds the statuses to the get deals params
func (o *GetDealsParams) WithStatuses(Statuses []string) *GetDealsParams {
	o.Statuses = Statuses
	return o
}

// WithStoreID adds the storeId to the get deals params
func (o *GetDealsParams) WithStoreID(StoreID *int64) *GetDealsParams {
	o.StoreID = StoreID
	return o
}

// WithUpdatedSince adds the updatedSince to the get deals params
func (o *GetDealsParams) WithUpdatedSince(UpdatedSince *strfmt.DateTime) *GetDealsParams {
	o.UpdatedSince = UpdatedSince
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetDealsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.All != nil {

		// query param all
		var qrAll bool
		if o.All != nil {
			qrAll = *o.All
		}
		qAll := swag.FormatBool(qrAll)
		if qAll != "" {
			if err := r.SetQueryParam("all", qAll); err != nil {
				return err
			}
		}

	}

	if o.CampaignCode != nil {

		// query param campaign_code
		var qrCampaignCode string
		if o.CampaignCode != nil {
			qrCampaignCode = *o.CampaignCode
		}
		qCampaignCode := qrCampaignCode
		if qCampaignCode != "" {
			if err := r.SetQueryParam("campaign_code", qCampaignCode); err != nil {
				return err
			}
		}

	}

	if o.CampaignID != nil {

		// query param campaign_id
		var qrCampaignID int64
		if o.CampaignID != nil {
			qrCampaignID = *o.CampaignID
		}
		qCampaignID := swag.FormatInt64(qrCampaignID)
		if qCampaignID != "" {
			if err := r.SetQueryParam("campaign_id", qCampaignID); err != nil {
				return err
			}
		}

	}

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

	var valuesDealIds []string
	for _, v := range o.DealIds {
		valuesDealIds = append(valuesDealIds, swag.FormatInt64(v))
	}

	joinedDealIds := swag.JoinByFormat(valuesDealIds, "csv")
	// query array param deal_ids
	if err := r.SetQueryParam("deal_ids", joinedDealIds...); err != nil {
		return err
	}

	if o.Deleted != nil {

		// query param deleted
		var qrDeleted bool
		if o.Deleted != nil {
			qrDeleted = *o.Deleted
		}
		qDeleted := swag.FormatBool(qrDeleted)
		if qDeleted != "" {
			if err := r.SetQueryParam("deleted", qDeleted); err != nil {
				return err
			}
		}

	}

	if o.Featured != nil {

		// query param featured
		var qrFeatured bool
		if o.Featured != nil {
			qrFeatured = *o.Featured
		}
		qFeatured := swag.FormatBool(qrFeatured)
		if qFeatured != "" {
			if err := r.SetQueryParam("featured", qFeatured); err != nil {
				return err
			}
		}

	}

	valuesFields := o.Fields

	joinedFields := swag.JoinByFormat(valuesFields, "csv")
	// query array param fields
	if err := r.SetQueryParam("fields", joinedFields...); err != nil {
		return err
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

	if o.PerPage != nil {

		// query param per_page
		var qrPerPage int64
		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt64(qrPerPage)
		if qPerPage != "" {
			if err := r.SetQueryParam("per_page", qPerPage); err != nil {
				return err
			}
		}

	}

	if o.RetailerID != nil {

		// query param retailer_id
		var qrRetailerID int64
		if o.RetailerID != nil {
			qrRetailerID = *o.RetailerID
		}
		qRetailerID := swag.FormatInt64(qrRetailerID)
		if qRetailerID != "" {
			if err := r.SetQueryParam("retailer_id", qRetailerID); err != nil {
				return err
			}
		}

	}

	valuesSort := o.Sort

	joinedSort := swag.JoinByFormat(valuesSort, "csv")
	// query array param sort
	if err := r.SetQueryParam("sort", joinedSort...); err != nil {
		return err
	}

	if o.State != nil {

		// query param state
		var qrState string
		if o.State != nil {
			qrState = *o.State
		}
		qState := qrState
		if qState != "" {
			if err := r.SetQueryParam("state", qState); err != nil {
				return err
			}
		}

	}

	valuesStatuses := o.Statuses

	joinedStatuses := swag.JoinByFormat(valuesStatuses, "csv")
	// query array param statuses
	if err := r.SetQueryParam("statuses", joinedStatuses...); err != nil {
		return err
	}

	if o.StoreID != nil {

		// query param store_id
		var qrStoreID int64
		if o.StoreID != nil {
			qrStoreID = *o.StoreID
		}
		qStoreID := swag.FormatInt64(qrStoreID)
		if qStoreID != "" {
			if err := r.SetQueryParam("store_id", qStoreID); err != nil {
				return err
			}
		}

	}

	if o.UpdatedSince != nil {

		// query param updated_since
		var qrUpdatedSince strfmt.DateTime
		if o.UpdatedSince != nil {
			qrUpdatedSince = *o.UpdatedSince
		}
		qUpdatedSince := qrUpdatedSince.String()
		if qUpdatedSince != "" {
			if err := r.SetQueryParam("updated_since", qUpdatedSince); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
