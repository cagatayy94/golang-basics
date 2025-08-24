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

	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(200, "Hello %s", name)
	})

	type Login struct {
		User     string `json:"user" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	r.POST("/login", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if json.User == "admin" && json.Password == "secret" {
			c.JSON(200, gin.H{"status": "you are logged in"})
		} else {
			c.JSON(401, gin.H{"status": "unauthorized"})
		}
	})

	r.Run() // default runs on :8080
}
