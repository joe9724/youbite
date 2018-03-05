// Code generated by go-swagger; DO NOT EDIT.

package banner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "youbite/models"
)

// BannerListOKCode is the HTTP code returned for type BannerListOK
const BannerListOKCode int = 200

/*BannerListOK Banner列表

swagger:response bannerListOK
*/
type BannerListOK struct {

	/*
	  In: Body
	*/
	Payload *models.InlineResponse20024 `json:"body,omitempty"`
}

// NewBannerListOK creates BannerListOK with default headers values
func NewBannerListOK() *BannerListOK {
	return &BannerListOK{}
}

// WithPayload adds the payload to the banner list o k response
func (o *BannerListOK) WithPayload(payload *models.InlineResponse20024) *BannerListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the banner list o k response
func (o *BannerListOK) SetPayload(payload *models.InlineResponse20024) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BannerListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*BannerListDefault error

swagger:response bannerListDefault
*/
type BannerListDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewBannerListDefault creates BannerListDefault with default headers values
func NewBannerListDefault(code int) *BannerListDefault {
	if code <= 0 {
		code = 500
	}

	return &BannerListDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the banner list default response
func (o *BannerListDefault) WithStatusCode(code int) *BannerListDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the banner list default response
func (o *BannerListDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the banner list default response
func (o *BannerListDefault) WithPayload(payload *models.Error) *BannerListDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the banner list default response
func (o *BannerListDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BannerListDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}