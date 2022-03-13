package route

import (
	"github.com/gin-gonic/gin"

	"zenrailz/route/nexus"
)

func Configure(engine *gin.Engine) {

	engine.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Welcome to ZenRailz")
	})

	health := engine.Group("/health")
	health.GET("/", nexus.HandleHealthStatus())
}
