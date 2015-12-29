package main

import (
	"ginlime/app"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// call priveet method
	router.GET("/hello1/:name", hello)

	// set direct function
	router.GET("/hello2/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(200, "hello-2 %s!\n", name)
	})

	// call other package
	router.GET("/hello3/:name", app.Hello)
	router.GET("/hobbies/:person_name", app.Hobbies)

	router.Run(":8080")
}

func hello(c *gin.Context) {
	name := c.Param("name")
	c.String(200, "hello-1 %s!\n", name)
}
