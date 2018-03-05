// Code generated by go-swagger; DO NOT EDIT.

package goods

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GoodsListHandlerFunc turns a function with the right signature into a goods list handler
type GoodsListHandlerFunc func(GoodsListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GoodsListHandlerFunc) Handle(params GoodsListParams) middleware.Responder {
	return fn(params)
}

// GoodsListHandler interface for that can handle valid goods list params
type GoodsListHandler interface {
	Handle(GoodsListParams) middleware.Responder
}

// NewGoodsList creates a new http.Handler for the goods list operation
func NewGoodsList(ctx *middleware.Context, handler GoodsListHandler) *GoodsList {
	return &GoodsList{Context: ctx, Handler: handler}
}

/*GoodsList swagger:route GET /goods/list Goods goodsList

获取商品列表

获取商品列表

*/
type GoodsList struct {
	Context *middleware.Context
	Handler GoodsListHandler
}

func (o *GoodsList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGoodsListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
