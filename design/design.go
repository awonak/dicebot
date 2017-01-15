package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("dicebot", func() {
	Title("Dice Roll Chatbot ")
	Description("A Chatbot for rolling dice")
	Scheme("http")
	Host("localhost:8080")
})

var _ = Resource("dice", func() {
	BasePath("/")
	DefaultMedia(RollMedia)

	Action("index", func() {
		Description("GET landing page")
		Routing(GET("/:rollPattern"))
		Params(func() {
			Param("rollPattern", String, "Roll Pattern")
		})
		Response(OK)
		Response(BadRequest, ErrorMedia)
		Response(NotFound)
	})

	Action("roll", func() {
		Description("Roll the dice")
		Routing(POST("/"))
		Payload(func() {
			Member("text")
			Required("text")
		})
		Response(OK)
		Response(BadRequest, ErrorMedia)
		Response(NotFound)
	})
})

var RollMedia = MediaType("application/vnd.goa.diceroll+json", func() {
	Description("Roll response message")
	Attributes(func() {
		Attribute("text", String, "Roll response text")
		Required("text")
	})
	View("default", func() {
		Attribute("text")
	})
})
