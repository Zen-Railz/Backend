package system

import (
	"os"

	"zenrailz/route"

	"github.com/gin-gonic/gin"
)

const ENV_PORT = "PORT"

type Server struct {
	engine *gin.Engine
}

func (s *Server) Run() {
	s.engine = gin.Default()

	route.Configure(s.engine)

	s.engine.Run(":" + os.Getenv(ENV_PORT))
}
