package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//docker run -p 8080:8080 -it ubuntu:20.04 bash
//
// root@b6d2f58860c7:/#
//
//docker cp myserver b6d2f58860c7:/
//
//docker run
// root@b6d2f58860c7:/# ulimit -n 25

// ## started the binary
// root@b6d2f58860c7:/# ./myserver

// test
// ./goloris.exe -victimUrl http://127.0.0.1:8080/greet
func main() {
	// Create a new Gin router
	router := gin.Default()
	// router.Use(gin.MaxReadTimeout(2 * time.Second))
	// Define a route for the index page
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	router.POST("/greet", func(c *gin.Context) {
		// Get the username from the request body
		var requestBody struct {
			Username string `json:"username"`
		}
		if err := c.BindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Respond with a greeting
		greeting := "Hello, " + requestBody.Username + "!"
		c.JSON(http.StatusOK, gin.H{"message": greeting})
	})
	// Run the server
	router.Run(":8080")

	// 改造
	// Set server timeouts
	// 疑似需把上述gin 路由全改为使用http
	// server := &http.Server{
	// 	Addr:         ":8080",
	// 	Handler:      router,
	// 	ReadTimeout:  10 * time.Second, // Set a maximum time to read the request headers
	// 	WriteTimeout: 10 * time.Second, // Set a maximum time to write the response
	// }
	// server.ListenAndServe()
}
