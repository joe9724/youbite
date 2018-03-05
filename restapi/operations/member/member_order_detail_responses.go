// Code generated by go-swagger; DO NOT EDIT.

package member

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "youbite/models"
)

// MemberOrderDetailOKCode is the HTTP code returned for type MemberOrderDetailOK
const MemberOrderDetailOKCode int = 200

/*MemberOrderDetailOK 订单详情

swagger:response memberOrderDetailOK
*/
type MemberOrderDetailOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse20019 `json:"body,omitempty"`
}

// NewMemberOrderDetailOK creates MemberOrderDetailOK with default headers values
func NewMemberOrderDetailOK() *MemberOrderDetailOK {
	return &MemberOrderDetailOK{}
}

// WithPayload adds the payload to the member order detail o k response
func (o *MemberOrderDetailOK) WithPayload(payload *models.InlineResponse20019) *MemberOrderDetailOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the member order detail o k response
func (o *MemberOrderDetailOK) SetPayload(payload *models.InlineResponse20019) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *MemberOrderDetailOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*MemberOrderDetailDefault error

swagger:response memberOrderDetailDefault
*/
type MemberOrderDetailDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewMemberOrderDetailDefault creates MemberOrderDetailDefault with default headers values
func NewMemberOrderDetailDefault(code int) *MemberOrderDetailDefault {
	if code <= 0 {
		code = 500
	}

	return &MemberOrderDetailDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the member order detail default response
func (o *MemberOrderDetailDefault) WithStatusCode(code int) *MemberOrderDetailDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the member order detail default response
func (o *MemberOrderDetailDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the member order detail default response
func (o *MemberOrderDetailDefault) WithPayload(payload *models.Error) *MemberOrderDetailDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the member order detail default response
func (o *MemberOrderDetailDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *MemberOrderDetailDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
