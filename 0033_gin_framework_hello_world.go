package main

import (
	//install it
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/hi", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	r.Run() // default runs on :8080
}
