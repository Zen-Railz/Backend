package system

import (
	"zenrailz/nexus"

	"github.com/gin-gonic/gin"
)

func ConfigureRoute(engine *gin.Engine, nexus *nexus.Store) {
	engine.StaticFile("/favicon.ico", "./static/favicon.ico")
	engine.StaticFile("/", "./static/index.html")

	health := engine.Group("/health")
	health.GET("/", nexus.SystemHealth())
	if gin.IsDebugging() {
		health.GET("/database", nexus.DatabaseHealth())
	}

	railway := engine.Group("/railway")
	railway.GET("/stations", nexus.RailwayStations())
	railway.GET("/lines", nexus.RailwayLines())
	railway.GET("/journey", nexus.RailwayJourney())
}
