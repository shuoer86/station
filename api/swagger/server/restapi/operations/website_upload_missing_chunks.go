// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WebsiteUploadMissingChunksHandlerFunc turns a function with the right signature into a website upload missing chunks handler
type WebsiteUploadMissingChunksHandlerFunc func(WebsiteUploadMissingChunksParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WebsiteUploadMissingChunksHandlerFunc) Handle(params WebsiteUploadMissingChunksParams) middleware.Responder {
	return fn(params)
}

// WebsiteUploadMissingChunksHandler interface for that can handle valid website upload missing chunks params
type WebsiteUploadMissingChunksHandler interface {
	Handle(WebsiteUploadMissingChunksParams) middleware.Responder
}

// NewWebsiteUploadMissingChunks creates a new http.Handler for the website upload missing chunks operation
func NewWebsiteUploadMissingChunks(ctx *middleware.Context, handler WebsiteUploadMissingChunksHandler) *WebsiteUploadMissingChunks {
	return &WebsiteUploadMissingChunks{Context: ctx, Handler: handler}
}

/*
	WebsiteUploadMissingChunks swagger:route POST /websiteCreator/uploadMissingChunks websiteUploadMissingChunks

WebsiteUploadMissingChunks website upload missing chunks API
*/
type WebsiteUploadMissingChunks struct {
	Context *middleware.Context
	Handler WebsiteUploadMissingChunksHandler
}

func (o *WebsiteUploadMissingChunks) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewWebsiteUploadMissingChunksParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}