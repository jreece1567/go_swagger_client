package m_o_v_i_e_i_n_f_o_r_m_a_t_i_o_n

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

// NewGetCentresCentreIDMoviesMovieIDParams creates a new GetCentresCentreIDMoviesMovieIDParams object
// with the default values initialized.
func NewGetCentresCentreIDMoviesMovieIDParams() *GetCentresCentreIDMoviesMovieIDParams {
	var ()
	return &GetCentresCentreIDMoviesMovieIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCentresCentreIDMoviesMovieIDParamsWithTimeout creates a new GetCentresCentreIDMoviesMovieIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCentresCentreIDMoviesMovieIDParamsWithTimeout(timeout time.Duration) *GetCentresCentreIDMoviesMovieIDParams {
	var ()
	return &GetCentresCentreIDMoviesMovieIDParams{

		timeout: timeout,
	}
}

/*GetCentresCentreIDMoviesMovieIDParams contains all the parameters to send to the API endpoint
for the get centres centre ID movies movie ID operation typically these are written to a http.Request
*/
type GetCentresCentreIDMoviesMovieIDParams struct {

	/*CentreID
	  Centre identifier. Request the movie with centre_id. Lowercase code name for a specific centre.

	*/
	CentreID string
	/*EndDate
	  End date. Request the sessions of this movie playing on or before end date of a range.

	*/
	EndDate *strfmt.Date
	/*Fields
	  Fields in response. Array that lists the fields requested. If empty or not supplied, all attributes will be returned.

	*/
	Fields []string
	/*MovieID
	  Movie identifier. Request the movie with movie_id.

	*/
	MovieID int64
	/*StartDate
	  Start date. Request the sessions of this movie playing on or after start date of a range.

	*/
	StartDate *strfmt.Date

	timeout time.Duration
}

// WithCentreID adds the centreId to the get centres centre ID movies movie ID params
func (o *GetCentresCentreIDMoviesMovieIDParams) WithCentreID(CentreID string) *GetCentresCentreIDMoviesMovieIDParams {
	o.CentreID = CentreID
	return o
}

// WithEndDate adds the endDate to the get centres centre ID movies movie ID params
func (o *GetCentresCentreIDMoviesMovieIDParams) WithEndDate(EndDate *strfmt.Date) *GetCentresCentreIDMoviesMovieIDParams {
	o.EndDate = EndDate
	return o
}

// WithFields adds the fields to the get centres centre ID movies movie ID params
func (o *GetCentresCentreIDMoviesMovieIDParams) WithFields(Fields []string) *GetCentresCentreIDMoviesMovieIDParams {
	o.Fields = Fields
	return o
}

// WithMovieID adds the movieId to the get centres centre ID movies movie ID params
func (o *GetCentresCentreIDMoviesMovieIDParams) WithMovieID(MovieID int64) *GetCentresCentreIDMoviesMovieIDParams {
	o.MovieID = MovieID
	return o
}

// WithStartDate adds the startDate to the get centres centre ID movies movie ID params
func (o *GetCentresCentreIDMoviesMovieIDParams) WithStartDate(StartDate *strfmt.Date) *GetCentresCentreIDMoviesMovieIDParams {
	o.StartDate = StartDate
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetCentresCentreIDMoviesMovieIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param centre_id
	if err := r.SetPathParam("centre_id", o.CentreID); err != nil {
		return err
	}

	if o.EndDate != nil {

		// query param end_date
		var qrEndDate strfmt.Date
		if o.EndDate != nil {
			qrEndDate = *o.EndDate
		}
		qEndDate := qrEndDate.String()
		if qEndDate != "" {
			if err := r.SetQueryParam("end_date", qEndDate); err != nil {
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

	// path param movie_id
	if err := r.SetPathParam("movie_id", swag.FormatInt64(o.MovieID)); err != nil {
		return err
	}

	if o.StartDate != nil {

		// query param start_date
		var qrStartDate strfmt.Date
		if o.StartDate != nil {
			qrStartDate = *o.StartDate
		}
		qStartDate := qrStartDate.String()
		if qStartDate != "" {
			if err := r.SetQueryParam("start_date", qStartDate); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}