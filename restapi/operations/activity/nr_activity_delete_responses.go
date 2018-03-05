// Code generated by go-swagger; DO NOT EDIT.

package activity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "youbite/models"
)

// ActivityDeleteOKCode is the HTTP code returned for type ActivityDeleteOK
const ActivityDeleteOKCode int = 200

/*ActivityDeleteOK 活动详情

swagger:response activityDeleteOK
*/
type ActivityDeleteOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse2002 `json:"body,omitempty"`
}

// NewActivityDeleteOK creates ActivityDeleteOK with default headers values
func NewActivityDeleteOK() *ActivityDeleteOK {
	return &ActivityDeleteOK{}
}

// WithPayload adds the payload to the activity delete o k response
func (o *ActivityDeleteOK) WithPayload(payload *models.InlineResponse2002) *ActivityDeleteOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the activity delete o k response
func (o *ActivityDeleteOK) SetPayload(payload *models.InlineResponse2002) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ActivityDeleteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NrActivityDeleteDefault error

swagger:response activityDeleteDefault
*/
type NrActivityDeleteDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNrActivityDeleteDefault creates NrActivityDeleteDefault with default headers values
func NewNrActivityDeleteDefault(code int) *NrActivityDeleteDefault {
	if code <= 0 {
		code = 500
	}

	return &NrActivityDeleteDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the activity delete default response
func (o *NrActivityDeleteDefault) WithStatusCode(code int) *NrActivityDeleteDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the activity delete default response
func (o *NrActivityDeleteDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the activity delete default response
func (o *NrActivityDeleteDefault) WithPayload(payload *models.Error) *NrActivityDeleteDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the activity delete default response
func (o *NrActivityDeleteDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NrActivityDeleteDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
