// Code generated by go-swagger; DO NOT EDIT.

package activity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "youbite/models"
)

// ActivityTagUploadOKCode is the HTTP code returned for type ActivityTagUploadOK
const ActivityTagUploadOKCode int = 200

/*ActivityTagUploadOK 添加活动成功，返回成功信息

swagger:response activityTagUploadOK
*/
type ActivityTagUploadOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse2003 `json:"body,omitempty"`
}

// NewActivityTagUploadOK creates ActivityTagUploadOK with default headers values
func NewActivityTagUploadOK() *ActivityTagUploadOK {
	return &ActivityTagUploadOK{}
}

// WithPayload adds the payload to the activity tag upload o k response
func (o *ActivityTagUploadOK) WithPayload(payload *models.InlineResponse2003) *ActivityTagUploadOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the activity tag upload o k response
func (o *ActivityTagUploadOK) SetPayload(payload *models.InlineResponse2003) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ActivityTagUploadOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ActivityTagUploadDefault error

swagger:response activityTagUploadDefault
*/
type ActivityTagUploadDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewActivityTagUploadDefault creates ActivityTagUploadDefault with default headers values
func NewActivityTagUploadDefault(code int) *ActivityTagUploadDefault {
	if code <= 0 {
		code = 500
	}

	return &ActivityTagUploadDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the activity tag upload default response
func (o *ActivityTagUploadDefault) WithStatusCode(code int) *ActivityTagUploadDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the activity tag upload default response
func (o *ActivityTagUploadDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the activity tag upload default response
func (o *ActivityTagUploadDefault) WithPayload(payload *models.Error) *ActivityTagUploadDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the activity tag upload default response
func (o *ActivityTagUploadDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ActivityTagUploadDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
