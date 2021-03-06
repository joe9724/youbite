// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Order order
// swagger:model Order
type Order struct {

	// goods Id
	GoodsID int64 `json:"goodsId,omitempty"`

	// goods name
	GoodsName string `json:"goodsName,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// member Id
	MemberID int64 `json:"memberId,omitempty"`

	// member name
	MemberName string `json:"memberName,omitempty"`

	// order no
	OrderNo string `json:"orderNo,omitempty"`

	// status
	Status int64 `json:"status,omitempty"`

	// time
	Time int64 `json:"time,omitempty"`

	// value
	Value float64 `json:"value,omitempty"`
}

// Validate validates this order
func (m *Order) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Order) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Order) UnmarshalBinary(b []byte) error {
	var res Order
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
