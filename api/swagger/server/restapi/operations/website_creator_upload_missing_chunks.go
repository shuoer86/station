// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// WebsiteCreatorUploadMissingChunksHandlerFunc turns a function with the right signature into a website creator upload missing chunks handler
type WebsiteCreatorUploadMissingChunksHandlerFunc func(WebsiteCreatorUploadMissingChunksParams) middleware.Responder

// Handle executing the request and returning a response
func (fn WebsiteCreatorUploadMissingChunksHandlerFunc) Handle(params WebsiteCreatorUploadMissingChunksParams) middleware.Responder {
	return fn(params)
}

// WebsiteCreatorUploadMissingChunksHandler interface for that can handle valid website creator upload missing chunks params
type WebsiteCreatorUploadMissingChunksHandler interface {
	Handle(WebsiteCreatorUploadMissingChunksParams) middleware.Responder
}

// NewWebsiteCreatorUploadMissingChunks creates a new http.Handler for the website creator upload missing chunks operation
func NewWebsiteCreatorUploadMissingChunks(ctx *middleware.Context, handler WebsiteCreatorUploadMissingChunksHandler) *WebsiteCreatorUploadMissingChunks {
	return &WebsiteCreatorUploadMissingChunks{Context: ctx, Handler: handler}
}

/*
	WebsiteCreatorUploadMissingChunks swagger:route POST /websiteCreator/uploadMissingChunks websiteCreatorUploadMissingChunks

WebsiteCreatorUploadMissingChunks website creator upload missing chunks API
*/
type WebsiteCreatorUploadMissingChunks struct {
	Context *middleware.Context
	Handler WebsiteCreatorUploadMissingChunksHandler
}

func (o *WebsiteCreatorUploadMissingChunks) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewWebsiteCreatorUploadMissingChunksParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
