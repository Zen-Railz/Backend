package mock

import (
	"zenrailz/errorr"
	"zenrailz/repository/railway"
)

func NewRailwayRepository() *RailwayRepository {
	return &RailwayRepository{}
}

type RailwayRepository struct {
	stations    []railway.Station
	sourceError errorr.Entity
}

func (r *RailwayRepository) Stations() ([]railway.Station, errorr.Entity) {
	return r.stations, r.sourceError
}

func (r *RailwayRepository) EmptyStations() *RailwayRepository {
	r.stations = []railway.Station{}
	return r
}

func (r *RailwayRepository) AddStation(name string, code string, number int) *RailwayRepository {
	r.stations = append(r.stations, railway.Station{
		Name:   name,
		Code:   code,
		Number: number,
	})
	return r
}

func (r *RailwayRepository) SetSourceError() *RailwayRepository {
	r.sourceError = NewError().
		SetCode(ErrorCode)
	return r
}
