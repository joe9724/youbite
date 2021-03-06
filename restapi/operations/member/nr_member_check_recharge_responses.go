// Code generated by go-swagger; DO NOT EDIT.

package member

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "youbite/models"
)

// MemberCheckRechargeOKCode is the HTTP code returned for type MemberCheckRechargeOK
const MemberCheckRechargeOKCode int = 200

/*MemberCheckRechargeOK 登录成功，返回登录信息

swagger:response memberCheckRechargeOK
*/
type MemberCheckRechargeOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse20017 `json:"body,omitempty"`
}

// NewMemberCheckRechargeOK creates MemberCheckRechargeOK with default headers values
func NewMemberCheckRechargeOK() *MemberCheckRechargeOK {
	return &MemberCheckRechargeOK{}
}

// WithPayload adds the payload to the member check recharge o k response
func (o *MemberCheckRechargeOK) WithPayload(payload *models.InlineResponse20017) *MemberCheckRechargeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the member check recharge o k response
func (o *MemberCheckRechargeOK) SetPayload(payload *models.InlineResponse20017) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *MemberCheckRechargeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NrMemberCheckRechargeDefault error

swagger:response memberCheckRechargeDefault
*/
type NrMemberCheckRechargeDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNrMemberCheckRechargeDefault creates NrMemberCheckRechargeDefault with default headers values
func NewNrMemberCheckRechargeDefault(code int) *NrMemberCheckRechargeDefault {
	if code <= 0 {
		code = 500
	}

	return &NrMemberCheckRechargeDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the member check recharge default response
func (o *NrMemberCheckRechargeDefault) WithStatusCode(code int) *NrMemberCheckRechargeDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the member check recharge default response
func (o *NrMemberCheckRechargeDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the member check recharge default response
func (o *NrMemberCheckRechargeDefault) WithPayload(payload *models.Error) *NrMemberCheckRechargeDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the member check recharge default response
func (o *NrMemberCheckRechargeDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NrMemberCheckRechargeDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
