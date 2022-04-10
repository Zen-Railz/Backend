package railway

import (
	"zenrailz/log"
	"zenrailz/repository"
)

type Service struct {
	logger      log.Logger
	railwayRepo repository.Railway
}

func NewService(logger log.Logger, railwayRepository repository.Railway) *Service {
	return &Service{
		logger:      logger,
		railwayRepo: railwayRepository,
	}
}
