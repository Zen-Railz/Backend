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
	lines       []railway.Line
	sourceError errorr.Entity
}

func (r *RailwayRepository) Stations() ([]railway.Station, errorr.Entity) {
	return r.stations, r.sourceError
}

func (r *RailwayRepository) Lines() ([]railway.Line, errorr.Entity) {
	return r.lines, r.sourceError
}

func (r *RailwayRepository) EmptyStations() *RailwayRepository {
	r.stations = []railway.Station{}
	return r
}

func (r *RailwayRepository) AddStation(name string, prefix string, number int) *RailwayRepository {
	r.stations = append(r.stations, railway.Station{
		Name:   name,
		Prefix: prefix,
		Number: number,
	})
	return r
}

func (r *RailwayRepository) EmptyLines() *RailwayRepository {
	r.lines = []railway.Line{}
	return r
}

func (r *RailwayRepository) AddLine(name string, code string, lineType string, isActive bool, announcement string) *RailwayRepository {
	r.lines = append(r.lines, railway.Line{
		Name:         name,
		Code:         code,
		Type:         lineType,
		IsActive:     isActive,
		Announcement: announcement,
	})
	return r
}

func (r *RailwayRepository) SetSourceError() *RailwayRepository {
	r.sourceError = NewError().
		SetCode(ErrorCode)
	return r
}
