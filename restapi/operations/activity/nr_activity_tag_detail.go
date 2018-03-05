// Code generated by go-swagger; DO NOT EDIT.

package activity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// NrActivityTagDetailHandlerFunc turns a function with the right signature into a activity tag detail handler
type NrActivityTagDetailHandlerFunc func(NrActivityTagDetailParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrActivityTagDetailHandlerFunc) Handle(params NrActivityTagDetailParams) middleware.Responder {
	return fn(params)
}

// NrActivityTagDetailHandler interface for that can handle valid activity tag detail params
type NrActivityTagDetailHandler interface {
	Handle(NrActivityTagDetailParams) middleware.Responder
}

// NewNrActivityTagDetail creates a new http.Handler for the activity tag detail operation
func NewNrActivityTagDetail(ctx *middleware.Context, handler NrActivityTagDetailHandler) *NrActivityTagDetail {
	return &NrActivityTagDetail{Context: ctx, Handler: handler}
}

/*NrActivityTagDetail swagger:route GET /activity/tag/detail Activity activityTagDetail

活动Tag详情

活动tag详情

*/
type NrActivityTagDetail struct {
	Context *middleware.Context
	Handler NrActivityTagDetailHandler
}

func (o *NrActivityTagDetail) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrActivityTagDetailParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
