// Code generated by go-swagger; DO NOT EDIT.

package member

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// NrMemberFansListHandlerFunc turns a function with the right signature into a member fans list handler
type NrMemberFansListHandlerFunc func(NrMemberFansListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrMemberFansListHandlerFunc) Handle(params NrMemberFansListParams) middleware.Responder {
	return fn(params)
}

// NrMemberFansListHandler interface for that can handle valid member fans list params
type NrMemberFansListHandler interface {
	Handle(NrMemberFansListParams) middleware.Responder
}

// NewNrMemberFansList creates a new http.Handler for the member fans list operation
func NewNrMemberFansList(ctx *middleware.Context, handler NrMemberFansListHandler) *NrMemberFansList {
	return &NrMemberFansList{Context: ctx, Handler: handler}
}

/*NrMemberFansList swagger:route GET /member/fans/list Member memberFansList

获取关注xx会员的人列表(含条件检索)

获取关注xx会员的人列表

*/
type NrMemberFansList struct {
	Context *middleware.Context
	Handler NrMemberFansListHandler
}

func (o *NrMemberFansList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrMemberFansListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
