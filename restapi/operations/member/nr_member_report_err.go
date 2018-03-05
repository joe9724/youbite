// Code generated by go-swagger; DO NOT EDIT.

package member

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// NrMemberReportErrHandlerFunc turns a function with the right signature into a member report err handler
type NrMemberReportErrHandlerFunc func(NrMemberReportErrParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrMemberReportErrHandlerFunc) Handle(params NrMemberReportErrParams) middleware.Responder {
	return fn(params)
}

// NrMemberReportErrHandler interface for that can handle valid member report err params
type NrMemberReportErrHandler interface {
	Handle(NrMemberReportErrParams) middleware.Responder
}

// NewNrMemberReportErr creates a new http.Handler for the member report err operation
func NewNrMemberReportErr(ctx *middleware.Context, handler NrMemberReportErrHandler) *NrMemberReportErr {
	return &NrMemberReportErr{Context: ctx, Handler: handler}
}

/*NrMemberReportErr swagger:route POST /member/reportErr Member memberReportErr

页面报错(备用)

页面报错

*/
type NrMemberReportErr struct {
	Context *middleware.Context
	Handler NrMemberReportErrHandler
}

func (o *NrMemberReportErr) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrMemberReportErrParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}