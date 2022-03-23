package nexus

import (
	"net/http"
	"zenrailz/service/railway"

	"github.com/gin-gonic/gin"
)

func RailwayStations() gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := railway.Stations()
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Code:    err.Code,
				Message: "Unable to retrieve stations.",
			})
		} else {
			c.JSON(http.StatusOK, result)
		}
	}
}
