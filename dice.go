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

// Roll runs the roll action.
func (c *DiceController) Roll(ctx *app.RollDiceContext) error {
	// Get the dice roll pattern from the request context
	pattern := ctx.Payload.Text

	d, _ := dice.NewDiceRoll(pattern)

	res := &app.GoaDiceroll{
		Text: fmt.Sprintf("Rolling Dice [%v]: %v", pattern, d.Roll()),
	}
	return ctx.OK(res)
}
