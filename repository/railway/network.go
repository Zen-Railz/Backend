package railway

import (
	"strconv"
	"zenrailz/code"
	"zenrailz/errorr"
	"zenrailz/repository/common"
)

func (r *Repository) Network() (map[string]*NetworkNode, errorr.Entity) {
	stationNameMap := make(map[string]*NetworkNode)

	rows, queryErr := r.database.Query("select s.name, s.prefix, s.number, s.is_active, l.name from station s, line l where s.line = l.code order by s.prefix, s.number")
	if queryErr != nil {
		err := common.ParseError(code.DatabaseQueryFailure, "Unable to get stations from database.", queryErr)
		return stationNameMap, err.Trace()
	}
	defer rows.Close()

	var previousNetworkNode *NetworkNode
	var previousStationIdentityPrefix string

	for rows.Next() {
		var stationName string
		var stationIdentityPrefix string
		var stationIdentityNumber int
		var stationIdentityIsActive bool
		var stationIdentityLine string

		if scanErr := rows.Scan(&stationName, &stationIdentityPrefix, &stationIdentityNumber, &stationIdentityIsActive, &stationIdentityLine); scanErr != nil {
			err := common.ParseError(code.DatabaseRowScanFailure, "Unable to read a station from database.", scanErr)
			return stationNameMap, err.Trace()
		}

		stationIdentityCode := stationIdentityPrefix + strconv.Itoa(stationIdentityNumber)

		currentNetworkNode, stationExist := stationNameMap[stationName]
		if stationExist {
			currentNetworkNode.StationIdentities[stationIdentityCode] = StationIdentity{
				Prefix:   stationIdentityPrefix,
				Number:   stationIdentityNumber,
				IsActive: stationIdentityIsActive,
				Line:     stationIdentityLine,
			}
		} else {
			currentNetworkNode = &NetworkNode{
				StationName: stationName,
				StationIdentities: map[string]StationIdentity{
					stationIdentityCode: {
						Prefix:   stationIdentityPrefix,
						Number:   stationIdentityNumber,
						IsActive: stationIdentityIsActive,
						Line:     stationIdentityLine,
					},
				},
				AdjacentNodes: make(map[string]*NetworkNode),
			}
		}

		stationNameMap[stationName] = currentNetworkNode

		if previousStationIdentityPrefix == stationIdentityPrefix {
			currentNetworkNode.AdjacentNodes[previousNetworkNode.StationName] = previousNetworkNode
			previousNetworkNode.AdjacentNodes[currentNetworkNode.StationName] = currentNetworkNode
		}

		previousNetworkNode = currentNetworkNode
		previousStationIdentityPrefix = stationIdentityPrefix
	}

	return stationNameMap, nil
}

type NetworkNode struct {
	StationName       string
	StationIdentities map[string]StationIdentity
	AdjacentNodes     map[string]*NetworkNode
}
