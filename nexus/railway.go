package nexus

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Store) RailwayStations() gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := s.railwaySvc.Stations()
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
