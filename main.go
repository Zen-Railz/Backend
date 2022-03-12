package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome")
	})

	port := os.Getenv("PORT")
	r.Run(":" + port)
}
