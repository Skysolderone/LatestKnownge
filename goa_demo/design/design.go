package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("hello", func() {
	Description("The hello service says hello")
	Method("say", func() {
		Payload(String, "name or message")
		Result(String, "this is a result")
		HTTP(func() {
			GET("/hello/{name}")
		})
	})
})
