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

// NrMemberAttentionListURL generates an URL for the member attention list operation
type NrMemberAttentionListURL struct {
	EndTime    *int64
	MemberID   *int64
	MemberName *string
	PageIndex  *int64
	PageSize   *int64
	StartTime  *int64
	Ts         *int64
	Userid     *string
	Val        *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *NrMemberAttentionListURL) WithBasePath(bp string) *NrMemberAttentionListURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *NrMemberAttentionListURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *NrMemberAttentionListURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/member/attention/list"

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

	var memberID string
	if o.MemberID != nil {
		memberID = swag.FormatInt64(*o.MemberID)
	}
	if memberID != "" {
		qs.Set("memberId", memberID)
	}

	var memberName string
	if o.MemberName != nil {
		memberName = *o.MemberName
	}
	if memberName != "" {
		qs.Set("memberName", memberName)
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
func (o *NrMemberAttentionListURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *NrMemberAttentionListURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *NrMemberAttentionListURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on NrMemberAttentionListURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on NrMemberAttentionListURL")
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
func (o *NrMemberAttentionListURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
