// Code generated by go-swagger; DO NOT EDIT.

package community

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "youbite/models"
)

// CommunityDetailOKCode is the HTTP code returned for type CommunityDetailOK
const CommunityDetailOKCode int = 200

/*CommunityDetailOK 圈子详情

swagger:response communityDetailOK
*/
type CommunityDetailOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse20013 `json:"body,omitempty"`
}

// NewCommunityDetailOK creates CommunityDetailOK with default headers values
func NewCommunityDetailOK() *CommunityDetailOK {
	return &CommunityDetailOK{}
}

// WithPayload adds the payload to the community detail o k response
func (o *CommunityDetailOK) WithPayload(payload *models.InlineResponse20013) *CommunityDetailOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the community detail o k response
func (o *CommunityDetailOK) SetPayload(payload *models.InlineResponse20013) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CommunityDetailOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*NrCommunityDetailDefault error

swagger:response communityDetailDefault
*/
type NrCommunityDetailDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewNrCommunityDetailDefault creates NrCommunityDetailDefault with default headers values
func NewNrCommunityDetailDefault(code int) *NrCommunityDetailDefault {
	if code <= 0 {
		code = 500
	}

	return &NrCommunityDetailDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the community detail default response
func (o *NrCommunityDetailDefault) WithStatusCode(code int) *NrCommunityDetailDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the community detail default response
func (o *NrCommunityDetailDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the community detail default response
func (o *NrCommunityDetailDefault) WithPayload(payload *models.Error) *NrCommunityDetailDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the community detail default response
func (o *NrCommunityDetailDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NrCommunityDetailDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}