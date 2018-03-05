// Code generated by go-swagger; DO NOT EDIT.

package topic

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// NrTopicDetailHandlerFunc turns a function with the right signature into a topic detail handler
type NrTopicDetailHandlerFunc func(NrTopicDetailParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrTopicDetailHandlerFunc) Handle(params NrTopicDetailParams) middleware.Responder {
	return fn(params)
}

// NrTopicDetailHandler interface for that can handle valid topic detail params
type NrTopicDetailHandler interface {
	Handle(NrTopicDetailParams) middleware.Responder
}

// NewNrTopicDetail creates a new http.Handler for the topic detail operation
func NewNrTopicDetail(ctx *middleware.Context, handler NrTopicDetailHandler) *NrTopicDetail {
	return &NrTopicDetail{Context: ctx, Handler: handler}
}

/*NrTopicDetail swagger:route GET /topic/detail Topic topicDetail

话题详情

话题详情

*/
type NrTopicDetail struct {
	Context *middleware.Context
	Handler NrTopicDetailHandler
}

func (o *NrTopicDetail) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrTopicDetailParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}