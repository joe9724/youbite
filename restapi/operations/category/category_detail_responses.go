// Code generated by go-swagger; DO NOT EDIT.

package category

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "youbite/models"
)

// CategoryDetailOKCode is the HTTP code returned for type CategoryDetailOK
const CategoryDetailOKCode int = 200

/*CategoryDetailOK 分类详情

swagger:response categoryDetailOK
*/
type CategoryDetailOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse20029 `json:"body,omitempty"`
}

// NewCategoryDetailOK creates CategoryDetailOK with default headers values
func NewCategoryDetailOK() *CategoryDetailOK {
	return &CategoryDetailOK{}
}

// WithPayload adds the payload to the category detail o k response
func (o *CategoryDetailOK) WithPayload(payload *models.InlineResponse20029) *CategoryDetailOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the category detail o k response
func (o *CategoryDetailOK) SetPayload(payload *models.InlineResponse20029) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CategoryDetailOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CategoryDetailDefault error

swagger:response categoryDetailDefault
*/
type CategoryDetailDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCategoryDetailDefault creates CategoryDetailDefault with default headers values
func NewCategoryDetailDefault(code int) *CategoryDetailDefault {
	if code <= 0 {
		code = 500
	}

	return &CategoryDetailDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the category detail default response
func (o *CategoryDetailDefault) WithStatusCode(code int) *CategoryDetailDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the category detail default response
func (o *CategoryDetailDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the category detail default response
func (o *CategoryDetailDefault) WithPayload(payload *models.Error) *CategoryDetailDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the category detail default response
func (o *CategoryDetailDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CategoryDetailDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
