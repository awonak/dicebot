package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("dicebot", func() {
	Title("Dice Roll Chatbot")
	Description("A Chatbot for rolling dice")
	Scheme("http")
	Host("localhost:8080")
})

var _ = Resource("dice", func() {
	BasePath("/")
	DefaultMedia(DiceRollMedia)

	Action("index", func() {
		Description("GET landing page")
		Routing(GET("/:pattern"))
		Params(func() {
			Param("pattern")
		})
		Response(OK)
		Response(BadRequest)
		Response(NotFound)
	})

	Action("roll", func() {
		Description("Roll the dice")
		Routing(POST("/"))
		Payload(DiceRollPayload)
		Response(OK)
		Response(BadRequest)
		Response(NotFound)
	})
})

var DiceRollPayload = Type("DiceRollPayload", func() {
	Attribute("pattern", String, "Name of thingy", func() {
		MinLength(3)
		MaxLength(6)
		Pattern("^(\\d+)d(\\d+)$")
	})
	Attribute("numDice", Integer, "Number of dice to roll", func() {
		Minimum(1)
		Maximum(6)
	})
	Attribute("sides", Integer, "Number of sides on the dice", func() {
		Minimum(1)
		Maximum(100)
	})
	Attribute("roll", Integer, "The value of a roll for given pattern")
	Required("pattern")
})

var DiceRollMedia = MediaType("application/vnd.goa.diceroll", func() {
	Description("Roll response message")
	ContentType("application/json")
	Reference(DiceRollPayload)
	Attributes(func() {
		Attribute("pattern")
		Attribute("numDice")
		Attribute("sides")
		Attribute("roll")
	})
	View("default", func() {
		Attribute("pattern")
		Attribute("numDice")
		Attribute("sides")
		Attribute("roll")
	})
	View("roll", func() {
		Attribute("roll")
	})
})
