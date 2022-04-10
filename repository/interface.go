package repository

import (
	"zenrailz/anomaly"
	"zenrailz/repository/railway"
)

type Database interface {
	Ping() *anomaly.ServiceError
}

type Railway interface {
	Stations() ([]railway.Station, *anomaly.ServiceError)
}
