package repository

import (
	"zenrailz/errorr"
	"zenrailz/repository/railway"
)

type Database interface {
	Ping() errorr.Entity
}

type Railway interface {
	Stations() ([]railway.Station, errorr.Entity)
	Lines() ([]railway.Line, errorr.Entity)
}
