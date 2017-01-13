// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=dicebot/design
// --out=$(GOPATH)/src/dicebot
// --version=v1.1.0-dirty
//
// API "dicebot": Application Contexts
//
// The content of this file is auto-generated, DO NOT MODIFY

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
)

// RollDiceContext provides the dice roll action context.
type RollDiceContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *RollDicePayload
}

// NewRollDiceContext parses the incoming request URL and body, performs validations and creates the
// context used by the dice controller roll action.
func NewRollDiceContext(ctx context.Context, service *goa.Service) (*RollDiceContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := RollDiceContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// rollDicePayload is the dice roll action payload.
type rollDicePayload struct {
	// Roll response text
	Text *string `form:"text,omitempty" json:"text,omitempty" xml:"text,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *rollDicePayload) Validate() (err error) {
	if payload.Text == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "text"))
	}
	return
}

// Publicize creates RollDicePayload from rollDicePayload
func (payload *rollDicePayload) Publicize() *RollDicePayload {
	var pub RollDicePayload
	if payload.Text != nil {
		pub.Text = *payload.Text
	}
	return &pub
}

// RollDicePayload is the dice roll action payload.
type RollDicePayload struct {
	// Roll response text
	Text string `form:"text" json:"text" xml:"text"`
}

// Validate runs the validation rules defined in the design.
func (payload *RollDicePayload) Validate() (err error) {
	if payload.Text == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`raw`, "text"))
	}
	return
}

// OK sends a HTTP response with status code 200.
func (ctx *RollDiceContext) OK(r *GoaDiceroll) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.diceroll+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *RollDiceContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}