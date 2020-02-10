// Code generated by go-swagger; DO NOT EDIT.

package service_resource

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/keptn/keptn/api/models"
)

// PostProjectProjectNameStageStageNameServiceServiceNameResourceCreatedCode is the HTTP code returned for type PostProjectProjectNameStageStageNameServiceServiceNameResourceCreated
const PostProjectProjectNameStageStageNameServiceServiceNameResourceCreatedCode int = 201

/*PostProjectProjectNameStageStageNameServiceServiceNameResourceCreated Success. Service resources have been uploaded

swagger:response postProjectProjectNameStageStageNameServiceServiceNameResourceCreated
*/
type PostProjectProjectNameStageStageNameServiceServiceNameResourceCreated struct {
}

// NewPostProjectProjectNameStageStageNameServiceServiceNameResourceCreated creates PostProjectProjectNameStageStageNameServiceServiceNameResourceCreated with default headers values
func NewPostProjectProjectNameStageStageNameServiceServiceNameResourceCreated() *PostProjectProjectNameStageStageNameServiceServiceNameResourceCreated {

	return &PostProjectProjectNameStageStageNameServiceServiceNameResourceCreated{}
}

// WriteResponse to the client
func (o *PostProjectProjectNameStageStageNameServiceServiceNameResourceCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// PostProjectProjectNameStageStageNameServiceServiceNameResourceBadRequestCode is the HTTP code returned for type PostProjectProjectNameStageStageNameServiceServiceNameResourceBadRequest
const PostProjectProjectNameStageStageNameServiceServiceNameResourceBadRequestCode int = 400

/*PostProjectProjectNameStageStageNameServiceServiceNameResourceBadRequest Failed. Service resources could not be uploaded

swagger:response postProjectProjectNameStageStageNameServiceServiceNameResourceBadRequest
*/
type PostProjectProjectNameStageStageNameServiceServiceNameResourceBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostProjectProjectNameStageStageNameServiceServiceNameResourceBadRequest creates PostProjectProjectNameStageStageNameServiceServiceNameResourceBadRequest with default headers values
func NewPostProjectProjectNameStageStageNameServiceServiceNameResourceBadRequest() *PostProjectProjectNameStageStageNameServiceServiceNameResourceBadRequest {

	return &PostProjectProjectNameStageStageNameServiceServiceNameResourceBadRequest{}
}

// WithPayload adds the payload to the post project project name stage stage name service service name resource bad request response
func (o *PostProjectProjectNameStageStageNameServiceServiceNameResourceBadRequest) WithPayload(payload *models.Error) *PostProjectProjectNameStageStageNameServiceServiceNameResourceBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post project project name stage stage name service service name resource bad request response
func (o *PostProjectProjectNameStageStageNameServiceServiceNameResourceBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostProjectProjectNameStageStageNameServiceServiceNameResourceBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PostProjectProjectNameStageStageNameServiceServiceNameResourceDefault Error

swagger:response postProjectProjectNameStageStageNameServiceServiceNameResourceDefault
*/
type PostProjectProjectNameStageStageNameServiceServiceNameResourceDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostProjectProjectNameStageStageNameServiceServiceNameResourceDefault creates PostProjectProjectNameStageStageNameServiceServiceNameResourceDefault with default headers values
func NewPostProjectProjectNameStageStageNameServiceServiceNameResourceDefault(code int) *PostProjectProjectNameStageStageNameServiceServiceNameResourceDefault {
	if code <= 0 {
		code = 500
	}

	return &PostProjectProjectNameStageStageNameServiceServiceNameResourceDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post project project name stage stage name service service name resource default response
func (o *PostProjectProjectNameStageStageNameServiceServiceNameResourceDefault) WithStatusCode(code int) *PostProjectProjectNameStageStageNameServiceServiceNameResourceDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post project project name stage stage name service service name resource default response
func (o *PostProjectProjectNameStageStageNameServiceServiceNameResourceDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post project project name stage stage name service service name resource default response
func (o *PostProjectProjectNameStageStageNameServiceServiceNameResourceDefault) WithPayload(payload *models.Error) *PostProjectProjectNameStageStageNameServiceServiceNameResourceDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post project project name stage stage name service service name resource default response
func (o *PostProjectProjectNameStageStageNameServiceServiceNameResourceDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostProjectProjectNameStageStageNameServiceServiceNameResourceDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
