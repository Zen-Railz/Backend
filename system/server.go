package system

import (
	"zenrailz/environment"
	"zenrailz/log"
	"zenrailz/repository/database"

	"github.com/gin-gonic/gin"
)

func Start() {
	log.SetLevel()
	logger := log.New()

	defer recuperate(logger)

	run(logger)
}

func run(logger log.Logger) {
	db, err := database.New()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	engine := gin.Default()
	nexus := InitialiseNexus(logger, db)

	ConfigureRoute(engine, nexus)

	port, err := environment.ServerPort()

	if err == nil {
		engine.Run(":" + port)
	} else {
		panic(err)
	}
}

func recuperate(logger log.Logger) {
	if err := recover(); err != nil {
		logger.Error("Panic", err)
	}
}
