// Code generated by go-swagger; DO NOT EDIT.

package member

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "youbite/models"
)

// MemberDetailOKCode is the HTTP code returned for type MemberDetailOK
const MemberDetailOKCode int = 200

/*MemberDetailOK 会员详情

swagger:response memberDetailOK
*/
type MemberDetailOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse20018 `json:"body,omitempty"`
}

// NewMemberDetailOK creates MemberDetailOK with default headers values
func NewMemberDetailOK() *MemberDetailOK {
	return &MemberDetailOK{}
}

// WithPayload adds the payload to the member detail o k response
func (o *MemberDetailOK) WithPayload(payload *models.InlineResponse20018) *MemberDetailOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the member detail o k response
func (o *MemberDetailOK) SetPayload(payload *models.InlineResponse20018) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *MemberDetailOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*MemberDetailDefault error

swagger:response memberDetailDefault
*/
type MemberDetailDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewMemberDetailDefault creates MemberDetailDefault with default headers values
func NewMemberDetailDefault(code int) *MemberDetailDefault {
	if code <= 0 {
		code = 500
	}

	return &MemberDetailDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the member detail default response
func (o *MemberDetailDefault) WithStatusCode(code int) *MemberDetailDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the member detail default response
func (o *MemberDetailDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the member detail default response
func (o *MemberDetailDefault) WithPayload(payload *models.Error) *MemberDetailDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the member detail default response
func (o *MemberDetailDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *MemberDetailDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
