// Code generated by go-swagger; DO NOT EDIT.

package virtualserver

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/dpvs-agent/models"
)

// DeleteVsVipPortRsOKCode is the HTTP code returned for type DeleteVsVipPortRsOK
const DeleteVsVipPortRsOKCode int = 200

/*
DeleteVsVipPortRsOK Success

swagger:response deleteVsVipPortRsOK
*/
type DeleteVsVipPortRsOK struct {
}

// NewDeleteVsVipPortRsOK creates DeleteVsVipPortRsOK with default headers values
func NewDeleteVsVipPortRsOK() *DeleteVsVipPortRsOK {

	return &DeleteVsVipPortRsOK{}
}

// WriteResponse to the client
func (o *DeleteVsVipPortRsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DeleteVsVipPortRsNotFoundCode is the HTTP code returned for type DeleteVsVipPortRsNotFound
const DeleteVsVipPortRsNotFoundCode int = 404

/*
DeleteVsVipPortRsNotFound Service not found

swagger:response deleteVsVipPortRsNotFound
*/
type DeleteVsVipPortRsNotFound struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewDeleteVsVipPortRsNotFound creates DeleteVsVipPortRsNotFound with default headers values
func NewDeleteVsVipPortRsNotFound() *DeleteVsVipPortRsNotFound {

	return &DeleteVsVipPortRsNotFound{}
}

// WithPayload adds the payload to the delete vs vip port rs not found response
func (o *DeleteVsVipPortRsNotFound) WithPayload(payload string) *DeleteVsVipPortRsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete vs vip port rs not found response
func (o *DeleteVsVipPortRsNotFound) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteVsVipPortRsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DeleteVsVipPortRsInvalidFrontendCode is the HTTP code returned for type DeleteVsVipPortRsInvalidFrontend
const DeleteVsVipPortRsInvalidFrontendCode int = 460

/*
DeleteVsVipPortRsInvalidFrontend Invalid frontend in service configuration

swagger:response deleteVsVipPortRsInvalidFrontend
*/
type DeleteVsVipPortRsInvalidFrontend struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewDeleteVsVipPortRsInvalidFrontend creates DeleteVsVipPortRsInvalidFrontend with default headers values
func NewDeleteVsVipPortRsInvalidFrontend() *DeleteVsVipPortRsInvalidFrontend {

	return &DeleteVsVipPortRsInvalidFrontend{}
}

// WithPayload adds the payload to the delete vs vip port rs invalid frontend response
func (o *DeleteVsVipPortRsInvalidFrontend) WithPayload(payload models.Error) *DeleteVsVipPortRsInvalidFrontend {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete vs vip port rs invalid frontend response
func (o *DeleteVsVipPortRsInvalidFrontend) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteVsVipPortRsInvalidFrontend) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(460)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DeleteVsVipPortRsInvalidBackendCode is the HTTP code returned for type DeleteVsVipPortRsInvalidBackend
const DeleteVsVipPortRsInvalidBackendCode int = 461

/*
DeleteVsVipPortRsInvalidBackend Invalid backend in service configuration

swagger:response deleteVsVipPortRsInvalidBackend
*/
type DeleteVsVipPortRsInvalidBackend struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewDeleteVsVipPortRsInvalidBackend creates DeleteVsVipPortRsInvalidBackend with default headers values
func NewDeleteVsVipPortRsInvalidBackend() *DeleteVsVipPortRsInvalidBackend {

	return &DeleteVsVipPortRsInvalidBackend{}
}

// WithPayload adds the payload to the delete vs vip port rs invalid backend response
func (o *DeleteVsVipPortRsInvalidBackend) WithPayload(payload models.Error) *DeleteVsVipPortRsInvalidBackend {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete vs vip port rs invalid backend response
func (o *DeleteVsVipPortRsInvalidBackend) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteVsVipPortRsInvalidBackend) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(461)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DeleteVsVipPortRsFailureCode is the HTTP code returned for type DeleteVsVipPortRsFailure
const DeleteVsVipPortRsFailureCode int = 500

/*
DeleteVsVipPortRsFailure Service deletion failed

swagger:response deleteVsVipPortRsFailure
*/
type DeleteVsVipPortRsFailure struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewDeleteVsVipPortRsFailure creates DeleteVsVipPortRsFailure with default headers values
func NewDeleteVsVipPortRsFailure() *DeleteVsVipPortRsFailure {

	return &DeleteVsVipPortRsFailure{}
}

// WithPayload adds the payload to the delete vs vip port rs failure response
func (o *DeleteVsVipPortRsFailure) WithPayload(payload models.Error) *DeleteVsVipPortRsFailure {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete vs vip port rs failure response
func (o *DeleteVsVipPortRsFailure) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteVsVipPortRsFailure) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
