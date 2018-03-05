// Code generated by go-swagger; DO NOT EDIT.

package feedback

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "youbite/models"
)

// FeedbackDetailOKCode is the HTTP code returned for type FeedbackDetailOK
const FeedbackDetailOKCode int = 200

/*FeedbackDetailOK 反馈详情

swagger:response feedbackDetailOK
*/
type FeedbackDetailOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse20021 `json:"body,omitempty"`
}

// NewFeedbackDetailOK creates FeedbackDetailOK with default headers values
func NewFeedbackDetailOK() *FeedbackDetailOK {
	return &FeedbackDetailOK{}
}

// WithPayload adds the payload to the feedback detail o k response
func (o *FeedbackDetailOK) WithPayload(payload *models.InlineResponse20021) *FeedbackDetailOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the feedback detail o k response
func (o *FeedbackDetailOK) SetPayload(payload *models.InlineResponse20021) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FeedbackDetailOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*FeedbackDetailDefault error

swagger:response feedbackDetailDefault
*/
type FeedbackDetailDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewFeedbackDetailDefault creates FeedbackDetailDefault with default headers values
func NewFeedbackDetailDefault(code int) *FeedbackDetailDefault {
	if code <= 0 {
		code = 500
	}

	return &FeedbackDetailDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the feedback detail default response
func (o *FeedbackDetailDefault) WithStatusCode(code int) *FeedbackDetailDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the feedback detail default response
func (o *FeedbackDetailDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the feedback detail default response
func (o *FeedbackDetailDefault) WithPayload(payload *models.Error) *FeedbackDetailDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the feedback detail default response
func (o *FeedbackDetailDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FeedbackDetailDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
