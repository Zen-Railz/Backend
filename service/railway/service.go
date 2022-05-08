package railway

import (
	"zenrailz/log"
	"zenrailz/repository"
)

func NewService(logger log.Logger, configurationRepository repository.Configuration, railwayRepository repository.Railway) *Service {
	return &Service{
		logger:      logger,
		configRepo:  configurationRepository,
		railwayRepo: railwayRepository,
	}
}
