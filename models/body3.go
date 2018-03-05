// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Body3 body 3
// swagger:model body_3
type Body3 struct {

	// client
	Client string `json:"client,omitempty"`

	// content
	Content string `json:"content,omitempty"`

	// imei
	Imei string `json:"imei,omitempty"`

	// ts
	Ts int64 `json:"ts,omitempty"`

	// val
	Val string `json:"val,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this body 3
func (m *Body3) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Body3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Body3) UnmarshalBinary(b []byte) error {
	var res Body3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
