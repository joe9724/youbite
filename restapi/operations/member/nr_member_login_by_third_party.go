// Code generated by go-swagger; DO NOT EDIT.

package member

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// NrMemberLoginByThirdPartyHandlerFunc turns a function with the right signature into a member login by third party handler
type NrMemberLoginByThirdPartyHandlerFunc func(NrMemberLoginByThirdPartyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrMemberLoginByThirdPartyHandlerFunc) Handle(params NrMemberLoginByThirdPartyParams) middleware.Responder {
	return fn(params)
}

// NrMemberLoginByThirdPartyHandler interface for that can handle valid member login by third party params
type NrMemberLoginByThirdPartyHandler interface {
	Handle(NrMemberLoginByThirdPartyParams) middleware.Responder
}

// NewNrMemberLoginByThirdParty creates a new http.Handler for the member login by third party operation
func NewNrMemberLoginByThirdParty(ctx *middleware.Context, handler NrMemberLoginByThirdPartyHandler) *NrMemberLoginByThirdParty {
	return &NrMemberLoginByThirdParty{Context: ctx, Handler: handler}
}

/*NrMemberLoginByThirdParty swagger:route POST /member/loginByThirdParty Member memberLoginByThirdParty

第三方登录

第三方登录

*/
type NrMemberLoginByThirdParty struct {
	Context *middleware.Context
	Handler NrMemberLoginByThirdPartyHandler
}

func (o *NrMemberLoginByThirdParty) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrMemberLoginByThirdPartyParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
