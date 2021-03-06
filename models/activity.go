// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Activity 活动
// swagger:model Activity
type Activity struct {

	// 地区位置
	Area string `json:"area"`

	// content
	Content string `json:"content"`

	// cover
	Cover string `json:"cover"`

	// endt time
	EndtTime int64 `json:"endtTime"`

	// icon
	Icon string `json:"icon"`

	// id
	ID int64 `json:"id"`

	// lat
	Lat float64 `json:"lat"`

	// 参加人数上限 0=不限制
	Limit int64 `json:"limit"`

	// lon
	Lon float64 `json:"lon"`

	// name
	Name string `json:"name"`

	// price
	Price int64 `json:"price"`

	// price unit
	PriceUnit int64 `json:"priceUnit"`

	// start time
	StartTime int64 `json:"startTime"`

	// sub title
	SubTitle string `json:"subTitle"`

	// summary
	Summary string `json:"summary"`

	// tags
	Tags string `json:"tags"`

	// type
	Type int64 `json:"type"`

	// 描述活动的网页
	WebURL string `json:"webUrl"`

	ViewCount int64 `json:"viewCount"`

	Status int64 `json:"status"`

	JoinedCount int64 `json:"joinedCount"`

	FavCount int64 `json:"favCount"`

	PraiseCount int64 `json:"praiseCount"`
}

// Validate validates this activity
func (m *Activity) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Activity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Activity) UnmarshalBinary(b []byte) error {
	var res Activity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
