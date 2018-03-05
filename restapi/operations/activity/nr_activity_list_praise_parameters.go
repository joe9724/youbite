// Code generated by go-swagger; DO NOT EDIT.

package activity

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

// NewNrActivityListPraiseParams creates a new NrActivityListPraiseParams object
// with the default values initialized.
func NewNrActivityListPraiseParams() NrActivityListPraiseParams {
	var ()
	return NrActivityListPraiseParams{}
}

// NrActivityListPraiseParams contains all the bound params for the activity list praise operation
// typically these are obtained from a http.Request
//
// swagger:parameters /activity/list/praise
type NrActivityListPraiseParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*活动id
	  In: query
	*/
	ActivityID *int64
	/*结束时间
	  In: query
	*/
	EndTime *int64
	/*说说模糊匹配(可模糊检索)
	  In: query
	*/
	Keyword *string
	/*可用于搜索
	  In: query
	*/
	Lat *float64
	/*可用于搜索
	  In: query
	*/
	Lon *float64
	/*当前用户id
	  In: query
	*/
	MemberID *string
	/*分页索引
	  In: query
	*/
	PageIndex *int64
	/*分页尺寸
	  In: query
	*/
	PageSize *int64
	/*开始时间
	  In: query
	*/
	StartTime *int64
	/*
	  In: query
	*/
	Ts *int64
	/*
	  In: query
	*/
	Val *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *NrActivityListPraiseParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qActivityID, qhkActivityID, _ := qs.GetOK("activityId")
	if err := o.bindActivityID(qActivityID, qhkActivityID, route.Formats); err != nil {
		res = append(res, err)
	}

	qEndTime, qhkEndTime, _ := qs.GetOK("endTime")
	if err := o.bindEndTime(qEndTime, qhkEndTime, route.Formats); err != nil {
		res = append(res, err)
	}

	qKeyword, qhkKeyword, _ := qs.GetOK("keyword")
	if err := o.bindKeyword(qKeyword, qhkKeyword, route.Formats); err != nil {
		res = append(res, err)
	}

	qLat, qhkLat, _ := qs.GetOK("lat")
	if err := o.bindLat(qLat, qhkLat, route.Formats); err != nil {
		res = append(res, err)
	}

	qLon, qhkLon, _ := qs.GetOK("lon")
	if err := o.bindLon(qLon, qhkLon, route.Formats); err != nil {
		res = append(res, err)
	}

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

	qStartTime, qhkStartTime, _ := qs.GetOK("startTime")
	if err := o.bindStartTime(qStartTime, qhkStartTime, route.Formats); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NrActivityListPraiseParams) bindActivityID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("activityId", "query", "int64", raw)
	}
	o.ActivityID = &value

	return nil
}

func (o *NrActivityListPraiseParams) bindEndTime(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("endTime", "query", "int64", raw)
	}
	o.EndTime = &value

	return nil
}

func (o *NrActivityListPraiseParams) bindKeyword(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Keyword = &raw

	return nil
}

func (o *NrActivityListPraiseParams) bindLat(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertFloat64(raw)
	if err != nil {
		return errors.InvalidType("lat", "query", "float64", raw)
	}
	o.Lat = &value

	return nil
}

func (o *NrActivityListPraiseParams) bindLon(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertFloat64(raw)
	if err != nil {
		return errors.InvalidType("lon", "query", "float64", raw)
	}
	o.Lon = &value

	return nil
}

func (o *NrActivityListPraiseParams) bindMemberID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.MemberID = &raw

	return nil
}

func (o *NrActivityListPraiseParams) bindPageIndex(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *NrActivityListPraiseParams) bindPageSize(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *NrActivityListPraiseParams) bindStartTime(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("startTime", "query", "int64", raw)
	}
	o.StartTime = &value

	return nil
}

func (o *NrActivityListPraiseParams) bindTs(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

func (o *NrActivityListPraiseParams) bindVal(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
