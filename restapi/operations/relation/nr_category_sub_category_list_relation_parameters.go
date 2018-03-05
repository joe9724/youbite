// Code generated by go-swagger; DO NOT EDIT.

package relation

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

// NewNrCategorySubCategoryListRelationParams creates a new NrCategorySubCategoryListRelationParams object
// with the default values initialized.
func NewNrCategorySubCategoryListRelationParams() NrCategorySubCategoryListRelationParams {
	var ()
	return NrCategorySubCategoryListRelationParams{}
}

// NrCategorySubCategoryListRelationParams contains all the bound params for the category sub category list relation operation
// typically these are obtained from a http.Request
//
// swagger:parameters /Category/subCategoryList/relation
type NrCategorySubCategoryListRelationParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*会员id
	  In: query
	*/
	MemberID *int64
	/*分页索引
	  In: query
	*/
	PageIndex *int64
	/*分页尺寸
	  In: query
	*/
	PageSize *int64
	/*大类id
	  In: query
	*/
	SubCategoryID *string
	/*当前后台登陆用户id
	  In: query
	*/
	Userid *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *NrCategorySubCategoryListRelationParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qMemberID, qhkMemberID, _ := qs.GetOK("memberId")
	if err := o.bindMemberID(qMemberID, qhkMemberID, route.Formats); err != nil {
		res = append(res, err)
	}

	qPageIndex, qhkPageIndex, _ := qs.GetOK("pageIndex")
	if err := o.bindPageIndex(qPageIndex, qhkPageIndex, route.Formats); err != nil {
		res = append(res, err)
	}

	qPageSize, qhkPageSize, _ := qs.GetOK("pageSize")
	if err := o.bindPageSize(qPageSize, qhkPageSize, route.Formats); err != nil {
		res = append(res, err)
	}

	qSubCategoryID, qhkSubCategoryID, _ := qs.GetOK("subCategoryId")
	if err := o.bindSubCategoryID(qSubCategoryID, qhkSubCategoryID, route.Formats); err != nil {
		res = append(res, err)
	}

	qUserid, qhkUserid, _ := qs.GetOK("userid")
	if err := o.bindUserid(qUserid, qhkUserid, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NrCategorySubCategoryListRelationParams) bindMemberID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("memberId", "query", "int64", raw)
	}
	o.MemberID = &value

	return nil
}

func (o *NrCategorySubCategoryListRelationParams) bindPageIndex(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("pageIndex", "query", "int64", raw)
	}
	o.PageIndex = &value

	return nil
}

func (o *NrCategorySubCategoryListRelationParams) bindPageSize(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("pageSize", "query", "int64", raw)
	}
	o.PageSize = &value

	return nil
}

func (o *NrCategorySubCategoryListRelationParams) bindSubCategoryID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.SubCategoryID = &raw

	return nil
}

func (o *NrCategorySubCategoryListRelationParams) bindUserid(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Userid = &raw

	return nil
}
