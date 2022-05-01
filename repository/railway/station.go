package railway

import (
	"strconv"
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

func (r *Repository) Network() (map[string]*NetworkNode, map[string]*NetworkNode, errorr.Entity) {
	stationNameMap := make(map[string]*NetworkNode)
	stationIdentityCodeMap := make(map[string]*NetworkNode)

	rows, queryErr := r.database.Query("select s.name, s.prefix, s.number, s.is_active, l.name from station s, line l where s.line = l.code order by s.prefix, s.number")
	if queryErr != nil {
		err := common.ParseError(code.DatabaseQueryFailure, "Unable to get stations from database.", queryErr)
		return stationNameMap, stationIdentityCodeMap, err.Trace()
	}
	defer rows.Close()

	var previousNetworkNode *NetworkNode
	var previousStationIdentityPrefix string
	var previousStationName string

	for rows.Next() {
		var stationName string
		var stationIdentityPrefix string
		var stationIdentityNumber int
		var stationIdentityIsActive bool
		var stationIdentityLine string

		if scanErr := rows.Scan(&stationName, &stationIdentityPrefix, &stationIdentityNumber, &stationIdentityIsActive, &stationIdentityLine); scanErr != nil {
			err := common.ParseError(code.DatabaseRowScanFailure, "Unable to read a station from database.", scanErr)
			return stationNameMap, stationIdentityCodeMap, err.Trace()
		}

		stationIdentityCode := stationIdentityPrefix + strconv.Itoa(stationIdentityNumber)

		networkNode, stationExist := stationNameMap[stationName]
		if stationExist {
			networkNode.StationIdentities[stationIdentityCode] = StationIdentity{
				Prefix:   stationIdentityPrefix,
				Number:   stationIdentityNumber,
				IsActive: stationIdentityIsActive,
				Line:     stationIdentityLine,
			}
		} else {
			networkNode = &NetworkNode{
				StationName: stationName,
				StationIdentities: map[string]StationIdentity{
					stationIdentityCode: {
						Prefix:   stationIdentityPrefix,
						Number:   stationIdentityNumber,
						IsActive: stationIdentityIsActive,
						Line:     stationIdentityLine,
					},
				},
				PreviousStationNames: make(StationNameSet),
				NextStationNames:     make(StationNameSet),
			}
		}

		stationNameMap[stationName] = networkNode
		stationIdentityCodeMap[stationIdentityCode] = networkNode

		if previousStationIdentityPrefix == stationIdentityPrefix {
			networkNode.PreviousStationNames.Add(previousStationName)
			previousNetworkNode.NextStationNames.Add(stationName)
		}

		previousNetworkNode = networkNode
		previousStationIdentityPrefix = stationIdentityPrefix
		previousStationName = stationName
	}

	return stationNameMap, stationIdentityCodeMap, nil
}
