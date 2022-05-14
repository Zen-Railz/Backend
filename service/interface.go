package service

import (
	"zenrailz/errorr"
	"zenrailz/service/health"
	"zenrailz/service/railway"
)

type Health interface {
	System() *health.SystemState
	Database() *health.DatabaseState
}

type Railway interface {
	Stations() (interface{}, errorr.Entity)
	Lines() ([]railway.Line, errorr.Entity)
	Journey(originStationName string, destinationStationName string) ([][]railway.PathPoint, errorr.Entity)
}
