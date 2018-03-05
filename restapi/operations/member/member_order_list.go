// Code generated by go-swagger; DO NOT EDIT.

package member

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// MemberOrderListHandlerFunc turns a function with the right signature into a member order list handler
type MemberOrderListHandlerFunc func(MemberOrderListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn MemberOrderListHandlerFunc) Handle(params MemberOrderListParams) middleware.Responder {
	return fn(params)
}

// MemberOrderListHandler interface for that can handle valid member order list params
type MemberOrderListHandler interface {
	Handle(MemberOrderListParams) middleware.Responder
}

// NewMemberOrderList creates a new http.Handler for the member order list operation
func NewMemberOrderList(ctx *middleware.Context, handler MemberOrderListHandler) *MemberOrderList {
	return &MemberOrderList{Context: ctx, Handler: handler}
}

/*MemberOrderList swagger:route GET /member/order/list Member memberOrderList

获取订单列表(含条件检索)

获取订单列表

*/
type MemberOrderList struct {
	Context *middleware.Context
	Handler MemberOrderListHandler
}

func (o *MemberOrderList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewMemberOrderListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
