// Code generated by go-swagger; DO NOT EDIT.

package member

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewNrStartUpParams creates a new NrStartUpParams object
// with the default values initialized.
func NewNrStartUpParams() NrStartUpParams {
	var ()
	return NrStartUpParams{}
}

// NrStartUpParams contains all the bound params for the start up operation
// typically these are obtained from a http.Request
//
// swagger:parameters /startUp
type NrStartUpParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*客户端类型
	  In: query
	*/
	Client *string
	/*唯一识别号
	  In: query
	*/
	Imei *string
	/*
	  In: query
	*/
	Ts *int64
	/*
	  In: query
	*/
	Val *string
	/*版本号
	  In: query
	*/
	Version *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *NrStartUpParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qClient, qhkClient, _ := qs.GetOK("client")
	if err := o.bindClient(qClient, qhkClient, route.Formats); err != nil {
		res = append(res, err)
	}

	qImei, qhkImei, _ := qs.GetOK("imei")
	if err := o.bindImei(qImei, qhkImei, route.Formats); err != nil {
		res = append(res, err)
	}

	qTs, qhkTs, _ := qs.GetOK("ts")
	if err := o.bindTs(qTs, qhkTs, route.Formats); err != nil {
		res = append(res, err)
	}

	qVal, qhkVal, _ := qs.GetOK("val")
	if err := o.bindVal(qVal, qhkVal, route.Formats); err != nil {
		res = append(res, err)
	}

	qVersion, qhkVersion, _ := qs.GetOK("version")
	if err := o.bindVersion(qVersion, qhkVersion, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NrStartUpParams) bindClient(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Client = &raw

	return nil
}

func (o *NrStartUpParams) bindImei(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Imei = &raw

	return nil
}

func (o *NrStartUpParams) bindTs(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("ts", "query", "int64", raw)
	}
	o.Ts = &value

	return nil
}

func (o *NrStartUpParams) bindVal(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Val = &raw

	return nil
}

func (o *NrStartUpParams) bindVersion(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Version = &raw

	return nil
}
