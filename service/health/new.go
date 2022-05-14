package health

import (
	"zenrailz/log"
	"zenrailz/repository"
)

func NewService(logger log.Logger, databaseRepository repository.Database) *Service {
	return &Service{
		logger:       logger,
		databaseRepo: databaseRepository,
	}
}

type Service struct {
	logger       log.Logger
	databaseRepo repository.Database
}
