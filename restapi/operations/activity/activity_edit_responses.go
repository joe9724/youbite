// Code generated by go-swagger; DO NOT EDIT.

package activity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "youbite/models"
)

// ActivityEditOKCode is the HTTP code returned for type ActivityEditOK
const ActivityEditOKCode int = 200

/*ActivityEditOK 添加活动成功，返回成功信息

swagger:response activityEditOK
*/
type ActivityEditOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse2003 `json:"body,omitempty"`
}

// NewActivityEditOK creates ActivityEditOK with default headers values
func NewActivityEditOK() *ActivityEditOK {
	return &ActivityEditOK{}
}

// WithPayload adds the payload to the activity edit o k response
func (o *ActivityEditOK) WithPayload(payload *models.InlineResponse2003) *ActivityEditOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the activity edit o k response
func (o *ActivityEditOK) SetPayload(payload *models.InlineResponse2003) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ActivityEditOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ActivityEditDefault error

swagger:response activityEditDefault
*/
type ActivityEditDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewActivityEditDefault creates ActivityEditDefault with default headers values
func NewActivityEditDefault(code int) *ActivityEditDefault {
	if code <= 0 {
		code = 500
	}

	return &ActivityEditDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the activity edit default response
func (o *ActivityEditDefault) WithStatusCode(code int) *ActivityEditDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the activity edit default response
func (o *ActivityEditDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the activity edit default response
func (o *ActivityEditDefault) WithPayload(payload *models.Error) *ActivityEditDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the activity edit default response
func (o *ActivityEditDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ActivityEditDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}