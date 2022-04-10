package nexus

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Store) SystemHealth() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := s.nativeSvc.SystemHealth()
		c.JSON(http.StatusOK, response)
	}
}

func (s *Store) DatabaseHealth() gin.HandlerFunc {
	return func(c *gin.Context) {
		response := s.nativeSvc.DatabaseHealth()
		c.JSON(http.StatusOK, response)
	}
}
