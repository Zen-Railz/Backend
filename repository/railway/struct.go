package railway

import "database/sql"

type Repository struct {
	database *sql.DB
}

type Station struct {
	Name        string
	Identifiers []StationIdentity
}

type StationIdentity struct {
	Prefix   string
	Number   int
	Line     string
	IsActive bool
}

type Line struct {
	Name         string
	Code         string
	Type         string
	IsActive     bool
	Announcement string
}

type NetworkNode struct {
	StationName          string
	StationIdentities    map[string]StationIdentity
	PreviousStationNames StationNameSet
	NextStationNames     StationNameSet
}

type StationNameSet map[string]struct{}

func (s StationNameSet) Add(name string) {
	s[name] = struct{}{}
}
