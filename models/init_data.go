// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InitData init data
// swagger:model InitData
type InitData struct {

	// download Url
	DownloadURL string `json:"downloadUrl,omitempty"`

	// extra info
	ExtraInfo *InitDataExtraInfo `json:"extraInfo,omitempty"`

	// force
	Force int32 `json:"force,omitempty"`

	// msg
	Msg string `json:"msg,omitempty"`

	// number
	Number int64 `json:"number,omitempty"`
}

// Validate validates this init data
func (m *InitData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExtraInfo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InitData) validateExtraInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.ExtraInfo) { // not required
		return nil
	}

	if m.ExtraInfo != nil {

		if err := m.ExtraInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("extraInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InitData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InitData) UnmarshalBinary(b []byte) error {
	var res InitData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
