// Code generated by go-swagger; DO NOT EDIT.

package activity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// NrActivityTagListURL generates an URL for the activity tag list operation
type NrActivityTagListURL struct {
	EndTime   *int64
	Lat       *float64
	Lon       *float64
	PageIndex *int64
	PageSize  *int64
	StartTime *int64
	Tag       *string
	TagID     *int64
	Ts        *int64
	Userid    *string
	Val       *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *NrActivityTagListURL) WithBasePath(bp string) *NrActivityTagListURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *NrActivityTagListURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *NrActivityTagListURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/activity/tag/list"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/nanjingyouzi/YoubiteApi/1.0.0"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var endTime string
	if o.EndTime != nil {
		endTime = swag.FormatInt64(*o.EndTime)
	}
	if endTime != "" {
		qs.Set("endTime", endTime)
	}

	var lat string
	if o.Lat != nil {
		lat = swag.FormatFloat64(*o.Lat)
	}
	if lat != "" {
		qs.Set("lat", lat)
	}

	var lon string
	if o.Lon != nil {
		lon = swag.FormatFloat64(*o.Lon)
	}
	if lon != "" {
		qs.Set("lon", lon)
	}

	var pageIndex string
	if o.PageIndex != nil {
		pageIndex = swag.FormatInt64(*o.PageIndex)
	}
	if pageIndex != "" {
		qs.Set("pageIndex", pageIndex)
	}

	var pageSize string
	if o.PageSize != nil {
		pageSize = swag.FormatInt64(*o.PageSize)
	}
	if pageSize != "" {
		qs.Set("pageSize", pageSize)
	}

	var startTime string
	if o.StartTime != nil {
		startTime = swag.FormatInt64(*o.StartTime)
	}
	if startTime != "" {
		qs.Set("startTime", startTime)
	}

	var tag string
	if o.Tag != nil {
		tag = *o.Tag
	}
	if tag != "" {
		qs.Set("tag", tag)
	}

	var tagID string
	if o.TagID != nil {
		tagID = swag.FormatInt64(*o.TagID)
	}
	if tagID != "" {
		qs.Set("tagId", tagID)
	}

	var ts string
	if o.Ts != nil {
		ts = swag.FormatInt64(*o.Ts)
	}
	if ts != "" {
		qs.Set("ts", ts)
	}

	var userid string
	if o.Userid != nil {
		userid = *o.Userid
	}
	if userid != "" {
		qs.Set("userid", userid)
	}

	var val string
	if o.Val != nil {
		val = *o.Val
	}
	if val != "" {
		qs.Set("val", val)
	}

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *NrActivityTagListURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *NrActivityTagListURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *NrActivityTagListURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on NrActivityTagListURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on NrActivityTagListURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *NrActivityTagListURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
