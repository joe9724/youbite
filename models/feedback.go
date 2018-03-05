// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Feedback feedback
// swagger:model Feedback
type Feedback struct {

	// id
	ID int64 `json:"id,omitempty"`

	// status
	Status int64 `json:"status,omitempty"`

	// time
	Time int64 `json:"time,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// type
	Type int64 `json:"type,omitempty"`
}

// Validate validates this feedback
func (m *Feedback) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Feedback) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Feedback) UnmarshalBinary(b []byte) error {
	var res Feedback
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
