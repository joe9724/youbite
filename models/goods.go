// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Goods goods
// swagger:model Goods
type Goods struct {

	// author avatar
	AuthorAvatar string `json:"authorAvatar,omitempty"`

	// author name
	AuthorName string `json:"authorName,omitempty"`

	// books number
	BooksNumber int64 `json:"booksNumber,omitempty"`

	// icon
	Icon string `json:"icon,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// play count
	PlayCount int64 `json:"playCount,omitempty"`

	// show icon
	ShowIcon bool `json:"showIcon,omitempty"`

	// status
	Status int64 `json:"status,omitempty"`

	// sub category Id
	SubCategoryID int64 `json:"subCategoryId,omitempty"`

	// sub title
	SubTitle string `json:"subTitle,omitempty"`

	// summary
	Summary string `json:"summary,omitempty"`

	// time
	Time int64 `json:"time,omitempty"`

	// value
	Value float64 `json:"value,omitempty"`
}

// Validate validates this goods
func (m *Goods) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Goods) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Goods) UnmarshalBinary(b []byte) error {
	var res Goods
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
