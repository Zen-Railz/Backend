package railway

import (
	"zenrailz/log"
	"zenrailz/repository"
)

type Service struct {
	logger      log.Logger
	railwayRepo repository.Railway
}

type Line struct {
	Name         string
	Code         string
	Type         string
	IsActive     bool
	Announcement string
}

type Itinerary struct {
	Visited map[string]struct{}
	Path    []PathPoint
	Id      string
}

type PathPoint struct {
	StationName         string
	StationIdentityCode string
	StationLine         string
}
