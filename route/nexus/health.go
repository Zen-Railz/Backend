package nexus

import (
	"net/http"
	"zenrailz/service/health"

	"github.com/gin-gonic/gin"
)

func HealthStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := health.Status()
		c.JSON(http.StatusOK, response)
	}
}

func DatabaseStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := health.Database()
		c.JSON(http.StatusOK, response)
	}
}
