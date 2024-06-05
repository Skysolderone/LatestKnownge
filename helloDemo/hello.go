package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/go/:name", func(ctx *gin.Context) {
		q := ctx.Param("name")
		ctx.String(200, "Hello %s", q)
	})
	r.Run(":8079")
}
