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
