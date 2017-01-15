package main

import (
	"fmt"

	"dicebot/app"
	"dicebot/dice"
	"github.com/goadesign/goa"
)

// DiceController implements the dice resource.
type DiceController struct {
	*goa.Controller
}

// NewDiceController creates a dice controller.
func NewDiceController(service *goa.Service) *DiceController {
	return &DiceController{Controller: service.NewController("DiceController")}
}

// Roll runs the roll action for GET requests.
func (c *DiceController) Index(ctx *app.IndexDiceContext) error {
	// Get the dice roll pattern from the request context
	pattern := ctx.RollPattern
	res := getDiceRollResponse(pattern)
	return ctx.OK(res)
}

// Roll runs the roll action for POST requests.
func (c *DiceController) Roll(ctx *app.RollDiceContext) error {
	// Get the dice roll pattern from the request context
	pattern := ctx.Payload.Text
	res := getDiceRollResponse(pattern)
	return ctx.OK(res)
}

// getDiceRollResponse takes a pattern, creates a DiceRoll and formulates a
// GoaDiceroll response
func getDiceRollResponse(pattern string) *app.GoaDiceroll {
	d, _ := dice.NewDiceRoll(pattern)

	res := &app.GoaDiceroll{
		Text: fmt.Sprintf("Rolling Dice [%v]: %v", pattern, d.Roll()),
	}

	return res
}
