// Code generated by go-swagger; DO NOT EDIT.

package member

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
)

// NrMemberLoginByThirdPartyURL generates an URL for the member login by third party operation
type NrMemberLoginByThirdPartyURL struct {
	_basePath string
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *NrMemberLoginByThirdPartyURL) WithBasePath(bp string) *NrMemberLoginByThirdPartyURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *NrMemberLoginByThirdPartyURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *NrMemberLoginByThirdPartyURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/member/loginByThirdParty"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/nanjingyouzi/YoubiteApi/1.0.0"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *NrMemberLoginByThirdPartyURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *NrMemberLoginByThirdPartyURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *NrMemberLoginByThirdPartyURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on NrMemberLoginByThirdPartyURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on NrMemberLoginByThirdPartyURL")
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
func (o *NrMemberLoginByThirdPartyURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
