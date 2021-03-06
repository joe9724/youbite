// Code generated by go-swagger; DO NOT EDIT.

package activity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"mime/multipart"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewActivityTagUploadParams creates a new ActivityTagUploadParams object
// with the default values initialized.
func NewActivityTagUploadParams() ActivityTagUploadParams {
	var ()
	return ActivityTagUploadParams{}
}

// ActivityTagUploadParams contains all the bound params for the activity tag upload operation
// typically these are obtained from a http.Request
//
// swagger:parameters activity/tag/upload
type ActivityTagUploadParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: formData
	*/
	Content *string
	/*The file to upload.
	  In: formData
	*/
	Cover io.ReadCloser
	/*The file to upload.
	  In: formData
	*/
	Icon io.ReadCloser
	/*
	  In: formData
	*/
	SubTitle *string
	/*
	  In: formData
	*/
	Summary *string
	/*Description of file contents.
	  In: formData
	*/
	Title *string
	/*
	  In: query
	*/
	Ts *int64
	/*当前登录用户id
	  In: formData
	*/
	Userid *int64
	/*
	  In: query
	*/
	Val *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *ActivityTagUploadParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		if err != http.ErrNotMultipart {
			return err
		} else if err := r.ParseForm(); err != nil {
			return err
		}
	}
	fds := runtime.Values(r.Form)

	fdContent, fdhkContent, _ := fds.GetOK("content")
	if err := o.bindContent(fdContent, fdhkContent, route.Formats); err != nil {
		res = append(res, err)
	}

	cover, coverHeader, err := r.FormFile("cover")
	if err != nil && err != http.ErrMissingFile {
		res = append(res, errors.New(400, "reading file %q failed: %v", "cover", err))
	} else if err == http.ErrMissingFile {
		// no-op for missing but optional file parameter
	} else if err := o.bindCover(cover, coverHeader); err != nil {
		res = append(res, err)
	} else {
		o.Cover = &runtime.File{Data: cover, Header: coverHeader}
	}

	icon, iconHeader, err := r.FormFile("icon")
	if err != nil && err != http.ErrMissingFile {
		res = append(res, errors.New(400, "reading file %q failed: %v", "icon", err))
	} else if err == http.ErrMissingFile {
		// no-op for missing but optional file parameter
	} else if err := o.bindIcon(icon, iconHeader); err != nil {
		res = append(res, err)
	} else {
		o.Icon = &runtime.File{Data: icon, Header: iconHeader}
	}

	fdSubTitle, fdhkSubTitle, _ := fds.GetOK("subTitle")
	if err := o.bindSubTitle(fdSubTitle, fdhkSubTitle, route.Formats); err != nil {
		res = append(res, err)
	}

	fdSummary, fdhkSummary, _ := fds.GetOK("summary")
	if err := o.bindSummary(fdSummary, fdhkSummary, route.Formats); err != nil {
		res = append(res, err)
	}

	fdTitle, fdhkTitle, _ := fds.GetOK("title")
	if err := o.bindTitle(fdTitle, fdhkTitle, route.Formats); err != nil {
		res = append(res, err)
	}

	qTs, qhkTs, _ := qs.GetOK("ts")
	if err := o.bindTs(qTs, qhkTs, route.Formats); err != nil {
		res = append(res, err)
	}

	fdUserid, fdhkUserid, _ := fds.GetOK("userid")
	if err := o.bindUserid(fdUserid, fdhkUserid, route.Formats); err != nil {
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

func (o *ActivityTagUploadParams) bindContent(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Content = &raw

	return nil
}

func (o *ActivityTagUploadParams) bindCover(file multipart.File, header *multipart.FileHeader) error {

	return nil
}

func (o *ActivityTagUploadParams) bindIcon(file multipart.File, header *multipart.FileHeader) error {

	return nil
}

func (o *ActivityTagUploadParams) bindSubTitle(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.SubTitle = &raw

	return nil
}

func (o *ActivityTagUploadParams) bindSummary(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Summary = &raw

	return nil
}

func (o *ActivityTagUploadParams) bindTitle(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Title = &raw

	return nil
}

func (o *ActivityTagUploadParams) bindTs(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *ActivityTagUploadParams) bindUserid(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("userid", "formData", "int64", raw)
	}
	o.Userid = &value

	return nil
}

func (o *ActivityTagUploadParams) bindVal(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
