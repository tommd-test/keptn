// Code generated by go-swagger; DO NOT EDIT.

package project_resource

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/keptn/keptn/api/models"
)

// PostProjectProjectNameResourceCreatedCode is the HTTP code returned for type PostProjectProjectNameResourceCreated
const PostProjectProjectNameResourceCreatedCode int = 201

/*PostProjectProjectNameResourceCreated Success. Project resources have been uploaded

swagger:response postProjectProjectNameResourceCreated
*/
type PostProjectProjectNameResourceCreated struct {
}

// NewPostProjectProjectNameResourceCreated creates PostProjectProjectNameResourceCreated with default headers values
func NewPostProjectProjectNameResourceCreated() *PostProjectProjectNameResourceCreated {

	return &PostProjectProjectNameResourceCreated{}
}

// WriteResponse to the client
func (o *PostProjectProjectNameResourceCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// PostProjectProjectNameResourceBadRequestCode is the HTTP code returned for type PostProjectProjectNameResourceBadRequest
const PostProjectProjectNameResourceBadRequestCode int = 400

/*PostProjectProjectNameResourceBadRequest Failed. Project resources could not be uploaded

swagger:response postProjectProjectNameResourceBadRequest
*/
type PostProjectProjectNameResourceBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostProjectProjectNameResourceBadRequest creates PostProjectProjectNameResourceBadRequest with default headers values
func NewPostProjectProjectNameResourceBadRequest() *PostProjectProjectNameResourceBadRequest {

	return &PostProjectProjectNameResourceBadRequest{}
}

// WithPayload adds the payload to the post project project name resource bad request response
func (o *PostProjectProjectNameResourceBadRequest) WithPayload(payload *models.Error) *PostProjectProjectNameResourceBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post project project name resource bad request response
func (o *PostProjectProjectNameResourceBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostProjectProjectNameResourceBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PostProjectProjectNameResourceDefault Error

swagger:response postProjectProjectNameResourceDefault
*/
type PostProjectProjectNameResourceDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostProjectProjectNameResourceDefault creates PostProjectProjectNameResourceDefault with default headers values
func NewPostProjectProjectNameResourceDefault(code int) *PostProjectProjectNameResourceDefault {
	if code <= 0 {
		code = 500
	}

	return &PostProjectProjectNameResourceDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post project project name resource default response
func (o *PostProjectProjectNameResourceDefault) WithStatusCode(code int) *PostProjectProjectNameResourceDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post project project name resource default response
func (o *PostProjectProjectNameResourceDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post project project name resource default response
func (o *PostProjectProjectNameResourceDefault) WithPayload(payload *models.Error) *PostProjectProjectNameResourceDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post project project name resource default response
func (o *PostProjectProjectNameResourceDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostProjectProjectNameResourceDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
