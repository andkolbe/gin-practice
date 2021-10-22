package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H {
			"message": "pong",
		})
	})
	r.Run() // serves on :8080 unless a PORT env variable is defined
}