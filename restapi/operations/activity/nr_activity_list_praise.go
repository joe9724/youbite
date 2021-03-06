// Code generated by go-swagger; DO NOT EDIT.

package activity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// NrActivityListPraiseHandlerFunc turns a function with the right signature into a activity list praise handler
type NrActivityListPraiseHandlerFunc func(NrActivityListPraiseParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrActivityListPraiseHandlerFunc) Handle(params NrActivityListPraiseParams) middleware.Responder {
	return fn(params)
}

// NrActivityListPraiseHandler interface for that can handle valid activity list praise params
type NrActivityListPraiseHandler interface {
	Handle(NrActivityListPraiseParams) middleware.Responder
}

// NewNrActivityListPraise creates a new http.Handler for the activity list praise operation
func NewNrActivityListPraise(ctx *middleware.Context, handler NrActivityListPraiseHandler) *NrActivityListPraise {
	return &NrActivityListPraise{Context: ctx, Handler: handler}
}

/*NrActivityListPraise swagger:route GET /activity/list/praise Activity activityListPraise

获取活动的点赞列表(含条件检索)

获取活动的点赞列表(含条件检索)

*/
type NrActivityListPraise struct {
	Context *middleware.Context
	Handler NrActivityListPraiseHandler
}

func (o *NrActivityListPraise) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrActivityListPraiseParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
