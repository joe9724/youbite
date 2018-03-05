// Code generated by go-swagger; DO NOT EDIT.

package member

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "youbite/models"
)

// MemberRegisterSendSmsOKCode is the HTTP code returned for type MemberRegisterSendSmsOK
const MemberRegisterSendSmsOKCode int = 200

/*MemberRegisterSendSmsOK 登录成功，返回登录信息

swagger:response memberRegisterSendSmsOK
*/
type MemberRegisterSendSmsOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse2002 `json:"body,omitempty"`
}

// NewMemberRegisterSendSmsOK creates MemberRegisterSendSmsOK with default headers values
func NewMemberRegisterSendSmsOK() *MemberRegisterSendSmsOK {
	return &MemberRegisterSendSmsOK{}
}

// WithPayload adds the payload to the member register send sms o k response
func (o *MemberRegisterSendSmsOK) WithPayload(payload *models.InlineResponse2002) *MemberRegisterSendSmsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the member register send sms o k response
func (o *MemberRegisterSendSmsOK) SetPayload(payload *models.InlineResponse2002) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *MemberRegisterSendSmsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NrMemberRegisterSendSmsDefault error

swagger:response memberRegisterSendSmsDefault
*/
type NrMemberRegisterSendSmsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNrMemberRegisterSendSmsDefault creates NrMemberRegisterSendSmsDefault with default headers values
func NewNrMemberRegisterSendSmsDefault(code int) *NrMemberRegisterSendSmsDefault {
	if code <= 0 {
		code = 500
	}

	return &NrMemberRegisterSendSmsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the member register send sms default response
func (o *NrMemberRegisterSendSmsDefault) WithStatusCode(code int) *NrMemberRegisterSendSmsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the member register send sms default response
func (o *NrMemberRegisterSendSmsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the member register send sms default response
func (o *NrMemberRegisterSendSmsDefault) WithPayload(payload *models.Error) *NrMemberRegisterSendSmsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the member register send sms default response
func (o *NrMemberRegisterSendSmsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NrMemberRegisterSendSmsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
