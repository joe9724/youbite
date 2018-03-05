// Code generated by go-swagger; DO NOT EDIT.

package member

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "youbite/models"
)

// MemberEditOKCode is the HTTP code returned for type MemberEditOK
const MemberEditOKCode int = 200

/*MemberEditOK 操作成功

swagger:response memberEditOK
*/
type MemberEditOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse2002 `json:"body,omitempty"`
}

// NewMemberEditOK creates MemberEditOK with default headers values
func NewMemberEditOK() *MemberEditOK {
	return &MemberEditOK{}
}

// WithPayload adds the payload to the member edit o k response
func (o *MemberEditOK) WithPayload(payload *models.InlineResponse2002) *MemberEditOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the member edit o k response
func (o *MemberEditOK) SetPayload(payload *models.InlineResponse2002) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *MemberEditOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NrMemberEditDefault error

swagger:response memberEditDefault
*/
type NrMemberEditDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNrMemberEditDefault creates NrMemberEditDefault with default headers values
func NewNrMemberEditDefault(code int) *NrMemberEditDefault {
	if code <= 0 {
		code = 500
	}

	return &NrMemberEditDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the member edit default response
func (o *NrMemberEditDefault) WithStatusCode(code int) *NrMemberEditDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the member edit default response
func (o *NrMemberEditDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the member edit default response
func (o *NrMemberEditDefault) WithPayload(payload *models.Error) *NrMemberEditDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the member edit default response
func (o *NrMemberEditDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NrMemberEditDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}