package mock

import (
	"zenrailz/errorr"
	"zenrailz/repository/railway"
)

func NewRailwayRepository() *RailwayRepository {
	return &RailwayRepository{}
}

type RailwayRepository struct {
	stations    map[string]railway.Station
	lines       []railway.Line
	sourceError errorr.Entity
}

func (r *RailwayRepository) Stations() (map[string]railway.Station, errorr.Entity) {
	return r.stations, r.sourceError
}

func (r *RailwayRepository) Lines() ([]railway.Line, errorr.Entity) {
	return r.lines, r.sourceError
}

func (r *RailwayRepository) Network() (map[string]*railway.NetworkNode, errorr.Entity) {
	return nil, nil
}

func (r *RailwayRepository) EmptyStations() *RailwayRepository {
	r.stations = make(map[string]railway.Station)
	return r
}

func (r *RailwayRepository) AddStation(name string, prefix string, number int) *RailwayRepository {
	station, stationExist := r.stations[name]
	if stationExist {
		station.Identifiers = append(station.Identifiers, railway.StationIdentity{
			Prefix: prefix,
			Number: number,
		})
		r.stations[name] = station
	} else {
		r.stations[name] = railway.Station{
			Name: name,
			Identifiers: []railway.StationIdentity{
				{
					Prefix: prefix,
					Number: number,
				},
			},
		}
	}
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
