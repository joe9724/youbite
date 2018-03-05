// Code generated by go-swagger; DO NOT EDIT.

package shuoshuo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// NrShuoshuoDetailURL generates an URL for the shuoshuo detail operation
type NrShuoshuoDetailURL struct {
	ShuoshuoID *int64
	Ts         *int64
	Userid     *int64
	Val        *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *NrShuoshuoDetailURL) WithBasePath(bp string) *NrShuoshuoDetailURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *NrShuoshuoDetailURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *NrShuoshuoDetailURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/shuoshuo/detail"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/nanjingyouzi/YoubiteApi/1.0.0"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var shuoshuoID string
	if o.ShuoshuoID != nil {
		shuoshuoID = swag.FormatInt64(*o.ShuoshuoID)
	}
	if shuoshuoID != "" {
		qs.Set("shuoshuoId", shuoshuoID)
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
		userid = swag.FormatInt64(*o.Userid)
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
func (o *NrShuoshuoDetailURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *NrShuoshuoDetailURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *NrShuoshuoDetailURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on NrShuoshuoDetailURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on NrShuoshuoDetailURL")
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
func (o *NrShuoshuoDetailURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}