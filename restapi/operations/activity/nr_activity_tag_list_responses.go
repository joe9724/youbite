// Code generated by go-swagger; DO NOT EDIT.

package activity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "youbite/models"
)

// ActivityTagListOKCode is the HTTP code returned for type ActivityTagListOK
const ActivityTagListOKCode int = 200

/*ActivityTagListOK 获取tag列表

swagger:response activityTagListOK
*/
type ActivityTagListOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse200 `json:"body,omitempty"`
}

// NewActivityTagListOK creates ActivityTagListOK with default headers values
func NewActivityTagListOK() *ActivityTagListOK {
	return &ActivityTagListOK{}
}

// WithPayload adds the payload to the activity tag list o k response
func (o *ActivityTagListOK) WithPayload(payload *models.InlineResponse200) *ActivityTagListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the activity tag list o k response
func (o *ActivityTagListOK) SetPayload(payload *models.InlineResponse200) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ActivityTagListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NrActivityTagListDefault error

swagger:response activityTagListDefault
*/
type NrActivityTagListDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNrActivityTagListDefault creates NrActivityTagListDefault with default headers values
func NewNrActivityTagListDefault(code int) *NrActivityTagListDefault {
	if code <= 0 {
		code = 500
	}

	return &NrActivityTagListDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the activity tag list default response
func (o *NrActivityTagListDefault) WithStatusCode(code int) *NrActivityTagListDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the activity tag list default response
func (o *NrActivityTagListDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the activity tag list default response
func (o *NrActivityTagListDefault) WithPayload(payload *models.Error) *NrActivityTagListDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the activity tag list default response
func (o *NrActivityTagListDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NrActivityTagListDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
