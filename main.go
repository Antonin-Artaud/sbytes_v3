package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	ginServer := gin.Default()

	ginServer.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Space!",
		})
	})

	_ = ginServer.Run(":8080")
}
