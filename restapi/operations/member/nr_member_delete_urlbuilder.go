// Code generated by go-swagger; DO NOT EDIT.

package member

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// NrMemberDeleteURL generates an URL for the member delete operation
type NrMemberDeleteURL struct {
	MemberID *int64
	Ts       *int64
	Userid   *int64
	Val      *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *NrMemberDeleteURL) WithBasePath(bp string) *NrMemberDeleteURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *NrMemberDeleteURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *NrMemberDeleteURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/member/delete"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/nanjingyouzi/YoubiteApi/1.0.0"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var memberID string
	if o.MemberID != nil {
		memberID = swag.FormatInt64(*o.MemberID)
	}
	if memberID != "" {
		qs.Set("memberId", memberID)
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
func (o *NrMemberDeleteURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *NrMemberDeleteURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *NrMemberDeleteURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on NrMemberDeleteURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on NrMemberDeleteURL")
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
func (o *NrMemberDeleteURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
