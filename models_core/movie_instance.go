package models_core

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*MovieInstance movie instance

swagger:model movieInstance
*/
type MovieInstance struct {

	/* links

	Required: true
	*/
	Links *MovieLinks `json:"_links"`

	/* classification

	Required: true
	*/
	Classification *MovieClassification `json:"classification"`

	/* Organisation which set the classification

	Required: true
	*/
	ClassificationBody *string `json:"classification_body"`

	/* Information about the classification

	Required: true
	*/
	ClassificationDetail *string `json:"classification_detail"`

	/* Date and time the movie was marked as deleted
	 */
	DeletedAt strfmt.DateTime `json:"deleted_at,omitempty"`

	/* directors

	Required: true
	*/
	Directors []string `json:"directors"`

	/* Movie genres

	Required: true
	*/
	Genres []string `json:"genres"`

	/* Movie ID

	Required: true
	*/
	MovieID *int64 `json:"movie_id"`

	/* Runtime of movie (in minutes)

	Required: true
	*/
	RunTime *int64 `json:"run_time"`

	/* sessions

	Required: true
	*/
	Sessions []*MovieSession `json:"sessions"`

	/* The text for the source attribution

	Required: true
	*/
	SourceAttributionText *string `json:"source_attribution_text"`

	/* The URI for the source attribution

	Required: true
	*/
	SourceAttributionURI *string `json:"source_attribution_uri"`

	/* Movie synopsis

	Required: true
	*/
	Synopsis *string `json:"synopsis"`

	/* The name of the movie session's time zone (e.g. 'US/Pacific')

	Required: true
	*/
	TimeZone *string `json:"time_zone"`

	/* Movie title

	Required: true
	*/
	Title *string `json:"title"`

	/* Top cast members

	Required: true
	*/
	TopCast []string `json:"top_cast"`

	/* Date and time the movie information was last updated

	Required: true
	*/
	UpdatedAt *strfmt.DateTime `json:"updated_at"`
}

// Validate validates this movie instance
func (m *MovieInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateClassification(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateClassificationBody(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateClassificationDetail(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDirectors(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGenres(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMovieID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRunTime(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSessions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSourceAttributionText(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSourceAttributionURI(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSynopsis(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTimeZone(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTopCast(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MovieInstance) validateLinks(formats strfmt.Registry) error {

	if err := validate.Required("_links", "body", m.Links); err != nil {
		return err
	}

	if m.Links != nil {

		if err := m.Links.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *MovieInstance) validateClassification(formats strfmt.Registry) error {

	if err := validate.Required("classification", "body", m.Classification); err != nil {
		return err
	}

	if m.Classification != nil {

		if err := m.Classification.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *MovieInstance) validateClassificationBody(formats strfmt.Registry) error {

	if err := validate.Required("classification_body", "body", m.ClassificationBody); err != nil {
		return err
	}

	return nil
}

func (m *MovieInstance) validateClassificationDetail(formats strfmt.Registry) error {

	if err := validate.Required("classification_detail", "body", m.ClassificationDetail); err != nil {
		return err
	}

	return nil
}

func (m *MovieInstance) validateDirectors(formats strfmt.Registry) error {

	if err := validate.Required("directors", "body", m.Directors); err != nil {
		return err
	}

	return nil
}

func (m *MovieInstance) validateGenres(formats strfmt.Registry) error {

	if err := validate.Required("genres", "body", m.Genres); err != nil {
		return err
	}

	return nil
}

func (m *MovieInstance) validateMovieID(formats strfmt.Registry) error {

	if err := validate.Required("movie_id", "body", m.MovieID); err != nil {
		return err
	}

	return nil
}

func (m *MovieInstance) validateRunTime(formats strfmt.Registry) error {

	if err := validate.Required("run_time", "body", m.RunTime); err != nil {
		return err
	}

	return nil
}

func (m *MovieInstance) validateSessions(formats strfmt.Registry) error {

	if err := validate.Required("sessions", "body", m.Sessions); err != nil {
		return err
	}

	for i := 0; i < len(m.Sessions); i++ {

		if swag.IsZero(m.Sessions[i]) { // not required
			continue
		}

		if m.Sessions[i] != nil {

			if err := m.Sessions[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *MovieInstance) validateSourceAttributionText(formats strfmt.Registry) error {

	if err := validate.Required("source_attribution_text", "body", m.SourceAttributionText); err != nil {
		return err
	}

	return nil
}

func (m *MovieInstance) validateSourceAttributionURI(formats strfmt.Registry) error {

	if err := validate.Required("source_attribution_uri", "body", m.SourceAttributionURI); err != nil {
		return err
	}

	return nil
}

func (m *MovieInstance) validateSynopsis(formats strfmt.Registry) error {

	if err := validate.Required("synopsis", "body", m.Synopsis); err != nil {
		return err
	}

	return nil
}

func (m *MovieInstance) validateTimeZone(formats strfmt.Registry) error {

	if err := validate.Required("time_zone", "body", m.TimeZone); err != nil {
		return err
	}

	return nil
}

func (m *MovieInstance) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *MovieInstance) validateTopCast(formats strfmt.Registry) error {

	if err := validate.Required("top_cast", "body", m.TopCast); err != nil {
		return err
	}

	return nil
}

func (m *MovieInstance) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", m.UpdatedAt); err != nil {
		return err
	}

	return nil
}
