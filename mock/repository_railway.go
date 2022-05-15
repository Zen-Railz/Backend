package mock

import (
	"zenrailz/errorr"
	"zenrailz/repository/railway"
)

func NewRailwayRepository() *RailwayRepository {
	return &RailwayRepository{}
}

func (r *RailwayRepository) Stations() (map[string]railway.Station, errorr.Entity) {
	return r.stations, r.sourceError
}

func (r *RailwayRepository) Lines() ([]railway.Line, errorr.Entity) {
	return r.lines, r.sourceError
}

func (r *RailwayRepository) Network() (map[string]*railway.NetworkNode, errorr.Entity) {
	return r.network, r.sourceError
}

func (r *RailwayRepository) EmptyStations() *RailwayRepository {
	r.stations = make(map[string]railway.Station)
	return r
}

func (r *RailwayRepository) AddStation(name string, prefix string, number int) *RailwayRepository {
	if r.stations == nil {
		r.EmptyStations()
	}

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
	if r.lines == nil {
		r.EmptyLines()
	}

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

func (r *RailwayRepository) EmptyNetwork() *RailwayRepository {
	r.network = make(map[string]*railway.NetworkNode)
	return r
}

func (r *RailwayRepository) AddNetworkNode(stationName string, stationIdentities map[string]railway.StationIdentity, adjacentNodes map[string]*railway.NetworkNode) *RailwayRepository {
	if r.network == nil {
		r.EmptyNetwork()
	}

	r.network[stationName] = &railway.NetworkNode{
		StationName:       stationName,
		StationIdentities: stationIdentities,
		AdjacentNodes:     adjacentNodes,
	}

	return r
}

type RailwayRepository struct {
	stations    map[string]railway.Station
	lines       []railway.Line
	network     map[string]*railway.NetworkNode
	sourceError errorr.Entity
}

const (
	RailwayJourneyOrigin      = "mockOrigin"
	RailwayJourneyDestination = "mockDestination"
)
