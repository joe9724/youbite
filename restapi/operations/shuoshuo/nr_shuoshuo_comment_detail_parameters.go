// Code generated by go-swagger; DO NOT EDIT.

package shuoshuo

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

// NewNrShuoshuoCommentDetailParams creates a new NrShuoshuoCommentDetailParams object
// with the default values initialized.
func NewNrShuoshuoCommentDetailParams() NrShuoshuoCommentDetailParams {
	var ()
	return NrShuoshuoCommentDetailParams{}
}

// NrShuoshuoCommentDetailParams contains all the bound params for the shuoshuo comment detail operation
// typically these are obtained from a http.Request
//
// swagger:parameters /shuoshuo/comment/detail
type NrShuoshuoCommentDetailParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*评论id
	  In: query
	*/
	CommentID *int64
	/*
	  In: query
	*/
	Ts *int64
	/*当前后台登录id
	  In: query
	*/
	Userid *int64
	/*
	  In: query
	*/
	Val *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *NrShuoshuoCommentDetailParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qCommentID, qhkCommentID, _ := qs.GetOK("commentId")
	if err := o.bindCommentID(qCommentID, qhkCommentID, route.Formats); err != nil {
		res = append(res, err)
	}

	qTs, qhkTs, _ := qs.GetOK("ts")
	if err := o.bindTs(qTs, qhkTs, route.Formats); err != nil {
		res = append(res, err)
	}

	qUserid, qhkUserid, _ := qs.GetOK("userid")
	if err := o.bindUserid(qUserid, qhkUserid, route.Formats); err != nil {
		res = append(res, err)
	}

	qVal, qhkVal, _ := qs.GetOK("val")
	if err := o.bindVal(qVal, qhkVal, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NrShuoshuoCommentDetailParams) bindCommentID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("commentId", "query", "int64", raw)
	}
	o.CommentID = &value

	return nil
}

func (o *NrShuoshuoCommentDetailParams) bindTs(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *NrShuoshuoCommentDetailParams) bindUserid(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("userid", "query", "int64", raw)
	}
	o.Userid = &value

	return nil
}

func (o *NrShuoshuoCommentDetailParams) bindVal(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
