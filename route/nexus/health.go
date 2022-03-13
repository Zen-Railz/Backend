package nexus

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"zenrailz/service/health"
)

func HandleHealthStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")

		result := health.Status()

		response := map[string]string{
			"status": "success",
			"data":   result,
		}

		c.JSON(http.StatusOK, response)
	}
}
