package native

import (
	"zenrailz/log"
	"zenrailz/repository"
)

type Service struct {
	logger       log.Logger
	databaseRepo repository.Database
}

type SystemState struct {
	Status   string `json:"status"`
	Database string `json:"database"`
}

type DatabaseState struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}
