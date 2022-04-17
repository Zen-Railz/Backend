package railway

import (
	"zenrailz/log"
	"zenrailz/repository"
)

func NewService(logger log.Logger, railwayRepository repository.Railway) *Service {
	return &Service{
		logger:      logger,
		railwayRepo: railwayRepository,
	}
}
