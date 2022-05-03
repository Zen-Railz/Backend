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

type BfeQueueObject struct {
	Visited map[string]struct{}
	Path    []PathPoint
}

type PathPoint struct {
	StationName       string
	StationIdentities []StationIdentity
}

type StationIdentity struct {
	Code string
	Line string
}
