// Code generated by go-swagger; DO NOT EDIT.

package relation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "youbite/models"
)

// SubCategoryGoodsListRelationOKCode is the HTTP code returned for type SubCategoryGoodsListRelationOK
const SubCategoryGoodsListRelationOKCode int = 200

/*SubCategoryGoodsListRelationOK 登录成功，返回登录信息

swagger:response subCategoryGoodsListRelationOK
*/
type SubCategoryGoodsListRelationOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse20034 `json:"body,omitempty"`
}

// NewSubCategoryGoodsListRelationOK creates SubCategoryGoodsListRelationOK with default headers values
func NewSubCategoryGoodsListRelationOK() *SubCategoryGoodsListRelationOK {
	return &SubCategoryGoodsListRelationOK{}
}

// WithPayload adds the payload to the sub category goods list relation o k response
func (o *SubCategoryGoodsListRelationOK) WithPayload(payload *models.InlineResponse20034) *SubCategoryGoodsListRelationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the sub category goods list relation o k response
func (o *SubCategoryGoodsListRelationOK) SetPayload(payload *models.InlineResponse20034) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SubCategoryGoodsListRelationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NrSubCategoryGoodsListRelationDefault error

swagger:response subCategoryGoodsListRelationDefault
*/
type NrSubCategoryGoodsListRelationDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNrSubCategoryGoodsListRelationDefault creates NrSubCategoryGoodsListRelationDefault with default headers values
func NewNrSubCategoryGoodsListRelationDefault(code int) *NrSubCategoryGoodsListRelationDefault {
	if code <= 0 {
		code = 500
	}

	return &NrSubCategoryGoodsListRelationDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the sub category goods list relation default response
func (o *NrSubCategoryGoodsListRelationDefault) WithStatusCode(code int) *NrSubCategoryGoodsListRelationDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the sub category goods list relation default response
func (o *NrSubCategoryGoodsListRelationDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the sub category goods list relation default response
func (o *NrSubCategoryGoodsListRelationDefault) WithPayload(payload *models.Error) *NrSubCategoryGoodsListRelationDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the sub category goods list relation default response
func (o *NrSubCategoryGoodsListRelationDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NrSubCategoryGoodsListRelationDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
