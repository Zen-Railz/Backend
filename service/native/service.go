package native

import (
	"zenrailz/log"
	"zenrailz/repository"
)

type Service struct {
	logger       log.Logger
	databaseRepo repository.Database
}

func NewService(logger log.Logger, databaseRepository repository.Database) *Service {
	return &Service{
		logger:       logger,
		databaseRepo: databaseRepository,
	}
}

const (
	Healthy   = "Healthy"
	Unhealthy = "Unhealthy"
)
