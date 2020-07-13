package main

import (

	"github.com/gin-gonic/gin"

)

func sayHello(c *gin.Context)  {
	c.JSON(200,gin.H{
		"message" : "hello",
	})
}
func main() {
	r := gin.Default()
	r.GET("/hello",sayHello)
	r.GET("/book")
	r.POST("/book")
	r.PUT("/book")
	r.DELETE("/book")

	r.Run(":9090")
}
