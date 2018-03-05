// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Topic 话题
// swagger:model Topic
type Topic struct {

	// content
	Content string `json:"content,omitempty"`

	// cover
	Cover string `json:"cover,omitempty"`

	// icon
	Icon string `json:"icon,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// lat
	Lat float64 `json:"lat,omitempty"`

	// lon
	Lon float64 `json:"lon,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// sub title
	SubTitle string `json:"subTitle,omitempty"`

	// summary
	Summary string `json:"summary,omitempty"`

	// time
	Time int64 `json:"time,omitempty"`
}

// Validate validates this topic
func (m *Topic) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Topic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Topic) UnmarshalBinary(b []byte) error {
	var res Topic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}