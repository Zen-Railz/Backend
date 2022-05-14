package railway

import (
	"zenrailz/code"
	"zenrailz/errorr"
	"zenrailz/repository/common"
)

func (r *Repository) Stations() (map[string]Station, errorr.Entity) {
	stations := make(map[string]Station)

	rows, queryErr := r.database.Query("select s.name, s.prefix, s.number, s.is_active, l.name from station s, line l where s.line = l.code order by s.name, s.prefix, s.number")
	if queryErr != nil {
		err := common.ParseError(code.DatabaseQueryFailure, "Unable to get stations from database.", queryErr)
		return stations, err.Trace()
	}
	defer rows.Close()

	for rows.Next() {
		var stationName string
		var stationIdentityPrefix string
		var stationIdentityNumber int
		var stationIdentityIsActive bool
		var stationIdentityLine string

		if scanErr := rows.Scan(&stationName, &stationIdentityPrefix, &stationIdentityNumber, &stationIdentityIsActive, &stationIdentityLine); scanErr != nil {
			err := common.ParseError(code.DatabaseRowScanFailure, "Unable to read a station from database.", scanErr)
			return stations, err.Trace()
		}

		station, stationExist := stations[stationName]
		if stationExist {
			station.Identifiers = append(station.Identifiers, StationIdentity{
				Prefix:   stationIdentityPrefix,
				Number:   stationIdentityNumber,
				Line:     stationIdentityLine,
				IsActive: stationIdentityIsActive,
			})
			stations[stationName] = station
		} else {
			stations[stationName] = Station{
				Name: stationName,
				Identifiers: []StationIdentity{
					{
						Prefix:   stationIdentityPrefix,
						Number:   stationIdentityNumber,
						Line:     stationIdentityLine,
						IsActive: stationIdentityIsActive,
					},
				},
			}
		}
	}

	if rowErr := rows.Err(); rowErr != nil {
		err := common.ParseError(code.DatabaseRowError, "Database Row Operation encountered an error.", rowErr)
		return stations, err.Trace()
	}

	return stations, nil
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
