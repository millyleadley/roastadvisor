package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("roastadvisor", func() {
	Title("Roast Advisor API")
	Description("Returns reviews of Sunday roasts")
	Server("roastadvisor", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
		})
	})
})

var _ = Service("Reviews", func() {
	Description("The reviews service provides operations to manage reviews.")

	Method("List", func() {
		Description("List all reviews.")

		Result(ArrayOf(Review))

		HTTP(func() {
			GET("/reviews")
			Response(StatusOK)
		})
	})
})

var Review = Type("Review", func() {
	Attribute("id", String, "Unique internal ID of the incident debrief", func() {
		Example("01G0J1EXE7AXZ2C93K61WBPYEH")
	})
	Attribute("user_id", String, "ID of the user who left the review", func() {
		Example("01G0J1EXE7AXZ2C93K61WBPYEH")
	})
	Attribute("restaurant_id", String, "ID of the restaurant the review is for", func() {
		Example("01G0J1EXE7AXZ2C93K61WBPYEH")
	})
	Required("id", "user_id", "restaurant_id")
})
