// Code generated by go-swagger; DO NOT EDIT.

package virtualserver

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PutVsVipPortAllowHandlerFunc turns a function with the right signature into a put vs vip port allow handler
type PutVsVipPortAllowHandlerFunc func(PutVsVipPortAllowParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutVsVipPortAllowHandlerFunc) Handle(params PutVsVipPortAllowParams) middleware.Responder {
	return fn(params)
}

// PutVsVipPortAllowHandler interface for that can handle valid put vs vip port allow params
type PutVsVipPortAllowHandler interface {
	Handle(PutVsVipPortAllowParams) middleware.Responder
}

// NewPutVsVipPortAllow creates a new http.Handler for the put vs vip port allow operation
func NewPutVsVipPortAllow(ctx *middleware.Context, handler PutVsVipPortAllowHandler) *PutVsVipPortAllow {
	return &PutVsVipPortAllow{Context: ctx, Handler: handler}
}

/*
	PutVsVipPortAllow swagger:route PUT /vs/{VipPort}/allow virtualserver putVsVipPortAllow

Add a set of ip from white list to vip:port:proto
*/
type PutVsVipPortAllow struct {
	Context *middleware.Context
	Handler PutVsVipPortAllowHandler
}

func (o *PutVsVipPortAllow) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPutVsVipPortAllowParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
