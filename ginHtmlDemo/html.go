package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(200, "./index.html", gin.H{
			"title": "首页",
		})
	})
	r.Run()
}
