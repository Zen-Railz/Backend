package route

import (
	"zenrailz/route/nexus"

	"github.com/gin-gonic/gin"
)

func Configure(engine *gin.Engine) {

	engine.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Welcome to ZenRailz")
	})

	health := engine.Group("/health")
	health.GET("/", nexus.HealthStatus())
	if gin.IsDebugging() {
		health.GET("/database", nexus.DatabaseStatus())
	}

	railway := engine.Group("/railway")
	railway.GET("/stations", nexus.RailwayStations())
}
