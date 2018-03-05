// Code generated by go-swagger; DO NOT EDIT.

package upload

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

// NewFileUploadParams creates a new FileUploadParams object
// with the default values initialized.
func NewFileUploadParams() FileUploadParams {
	var ()
	return FileUploadParams{}
}

// FileUploadParams contains all the bound params for the file upload operation
// typically these are obtained from a http.Request
//
// swagger:parameters file/upload/
type FileUploadParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*文件后缀
	  In: formData
	*/
	FileExt *string
	/*上传文件类型
	  In: formData
	*/
	FileType *string
	/*The file to upload.
	  In: formData
	*/
	Filename io.ReadCloser
	/*时间戳
	  In: formData
	*/
	Ts *int64
	/*当前登录用户id
	  In: formData
	*/
	Userid *int64
	/*加密串
	  In: formData
	*/
	Val *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *FileUploadParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		if err != http.ErrNotMultipart {
			return err
		} else if err := r.ParseForm(); err != nil {
			return err
		}
	}
	fds := runtime.Values(r.Form)

	fdFileExt, fdhkFileExt, _ := fds.GetOK("fileExt")
	if err := o.bindFileExt(fdFileExt, fdhkFileExt, route.Formats); err != nil {
		res = append(res, err)
	}

	fdFileType, fdhkFileType, _ := fds.GetOK("fileType")
	if err := o.bindFileType(fdFileType, fdhkFileType, route.Formats); err != nil {
		res = append(res, err)
	}

	filename, filenameHeader, err := r.FormFile("filename")
	if err != nil && err != http.ErrMissingFile {
		res = append(res, errors.New(400, "reading file %q failed: %v", "filename", err))
	} else if err == http.ErrMissingFile {
		// no-op for missing but optional file parameter
	} else if err := o.bindFilename(filename, filenameHeader); err != nil {
		res = append(res, err)
	} else {
		o.Filename = &runtime.File{Data: filename, Header: filenameHeader}
	}

	fdTs, fdhkTs, _ := fds.GetOK("ts")
	if err := o.bindTs(fdTs, fdhkTs, route.Formats); err != nil {
		res = append(res, err)
	}

	fdUserid, fdhkUserid, _ := fds.GetOK("userid")
	if err := o.bindUserid(fdUserid, fdhkUserid, route.Formats); err != nil {
		res = append(res, err)
	}

	fdVal, fdhkVal, _ := fds.GetOK("val")
	if err := o.bindVal(fdVal, fdhkVal, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *FileUploadParams) bindFileExt(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.FileExt = &raw

	return nil
}

func (o *FileUploadParams) bindFileType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.FileType = &raw

	return nil
}

func (o *FileUploadParams) bindFilename(file multipart.File, header *multipart.FileHeader) error {

	return nil
}

func (o *FileUploadParams) bindTs(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("ts", "formData", "int64", raw)
	}
	o.Ts = &value

	return nil
}

func (o *FileUploadParams) bindUserid(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *FileUploadParams) bindVal(rawData []string, hasKey bool, formats strfmt.Registry) error {
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