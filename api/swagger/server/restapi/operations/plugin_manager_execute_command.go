// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PluginManagerExecuteCommandHandlerFunc turns a function with the right signature into a plugin manager execute command handler
type PluginManagerExecuteCommandHandlerFunc func(PluginManagerExecuteCommandParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PluginManagerExecuteCommandHandlerFunc) Handle(params PluginManagerExecuteCommandParams) middleware.Responder {
	return fn(params)
}

// PluginManagerExecuteCommandHandler interface for that can handle valid plugin manager execute command params
type PluginManagerExecuteCommandHandler interface {
	Handle(PluginManagerExecuteCommandParams) middleware.Responder
}

// NewPluginManagerExecuteCommand creates a new http.Handler for the plugin manager execute command operation
func NewPluginManagerExecuteCommand(ctx *middleware.Context, handler PluginManagerExecuteCommandHandler) *PluginManagerExecuteCommand {
	return &PluginManagerExecuteCommand{Context: ctx, Handler: handler}
}

/*
	PluginManagerExecuteCommand swagger:route POST /plugin-manager/{id}/execute pluginManagerExecuteCommand

PluginManagerExecuteCommand plugin manager execute command API
*/
type PluginManagerExecuteCommand struct {
	Context *middleware.Context
	Handler PluginManagerExecuteCommandHandler
}

func (o *PluginManagerExecuteCommand) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPluginManagerExecuteCommandParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PluginManagerExecuteCommandBody plugin manager execute command body
//
// swagger:model PluginManagerExecuteCommandBody
type PluginManagerExecuteCommandBody struct {

	// Command to execute.
	// Required: true
	// Enum: [update stop start restart]
	Command string `json:"command"`
}

// Validate validates this plugin manager execute command body
func (o *PluginManagerExecuteCommandBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCommand(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var pluginManagerExecuteCommandBodyTypeCommandPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["update","stop","start","restart"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pluginManagerExecuteCommandBodyTypeCommandPropEnum = append(pluginManagerExecuteCommandBodyTypeCommandPropEnum, v)
	}
}

const (

	// PluginManagerExecuteCommandBodyCommandUpdate captures enum value "update"
	PluginManagerExecuteCommandBodyCommandUpdate string = "update"

	// PluginManagerExecuteCommandBodyCommandStop captures enum value "stop"
	PluginManagerExecuteCommandBodyCommandStop string = "stop"

	// PluginManagerExecuteCommandBodyCommandStart captures enum value "start"
	PluginManagerExecuteCommandBodyCommandStart string = "start"

	// PluginManagerExecuteCommandBodyCommandRestart captures enum value "restart"
	PluginManagerExecuteCommandBodyCommandRestart string = "restart"
)

// prop value enum
func (o *PluginManagerExecuteCommandBody) validateCommandEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, pluginManagerExecuteCommandBodyTypeCommandPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PluginManagerExecuteCommandBody) validateCommand(formats strfmt.Registry) error {

	if err := validate.RequiredString("body"+"."+"command", "body", o.Command); err != nil {
		return err
	}

	// value enum
	if err := o.validateCommandEnum("body"+"."+"command", "body", o.Command); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this plugin manager execute command body based on context it is used
func (o *PluginManagerExecuteCommandBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PluginManagerExecuteCommandBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PluginManagerExecuteCommandBody) UnmarshalBinary(b []byte) error {
	var res PluginManagerExecuteCommandBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}