// Code generated by go-swagger; DO NOT EDIT.

package relation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// NrSubCategoryGoodsListRelationHandlerFunc turns a function with the right signature into a sub category goods list relation handler
type NrSubCategoryGoodsListRelationHandlerFunc func(NrSubCategoryGoodsListRelationParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NrSubCategoryGoodsListRelationHandlerFunc) Handle(params NrSubCategoryGoodsListRelationParams) middleware.Responder {
	return fn(params)
}

// NrSubCategoryGoodsListRelationHandler interface for that can handle valid sub category goods list relation params
type NrSubCategoryGoodsListRelationHandler interface {
	Handle(NrSubCategoryGoodsListRelationParams) middleware.Responder
}

// NewNrSubCategoryGoodsListRelation creates a new http.Handler for the sub category goods list relation operation
func NewNrSubCategoryGoodsListRelation(ctx *middleware.Context, handler NrSubCategoryGoodsListRelationHandler) *NrSubCategoryGoodsListRelation {
	return &NrSubCategoryGoodsListRelation{Context: ctx, Handler: handler}
}

/*NrSubCategoryGoodsListRelation swagger:route GET /relation/subCategory/goodsList Relation subCategoryGoodsListRelation

获取子类下商品集合

获取子类下商品集合

*/
type NrSubCategoryGoodsListRelation struct {
	Context *middleware.Context
	Handler NrSubCategoryGoodsListRelationHandler
}

func (o *NrSubCategoryGoodsListRelation) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewNrSubCategoryGoodsListRelationParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}