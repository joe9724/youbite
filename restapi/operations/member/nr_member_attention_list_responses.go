// Code generated by go-swagger; DO NOT EDIT.

package member

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "youbite/models"
)

// MemberAttentionListOKCode is the HTTP code returned for type MemberAttentionListOK
const MemberAttentionListOKCode int = 200

/*MemberAttentionListOK 获取反馈列表

swagger:response memberAttentionListOK
*/
type MemberAttentionListOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse2005 `json:"body,omitempty"`
}

// NewMemberAttentionListOK creates MemberAttentionListOK with default headers values
func NewMemberAttentionListOK() *MemberAttentionListOK {
	return &MemberAttentionListOK{}
}

// WithPayload adds the payload to the member attention list o k response
func (o *MemberAttentionListOK) WithPayload(payload *models.InlineResponse2005) *MemberAttentionListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the member attention list o k response
func (o *MemberAttentionListOK) SetPayload(payload *models.InlineResponse2005) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *MemberAttentionListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NrMemberAttentionListDefault error

swagger:response memberAttentionListDefault
*/
type NrMemberAttentionListDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNrMemberAttentionListDefault creates NrMemberAttentionListDefault with default headers values
func NewNrMemberAttentionListDefault(code int) *NrMemberAttentionListDefault {
	if code <= 0 {
		code = 500
	}

	return &NrMemberAttentionListDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the member attention list default response
func (o *NrMemberAttentionListDefault) WithStatusCode(code int) *NrMemberAttentionListDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the member attention list default response
func (o *NrMemberAttentionListDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the member attention list default response
func (o *NrMemberAttentionListDefault) WithPayload(payload *models.Error) *NrMemberAttentionListDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the member attention list default response
func (o *NrMemberAttentionListDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NrMemberAttentionListDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}