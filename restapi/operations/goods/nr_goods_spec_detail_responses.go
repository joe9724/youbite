// Code generated by go-swagger; DO NOT EDIT.

package goods

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "youbite/models"
)

// GoodsSpecDetailOKCode is the HTTP code returned for type GoodsSpecDetailOK
const GoodsSpecDetailOKCode int = 200

/*GoodsSpecDetailOK 商品规格详情

swagger:response goodsSpecDetailOK
*/
type GoodsSpecDetailOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse20033 `json:"body,omitempty"`
}

// NewGoodsSpecDetailOK creates GoodsSpecDetailOK with default headers values
func NewGoodsSpecDetailOK() *GoodsSpecDetailOK {
	return &GoodsSpecDetailOK{}
}

// WithPayload adds the payload to the goods spec detail o k response
func (o *GoodsSpecDetailOK) WithPayload(payload *models.InlineResponse20033) *GoodsSpecDetailOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the goods spec detail o k response
func (o *GoodsSpecDetailOK) SetPayload(payload *models.InlineResponse20033) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GoodsSpecDetailOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NrGoodsSpecDetailDefault error

swagger:response goodsSpecDetailDefault
*/
type NrGoodsSpecDetailDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNrGoodsSpecDetailDefault creates NrGoodsSpecDetailDefault with default headers values
func NewNrGoodsSpecDetailDefault(code int) *NrGoodsSpecDetailDefault {
	if code <= 0 {
		code = 500
	}

	return &NrGoodsSpecDetailDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the goods spec detail default response
func (o *NrGoodsSpecDetailDefault) WithStatusCode(code int) *NrGoodsSpecDetailDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the goods spec detail default response
func (o *NrGoodsSpecDetailDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the goods spec detail default response
func (o *NrGoodsSpecDetailDefault) WithPayload(payload *models.Error) *NrGoodsSpecDetailDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the goods spec detail default response
func (o *NrGoodsSpecDetailDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NrGoodsSpecDetailDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
