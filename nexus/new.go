package nexus

import "zenrailz/service"

func NewStore(
	healthService service.Health,
	railwayService service.Railway,
) *Store {
	return &Store{
		healthSvc:  healthService,
		railwaySvc: railwayService,
	}
}

type Store struct {
	healthSvc  service.Health
	railwaySvc service.Railway
}
