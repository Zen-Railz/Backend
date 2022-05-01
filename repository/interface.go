package repository

import (
	"zenrailz/errorr"
	"zenrailz/repository/railway"
)

type Database interface {
	Ping() errorr.Entity
}

type Railway interface {
	Stations() (map[string]railway.Station, errorr.Entity)
	Lines() ([]railway.Line, errorr.Entity)
	Network() (map[string]*railway.NetworkNode, map[string]*railway.NetworkNode, errorr.Entity)
}
